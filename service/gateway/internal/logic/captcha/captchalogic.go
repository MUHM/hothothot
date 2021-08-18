package logic

import (
	"context"

	"hothothot/common/tools/captcha"
	"hothothot/service/gateway/internal/svc"
	"hothothot/service/gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) CaptchaLogic {
	return CaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CaptchaLogic) Captcha(req types.CaptchaReq) (*types.CaptchaResp, error) {
	captchaId, captchaBase64, err := captcha.RandomDigits(6)
	if err != nil {
		return nil, err
	}
	return &types.CaptchaResp{
		CaptchaId:     captchaId,
		CaptchaBase64: captchaBase64,
	}, nil
}
