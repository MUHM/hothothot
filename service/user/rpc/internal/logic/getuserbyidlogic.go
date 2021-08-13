package logic

import (
	"context"

	"hothothot/service/user/rpc/internal/svc"
	"hothothot/service/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByIdLogic) GetUserById(in *user.IdReq) (*user.UserInfoReply, error) {
	one, err := l.svcCtx.UserModel.FindOne(in.Id)
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
