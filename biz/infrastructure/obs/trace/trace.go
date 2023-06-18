package trace

import (
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

func Init() {
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(b3.New(), propagation.Baggage{}, propagation.TraceContext{}))
}
