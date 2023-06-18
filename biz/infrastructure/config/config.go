package config

import (
	"os"

	"github.com/google/wire"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
)

type Auth struct {
	AccessSecret string
	AccessExpire int64
}

type Config struct {
	Auth          Auth
	CdnHost       string
	AuthRPC       zrpc.RpcClientConf
	CollectionRPC zrpc.RpcClientConf
	MomentRPC     zrpc.RpcClientConf
	SystemRPC     zrpc.RpcClientConf
	LikeRPC       zrpc.RpcClientConf
	UserRPC       zrpc.RpcClientConf
	StsRPC        zrpc.RpcClientConf
	CommentRPC    zrpc.RpcClientConf
	PostRPC       zrpc.RpcClientConf
}

func NewConfig() (*Config, error) {
	c := new(Config)
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = "etc/config.yaml"
	}
	err := conf.Load(path, c)
	return c, err
}

var ProviderSet = wire.NewSet(
	NewConfig,
)
