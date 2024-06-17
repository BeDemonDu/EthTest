package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	EthClient EthClientConfig `json:"ethClient"`
}

type EthClientConfig struct {
	BaseUrl string `json:"baseUrl"`
}
