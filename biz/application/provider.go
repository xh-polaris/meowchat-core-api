package application

import (
	"github.com/google/wire"
	"github.com/xh-polaris/meowchat-core-api/biz/application/service"
)

var ProviderSet = wire.NewSet(
	service.CollectionServiceSet,
)
