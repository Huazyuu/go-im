package svc

import (
	"github.com/go-redis/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"server/core"
	"server/yu_user/user_api/internal/config"
	"server/yu_user/user_rpc/types/user_rpc"
	"server/yu_user/user_rpc/user"
)

type ServiceContext struct {
	Config  config.Config
	DB      *gorm.DB
	Redis   *redis.Client
	UserRpc user_rpc.UserClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		DB:      core.InitGorm(c.Mysql.DataSource),
		Redis:   core.InitRedis(c.Redis.Addr, c.Redis.Password, c.Redis.DB),
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
