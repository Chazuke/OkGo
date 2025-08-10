package otel

import (
	"okgo/logger"

	"github.com/go-logr/zapr"
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

	// Use zapr to bridge zap logger to logr.Logger interface for OpenTelemetry internal logging
	logrLogger := zapr.NewLogger(logger.Log)
	otel.SetLogger(logrLogger)
	
	return lp
}
