package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	OkGoLogger struct {
		env     string
		Log     *zap.Logger
		service string
	}

	OkGoLogOptions struct {
		WithTimestamp bool
		LogLevel      OkGoLogLevel
	}

	OkGoLogLevel = zapcore.Level
)

const (
	Debug OkGoLogLevel = zapcore.DebugLevel
	Info  OkGoLogLevel = zapcore.InfoLevel
	Warn  OkGoLogLevel = zapcore.WarnLevel
	Error OkGoLogLevel = zapcore.ErrorLevel
)

func DefaultOpts() *OkGoLogOptions {
	return &OkGoLogOptions{
		WithTimestamp: true,
		LogLevel:      Debug,
	}
}

func New(env string, service string, opts OkGoLogOptions) *OkGoLogger {
	// Create config based on environment
	var config zap.Config
	if env == "production" {
		config = zap.NewProductionConfig()
		config.EncoderConfig.TimeKey = "timestamp"
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	} else {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	// Set the log level directly (no conversion needed now)
	config.Level = zap.NewAtomicLevelAt(zapcore.Level(opts.LogLevel))

	// Disable timestamp if requested
	if !opts.WithTimestamp {
		config.EncoderConfig.TimeKey = ""
	}

	// Build the logger
	logger, err := config.Build(
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
	if err != nil {
		panic(err)
	}

	// Add service and environment fields
	logger = logger.With(
		zap.String("service", service),
		zap.String("env", env),
	)

	return &OkGoLogger{
		env:     env,
		Log:     logger,
		service: service,
	}
}

// Helper methods to match common logging patterns
func (l *OkGoLogger) Debug(msg string, fields ...zap.Field) {
	l.Log.Debug(msg, fields...)
}

func (l *OkGoLogger) Info(msg string, fields ...zap.Field) {
	l.Log.Info(msg, fields...)
}

func (l *OkGoLogger) Warn(msg string, fields ...zap.Field) {
	l.Log.Warn(msg, fields...)
}

func (l *OkGoLogger) Error(msg string, fields ...zap.Field) {
	l.Log.Error(msg, fields...)
}

func (l *OkGoLogger) Fatal(msg string, fields ...zap.Field) {
	l.Log.Fatal(msg, fields...)
}

// Sync flushes any buffered log entries
func (l *OkGoLogger) Sync() error {
	return l.Log.Sync()
}
