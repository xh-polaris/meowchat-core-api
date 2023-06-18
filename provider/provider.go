package provider

import (
	"github.com/google/wire"
	"github.com/xh-polaris/meowchat-core-api/biz/adaptor"
	"github.com/xh-polaris/meowchat-core-api/biz/application"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure"
)

var provider *Provider

func Init() {
	var err error
	provider, err = NewContainer()
	if err != nil {
		panic(err)
	}
}

type Provider struct {
	*adaptor.Adaptor
	*application.Application
	*infrastructure.Infrastructure
}

func Get() *Provider {
	return provider
}

var AllProvider = wire.NewSet(
	adaptor.ProviderSet,
	application.ProviderSet,
	infrastructure.ProviderSet,
)
