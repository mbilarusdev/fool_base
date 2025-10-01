package log

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

var logger *zap.Logger

func Init() {
	logLevelStr := os.Getenv("LOG_LEVEL")
	var logLevel zap.AtomicLevel

	switch logLevelStr {
	case "debug":
		logLevel = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		logLevel = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		logLevel = zap.NewAtomicLevelAt(zap.WarnLevel)
	default:
		logLevel = zap.NewAtomicLevelAt(zap.ErrorLevel)
	}

	config := zap.Config{
		Encoding:         "json",
		Level:            logLevel,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig:    zap.NewProductionEncoderConfig(),
	}

	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)

	logger, _ = config.Build()

	Info("Logger Zap inited!")
}

func Err(err error, msg string, fields ...zap.Field) error {
	logger.Error(fmt.Sprintf("%v: %v", msg, err), fields...)
	return err
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Sync() {
	logger.Sync()
}
