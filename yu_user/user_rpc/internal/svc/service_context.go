package svc

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"server/core"
	"server/yu_user/user_rpc/internal/config"
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
		Redis:  core.InitRedis(c.RedisRpc.Addr, c.RedisRpc.Password, c.RedisRpc.DB),
	}
}
