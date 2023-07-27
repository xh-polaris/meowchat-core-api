package meowchat_collection

import (
	"github.com/google/wire"
	"github.com/xh-polaris/gopkg/kitex/client"
	collection "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/collection/collectionservice"

	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
)

type IMeowchatCollection interface {
	collection.Client
}

type MeowchatCollection struct {
	collection.Client
}

var MeowchatCollectionSet = wire.NewSet(
	NewMeowchatCollection,
	wire.Struct(new(MeowchatCollection), "*"),
	wire.Bind(new(IMeowchatCollection), new(*MeowchatCollection)),
)

func NewMeowchatCollection(config *config.Config) collection.Client {
	return client.NewClient(config.Name, "meowchat.collection", collection.NewClient)
}
