package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	zapLogger *zap.Logger
)

func InitZaplogger() {
	// Create Zap logger
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	var err error
	zapLogger, err = config.Build()
	if err != nil {
		log.Fatalf("Failed to create logger: %v", err)
	}
}

func Info(msg string, fields ...zap.Field) {
	zapLogger.Info(msg, fields...)
}

// Debug logs a message at Debug level.
func Debug(msg string, fields ...zap.Field) {
	zapLogger.Debug(msg, fields...)
}

// Error logs a message at Error level.
func Error(msg string, fields ...zap.Field) {
	zapLogger.Error(msg, fields...)
}

// Fatal logs a message at Fatal level.
func Fatal(msg string, fields ...zap.Field) {
	zapLogger.Fatal(msg, fields...)
}
