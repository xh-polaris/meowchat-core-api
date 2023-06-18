package adaptor

import "github.com/google/wire"

type Adaptor struct {
	Extractor IExtractor
}

var ProviderSet = wire.NewSet(
	ExtractorSet,
	wire.Struct(new(Adaptor), "*"),
)
