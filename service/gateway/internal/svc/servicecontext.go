package svc

import (
	"hothothot/common/tools/captcha"
	"hothothot/service/gateway/internal/config"
	"hothothot/service/gateway/internal/middleware"
	"hothothot/service/system/rpc/systemclient"
	"hothothot/service/user/rpc/userclient"

	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	SystemRpc  systemclient.System
	UserRpc    userclient.User
	CasbinRbac rest.Middleware
	Captcha    captcha.Captcha
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		SystemRpc:  systemclient.NewSystem(zrpc.MustNewClient(c.SystemRpc)),
		UserRpc:    userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		CasbinRbac: middleware.NewCasbinRbacMiddleware(c.Mysql.DataSource).Handle,
		Captcha:    captcha.NewCaptcha(c.CacheRedis),
	}
}
