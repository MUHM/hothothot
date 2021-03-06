// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

//go:generate mockgen -destination ./user_mock.go -package userclient -source $GOFILE

package userclient

import (
	"context"

	"hothothot/service/user/rpc/user"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	OtpReply      = user.OtpReply
	IdReq         = user.IdReq
	NameReq       = user.NameReq
	MailReq       = user.MailReq
	UserInfoReply = user.UserInfoReply
	OtpReq        = user.OtpReq

	User interface {
		GetUserById(ctx context.Context, in *IdReq) (*UserInfoReply, error)
		GetUserByName(ctx context.Context, in *NameReq) (*UserInfoReply, error)
		GetUserByMail(ctx context.Context, in *MailReq) (*UserInfoReply, error)
		SaveOtpSecret(ctx context.Context, in *OtpReq) (*OtpReply, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) GetUserById(ctx context.Context, in *IdReq) (*UserInfoReply, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUserById(ctx, in)
}

func (m *defaultUser) GetUserByName(ctx context.Context, in *NameReq) (*UserInfoReply, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUserByName(ctx, in)
}

func (m *defaultUser) GetUserByMail(ctx context.Context, in *MailReq) (*UserInfoReply, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUserByMail(ctx, in)
}

func (m *defaultUser) SaveOtpSecret(ctx context.Context, in *OtpReq) (*OtpReply, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.SaveOtpSecret(ctx, in)
}
