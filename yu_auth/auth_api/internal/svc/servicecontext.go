package svc

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"server/core"
	"server/yu_auth/auth_api/internal/config"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	Redis  *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB:     core.InitGorm(c.Mysql.DataSource),
		Redis:  core.InitRedis(c.Redis.Addr, c.Redis.Password, c.Redis.DB),
	}
}
