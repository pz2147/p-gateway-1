package config

import "github.com/tal-tech/go-zero/rest"

type Config struct {
	rest.RestConf
	ForwardConf ForwardConf
}

type ForwardConf struct {
	Api1 ForwardApiModel
}

type ForwardApiModel struct {
	Host string
}
