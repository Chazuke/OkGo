package otel

import (
	"okgo/logger"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func initGlobalTracerProvider(serviceName string, e sdktrace.SpanExporter) *sdktrace.TracerProvider {
	// exporter, err := stdout.New(stdout.WithPrettyPrint())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(e),
		sdktrace.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String(serviceName),
			)),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp
}

func initGlobalLoggerProvider(serviceName string, e log.Exporter, logger *logger.OkGoLogger) *log.LoggerProvider {
	lp := log.NewLoggerProvider(
		log.WithProcessor(
			log.NewBatchProcessor(e),
		),
		log.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String(serviceName),
			),
		),
	)

	// Note: zap doesn't directly implement logr.Logger interface required by otel.SetLogger
	// If you need OpenTelemetry to use zap for its internal logging, consider using a bridge
	// like "github.com/go-logr/zapr" to create a logr.Logger from zap.Logger
	// For now, OpenTelemetry will use its default logger
	_ = logger // Mark as used to avoid unused parameter warning
	
	return lp
}
