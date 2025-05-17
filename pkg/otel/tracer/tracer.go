package tracer

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

var tracer trace.Tracer

func Init(t trace.Tracer) {
	tracer = t
}

func Start(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	return tracer.Start(ctx, spanName, opts...)
}
