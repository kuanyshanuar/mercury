package telemetry

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
)

// InitTelemetryProvider - setups a new telemetry provider.
func InitTelemetryProvider(url, environment string) (func(ctx context.Context) error, error) {
	// Create the Jaeger exporter
	//
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("mercury"),
			attribute.String("environment", environment),
		)),
	)

	// Register our TracerProvider as the global so any imported
	// instrumentation in the future will default to using it.
	otel.SetTracerProvider(tp)

	return tp.Shutdown, err
}

// StartServiceSpan - starts service span.
func StartServiceSpan(ctx context.Context, serviceName string, method string) (context.Context, trace.Span) {
	ctx, span := otel.Tracer(serviceName).Start(ctx, method)
	return ctx, span
}

// EndSpan - end the span and returns trace id and span id.
func EndSpan(span trace.Span, err error) (traceID string, spanID string) {
	if span == nil {
		return
	}

	traceID = span.SpanContext().TraceID().String()
	spanID = span.SpanContext().SpanID().String()

	span.SetStatus(codes.Ok, "")
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		span.RecordError(err)
	}
	span.End()

	return
}
