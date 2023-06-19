package adaptor

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	ExtractorSet,
	HandlerSet,
)
