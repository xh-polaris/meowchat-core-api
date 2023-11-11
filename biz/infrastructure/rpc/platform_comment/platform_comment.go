package platform_comment

import (
	"github.com/google/wire"
	"github.com/xh-polaris/gopkg/kitex/client"
	comment "github.com/xh-polaris/service-idl-gen-go/kitex_gen/platform/comment/commentservice"

	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
)

type IPlatformCommment interface {
	comment.Client
}

type PlatformComment struct {
	comment.Client
}

var PlatformCommentSet = wire.NewSet(
	NewPlatformComment,
	wire.Struct(new(PlatformComment), "*"),
	wire.Bind(new(IPlatformCommment), new(*PlatformComment)),
)

func NewPlatformComment(config *config.Config) comment.Client {
	return client.NewClient(config.Name, "platform.comment", comment.NewClient)
}
