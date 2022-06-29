package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	ForwardConf ForwardConf
	Auth JwtAuth
}

type ForwardConf struct {
	Api1 ForwardApiModel
}

type ForwardApiModel struct {
	Host string
}

type JwtAuth struct {
	AccessSecret string
	AccessExpire int64
}