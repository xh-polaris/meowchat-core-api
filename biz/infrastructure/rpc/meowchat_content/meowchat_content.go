package meowchat_content

import (
	"github.com/google/wire"
	"github.com/xh-polaris/gopkg/kitex/client"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
	content "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/content/contentservice"
)

type IMeowchatContent interface {
	content.Client
}

type MeowchatContent struct {
	content.Client
}

var MeowchatContentSet = wire.NewSet(
	NewMeowchatContent,
	wire.Struct(new(MeowchatContent), "*"),
	wire.Bind(new(IMeowchatContent), new(*MeowchatContent)),
)

func NewMeowchatContent(config *config.Config) content.Client {
	return client.NewClient(config.Name, "meowchat.content", content.NewClient)
}
