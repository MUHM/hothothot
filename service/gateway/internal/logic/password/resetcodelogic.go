package logic

import (
	"context"

	"hothothot/service/gateway/internal/svc"
	"hothothot/service/gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ResetCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResetCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) ResetCodeLogic {
	return ResetCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResetCodeLogic) ResetCode(req types.PwdResetCodeReq) (*types.PwdResetCodeResp, error) {
	// todo: add your logic here and delete this line

	return &types.PwdResetCodeResp{}, nil
}
