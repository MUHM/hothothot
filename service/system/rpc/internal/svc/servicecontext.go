package svc

import (
	"hothothot/service/system/model"
	"hothothot/service/system/rpc/internal/config"

	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config           config.Config
	SysSettingsModel model.SysSettingsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:           c,
		SysSettingsModel: model.NewSysSettingsModel(conn, c.CacheRedis),
	}
}
