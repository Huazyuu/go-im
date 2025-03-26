package svc

import (
	"gorm.io/gorm"
	"server/core"
	"server/yu_auth/auth_api/internal/config"
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
