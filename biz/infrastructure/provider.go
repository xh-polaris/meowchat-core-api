package infrastructure

import (
	"github.com/google/wire"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_collection"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/platform_authentication"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/platform_sts"
)

var RPCSet = wire.NewSet(
	meowchat_collection.MeowchatCollectionSet,
	platform_authentication.PlatformAuthenticationSet,
	platform_sts.PlatformStsSet,
)

var ProviderSet = wire.NewSet(
	config.NewConfig,
	RPCSet,
)
