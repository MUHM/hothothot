package logic

import (
	"context"
	"errors"

	"hothothot/service/user/model"
	"hothothot/service/user/rpc/internal/svc"
	"hothothot/service/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByNameLogic {
	return &GetUserByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByNameLogic) GetUserByName(in *user.NameReq) (*user.UserInfoReply, error) {
	one, err := l.svcCtx.UserModel.FindOneByUsername(in.Name)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errors.New("用户名不存在")
	default:
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
