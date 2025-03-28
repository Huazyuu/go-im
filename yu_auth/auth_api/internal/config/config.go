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

	Auth struct {
		AccessSecret string
		AccessExpire int
	}

	WhiteList []string // 白名单

	Redis struct {
		Addr     string
		Password string
		DB       int
	}

	OpenLoginList []struct {
		Name string
		Icon string
		Href string
	}

	GitHub struct {
		ClientID     string
		ClientSecret string
		Redirect     string
	}

	Gitee struct {
		ClientID     string
		ClientSecret string
		Redirect     string
	}
}
