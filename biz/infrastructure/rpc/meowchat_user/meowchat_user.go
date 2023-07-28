package meowchat_user

import (
	"github.com/google/wire"
	"github.com/xh-polaris/gopkg/kitex/client"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
	user "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/user/userservice"
)

type IMeowchatUser interface {
	user.Client
}

type MeowchatUser struct {
	user.Client
}

var MeowchatUserSet = wire.NewSet(
	NewMeowchatUser,
	wire.Struct(new(MeowchatUser), "*"),
	wire.Bind(new(IMeowchatUser), new(*MeowchatUser)),
)

func NewMeowchatUser(config *config.Config) user.Client {
	return client.NewClient(config.Name, "meowchat.user", user.NewClient)
}
