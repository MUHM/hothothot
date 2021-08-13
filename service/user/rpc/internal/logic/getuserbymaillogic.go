package logic

import (
	"context"

	"hothothot/service/user/rpc/internal/svc"
	"hothothot/service/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserByMailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByMailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByMailLogic {
	return &GetUserByMailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByMailLogic) GetUserByMail(in *user.MailReq) (*user.UserInfoReply, error) {
	one, err := l.svcCtx.UserModel.FindOneByUsername(in.Mail)
	if err != nil {
		return nil, err
	}

	return &user.UserInfoReply{
		Id:       one.Id,
		Username: one.Username,
		Nickname: one.Nickname,
		Gender:   one.Gender,
		Password: one.Password,
		Mail:     one.Mail,
	}, nil
}
