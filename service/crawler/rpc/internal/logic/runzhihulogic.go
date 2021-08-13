package logic

import (
	"context"

	"hothothot/service/crawler/rpc/crawler"
	"hothothot/service/crawler/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
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

	return &crawler.ZhiHuResp{}, nil
}
