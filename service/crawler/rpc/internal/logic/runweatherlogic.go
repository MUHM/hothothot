package logic

import (
	"context"

	"hothothot/service/crawler/rpc/crawler"
	"hothothot/service/crawler/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type RunWeatherLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRunWeatherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RunWeatherLogic {
	return &RunWeatherLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RunWeatherLogic) RunWeather(in *crawler.WeatherReq) (*crawler.WeatherNoResp, error) {
	// todo: add your logic here and delete this line

	return &crawler.WeatherNoResp{}, nil
}
