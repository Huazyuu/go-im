package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Mysql struct {
		DataSource string
	}
	Auth struct {
		AccessSecret string
		AccessExpire int
	}
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
