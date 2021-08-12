package logic

import (
	"context"

	"hothothot/service/system/rpc/internal/svc"
	"hothothot/service/system/rpc/system"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetByNameLogic {
	return &GetByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetByNameLogic) GetByName(in *system.NameReq) (*system.NameResp, error) {
	one, err := l.svcCtx.SysSettingsModel.FindOneByName(in.Name)
	if err != nil {
		return nil, err
	}
	return &system.NameResp{
		Id:       (*one).Id,
		Classify: (*one).Classify,
		Content:  (*one).Content,
	}, nil
}
