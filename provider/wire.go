//go:build wireinject
// +build wireinject

package provider

import (
	"github.com/google/wire"
)

func NewContainer() (*Provider, error) {
	wire.Build(
		AllProvider,
		wire.Struct(new(Provider), "*"),
	)
	return nil, nil
}
