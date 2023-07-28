package config

import (
	"github.com/zeromicro/go-zero/core/service"
	"os"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
)

type Auth struct {
	AccessSecret string
	AccessExpire int64
}

type Config struct {
	service.ServiceConf
	ListenOn string
	Auth     Auth
	CosHost  string
	CdnHost  string
	CIHost   string
	CosApi   struct {
		SecretId  string
		SecretKey string
	}
	AuthRPC    zrpc.RpcClientConf
	MomentRPC  zrpc.RpcClientConf
	SystemRPC  zrpc.RpcClientConf
	LikeRPC    zrpc.RpcClientConf
	UserRPC    zrpc.RpcClientConf
	StsRPC     zrpc.RpcClientConf
	CommentRPC zrpc.RpcClientConf
	PostRPC    zrpc.RpcClientConf
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
	return c, nil
}
