package infrastructure

import (
	"github.com/google/wire"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_content"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_system"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_user"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/platform_comment"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/platform_sts"
)

var RPCSet = wire.NewSet(
	meowchat_content.MeowchatContentSet,
	platform_sts.PlatformStsSet,
	platform_comment.PlatformCommentSet,
	meowchat_user.MeowchatUserSet,
	meowchat_system.MeowchatSystemSet,
)

var ProviderSet = wire.NewSet(
	config.NewConfig,
	RPCSet,
)
