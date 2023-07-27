package infrastructure

import (
	"github.com/google/wire"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_collection"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_like"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_moment"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_post"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_user"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/platform_comment"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/platform_sts"
)

var RPCSet = wire.NewSet(
	meowchat_collection.MeowchatCollectionSet,
	platform_sts.PlatformStsSet,
	platform_comment.PlatformCommentSet,
	meowchat_user.MeowchatUserSet,
	meowchat_moment.MeowchatMomentSet,
	meowchat_like.MeowchatLikeSet,
	meowchat_post.MeowchatPostSet,
)

var ProviderSet = wire.NewSet(
	config.NewConfig,
	RPCSet,
)
