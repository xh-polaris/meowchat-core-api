package infrastructure

import (
	"github.com/google/wire"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc"
)

type Infrastructure struct {
}

var RPCSet = wire.NewSet(
	rpc.MeowchatCollectionSet,
)

var ProviderSet = wire.NewSet(
	wire.Struct(new(Infrastructure), "*"),
	config.ProviderSet,
	RPCSet,
)
