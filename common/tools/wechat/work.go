package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
	jsonparse "hothothot/common/tools/json"
	wechatTypes "hothothot/common/types/wechat"
	"time"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/syncx"
	"github.com/valyala/fasthttp"
)

var (
	ExclusiveCalls = syncx.NewSharedCalls()
	Stats          = cache.NewStat("wechat")
	ErrorCaptcha   = errors.New("wechat: no access_token in result set")

	CachePrefix    = "wechat:work"
	CacheExpiry    = 7000 * time.Second
	GetTokenApi    = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s"
	SendMessageApi = "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s"
	TokenKey       = "access_token"
)

type (
	WorkConf struct {
		CorpId     string `json:"corpId"`
		CorpSecret string `json:"corpSecret"`
		AgentId    string `json:"agentId"`
	}
	WechatWork interface {
		GetToken() (string, error)
		SendText(text string)
	}

	defaultWechatWork struct {
		CacheRedis cache.Cache
		WorkConf   WorkConf
	}
)

func NewWechatWork(cacheConf cache.CacheConf, workConf WorkConf) WechatWork {
	return &defaultWechatWork{
		CacheRedis: cache.New(cacheConf, ExclusiveCalls, Stats, ErrorCaptcha, func(o *cache.Options) { o.Expiry = CacheExpiry }),
		WorkConf:   workConf,
	}
}

func (c *defaultWechatWork) GetToken() (string, error) {
	var accessToken string
	corpId := c.WorkConf.CorpId
	corpSecret := c.WorkConf.CorpSecret
	key := fmt.Sprintf("%s:%s:%s", CachePrefix, TokenKey, corpId)
	c.CacheRedis.Get(key, &accessToken)
	if accessToken != "" {
		return accessToken, nil
	}
	client := fasthttp.Client{}
	httpReq := fasthttp.AcquireRequest()
	httpRes := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(httpReq)
	defer fasthttp.ReleaseResponse(httpRes)
	httpReq.SetRequestURI(fmt.Sprintf(GetTokenApi, corpId, corpSecret))
	if err := client.DoTimeout(httpReq, httpRes, 30*time.Second); err != nil {
		return "", errors.New("请求超时")
	}
	if httpRes.StatusCode() != 200 {
		return "", errors.New("invalid statuscode")
	}
	tokenResponse := jsonparse.ParseJson(string(httpRes.Body()))
	if tokenResponse["errcode"].(float64) != 0 {
		return "", errors.New("invalid corpsecret")
	}
	accessToken = tokenResponse[TokenKey].(string)
	c.CacheRedis.SetWithExpire(fmt.Sprintf("%s:%s:%s", CachePrefix, TokenKey, corpId), accessToken, time.Duration(tokenResponse["expires_in"].(float64))*time.Second)
	return accessToken, nil
}

func (c *defaultWechatWork) SendText(text string) {
	postData := &wechatTypes.WechatJsonData{
		ToUser:                 "@all",
		AgentId:                c.WorkConf.AgentId,
		MsgType:                "text",
		DuplicateCheckInterval: 600,
	}
	postData.Text.Content = text
	accessToken, _ := c.GetToken()
	c.PostMsg(*postData, fmt.Sprintf(SendMessageApi, accessToken))
}

func (c *defaultWechatWork) PostMsg(postData wechatTypes.WechatJsonData, url string) (string, error) {
	client := fasthttp.Client{}
	httpReq := fasthttp.AcquireRequest()
	httpRes := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(httpReq)
	defer fasthttp.ReleaseResponse(httpRes)
	httpReq.Header.SetContentType("application/json")
	httpReq.Header.SetMethod("POST")
	httpReq.SetRequestURI(url)
	postJson, _ := json.Marshal(postData)
	requestBody := []byte(postJson)
	httpReq.SetBody(requestBody)
	if err := client.DoTimeout(httpReq, httpRes, 30*time.Second); err != nil {
		return "", errors.New("请求超时")
	}
	if httpRes.StatusCode() != 200 {
		return "", errors.New("invalid statuscode")
	}
	res := string(httpRes.Body())
	return res, nil
}
