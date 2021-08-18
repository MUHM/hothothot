package svc

import (
	"hothothot/service/crawler/rpc/internal/config"
	"hothothot/service/system/rpc/systemclient"

	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	SystemRpc systemclient.System
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		SystemRpc: systemclient.NewSystem(zrpc.MustNewClient(c.SystemRpc)),
	}
}
