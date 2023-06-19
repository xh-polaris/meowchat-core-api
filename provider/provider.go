package provider

import (
	"github.com/google/wire"

	"github.com/xh-polaris/meowchat-core-api/biz/adaptor"
	"github.com/xh-polaris/meowchat-core-api/biz/application"
	"github.com/xh-polaris/meowchat-core-api/biz/application/service"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
)

var provider *Provider

func Init() {
	var err error
	provider, err = NewProvider()
	if err != nil {
		panic(err)
	}
}

// Provider 提供controller依赖的对象
type Provider struct {
	Extractor         adaptor.IExtractor
	Handler           adaptor.IHandler
	Config            *config.Config
	CollectionService service.ICollectionService
}

func Get() *Provider {
	return provider
}

var AllProvider = wire.NewSet(
	adaptor.ProviderSet,
	application.ProviderSet,
	infrastructure.ProviderSet,
)
