package application

import (
	"github.com/google/wire"
	"github.com/xh-polaris/meowchat-core-api/biz/application/service"
)

type Application struct {
	CollectionService service.ICollectionService
}

var ProviderSet = wire.NewSet(
	service.CollectionServiceSet,
	wire.Struct(new(Application), "*"),
)
