package logic

import (
	"context"

	"hothothot/service/gateway/internal/svc"
	"hothothot/service/gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ResetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResetLogic(ctx context.Context, svcCtx *svc.ServiceContext) ResetLogic {
	return ResetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResetLogic) Reset(req types.PwdResetReq) (*types.PwdResetResp, error) {
	// todo: add your logic here and delete this line

	return &types.PwdResetResp{}, nil
}
