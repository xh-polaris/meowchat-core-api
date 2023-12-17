package meowchat_system

import (
	"github.com/google/wire"
	"github.com/xh-polaris/gopkg/kitex/client"
	system "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/system/systemrpc"

	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
)

type IMeowchatSystem interface {
	system.Client
}

type MeowchatSystem struct {
	system.Client
}

var MeowchatSystemSet = wire.NewSet(
	NewMeowchatSystem,
	wire.Struct(new(MeowchatSystem), "*"),
	wire.Bind(new(IMeowchatSystem), new(*MeowchatSystem)),
)

func NewMeowchatSystem(config *config.Config) system.Client {
	return client.NewClient(config.Name, "meowchat.system", system.NewClient)
}
