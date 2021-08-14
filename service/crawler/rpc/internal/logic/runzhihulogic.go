package logic

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	baseTypes "hothothot/common/types"
	"hothothot/service/crawler/rpc/crawler"
	"hothothot/service/crawler/rpc/internal/svc"
	"hothothot/service/system/rpc/systemclient"

	"github.com/tal-tech/go-zero/core/logx"
	"github.com/valyala/fasthttp"
)

type RunZhiHuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRunZhiHuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RunZhiHuLogic {
	return &RunZhiHuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RunZhiHuLogic) RunZhiHu(in *crawler.ZhiHuReq) (*crawler.ZhiHuResp, error) {
	// todo: add your logic here and delete this line

	zhihuSettings, _ := l.svcCtx.SystemRpc.GetByName(l.ctx, &systemclient.NameReq{Name: "crawler_zhihu"})
	var crawlerZhihu baseTypes.CrawlerZhihu
	json.Unmarshal([]byte(zhihuSettings.Content), &crawlerZhihu)
	client := fasthttp.Client{}
	httpReq := fasthttp.AcquireRequest()
	httpRes := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(httpReq)
	defer fasthttp.ReleaseResponse(httpRes)
	httpReq.SetRequestURI(crawlerZhihu.URL)
	httpReq.Header.Add("Accept", crawlerZhihu.Accept)
	httpReq.Header.SetUserAgent(crawlerZhihu.UserAgent)
	if err := client.DoTimeout(httpReq, httpRes, 30*time.Second); err != nil {
		return nil, errors.New("请求超时")
	}
	if httpRes.StatusCode() != 200 {
		return nil, errors.New("invalid statuscode")
	}
	// var result types.ZhihuResponse
	// json.Unmarshal(httpRes.Body(), &result)
	return &crawler.ZhiHuResp{}, nil
}
