package logic

import (
	"context"

	"hothothot/service/user/rpc/internal/svc"
	"hothothot/service/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type SaveOtpSecretLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveOtpSecretLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveOtpSecretLogic {
	return &SaveOtpSecretLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveOtpSecretLogic) SaveOtpSecret(in *user.OtpReq) (*user.OtpReply, error) {
	err := l.svcCtx.UserModel.SaveOtpSecret(in.Id, in.Secret)
	if err != nil {
		return nil, err
	}
	return &user.OtpReply{}, nil
}
