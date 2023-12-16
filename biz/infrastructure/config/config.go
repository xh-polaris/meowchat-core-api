package config

import (
	"os"

	"github.com/zeromicro/go-zero/core/service"

	"github.com/zeromicro/go-zero/core/conf"
)

var config *Config

type Auth struct {
	AccessSecret string
	AccessExpire int64
}

type Fish struct {
	SignIn  []int64
	Like    []int64
	Content []int64
	Comment []int64
}

type Config struct {
	service.ServiceConf
	ListenOn           string
	Auth               Auth
	CdnHost            string
	Fish               Fish
	DefaultCommunityId string
	MinVersion         string
}

func NewConfig() (*Config, error) {
	c := new(Config)
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = "etc/config.yaml"
	}
	err := conf.Load(path, c)
	if err != nil {
		return nil, err
	}
	err = c.SetUp()
	if err != nil {
		return nil, err
	}
	config = c
	return c, nil
}

func GetConfig() *Config {
	return config
}
