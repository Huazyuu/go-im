package svc

import (
	"gorm.io/gorm"
	"server/core"
	"server/yu_user/user_rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB:     core.InitGorm(c.Mysql.DataSource),
	}
}
