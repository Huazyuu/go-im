package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	UserRpc zrpc.RpcClientConf

	Mysql struct {
		DataSource string
	}
	Redis struct {
		Addr     string
		Password string
		DB       int
	}
	Etcd string
}
