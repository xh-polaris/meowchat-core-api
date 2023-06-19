package infrastructure

import (
	"github.com/google/wire"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc"
)

var RPCSet = wire.NewSet(
	rpc.MeowchatCollectionSet,
)

var ProviderSet = wire.NewSet(
	config.ProviderSet,
	RPCSet,
)
