package infrastructure

import (
	"github.com/google/wire"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_collection"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/platform_authentication"
)

var RPCSet = wire.NewSet(
	meowchat_collection.MeowchatCollectionSet,
	platform_authentication.PlatformAuthenticationSet,
)

var ProviderSet = wire.NewSet(
	config.NewConfig,
	RPCSet,
)
