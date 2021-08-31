package wechat

import (
	"errors"
	"hothothot/common/tools/jsonparse"
	"time"

	"github.com/valyala/fasthttp"
)

var GetTokenApi = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s"
var SendMessageApi = "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s"

func GetToken(corpid string, corpsecret string) (string, error) {

	client := fasthttp.Client{}
	httpReq := fasthttp.AcquireRequest()
	httpRes := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(httpReq)
	defer fasthttp.ReleaseResponse(httpRes)
	httpReq.SetRequestURI(GetTokenApi)
	if err := client.DoTimeout(httpReq, httpRes, 30*time.Second); err != nil {
		return "", errors.New("请求超时")
	}
	if httpRes.StatusCode() != 200 {
		return "", errors.New("invalid statuscode")
	}
	tokenResponse := jsonparse.ParseJson(httpReq.Body())
	if tokenResponse["errcode"] != 0 {
		return "", errors.New("invalid corpsecret")
	}
	accessToken := tokenResponse["access_token"].(string)
	return accessToken
}
