package zap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(level string) (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	if level == "debug" {
		config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	} else if level == "warn" {
		config.Level = zap.NewAtomicLevelAt(zapcore.WarnLevel)
	} else if level == "error" {
		config.Level = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	}
	return config.Build()
}

func Check(level string) bool {
	logger, _ := NewLogger(level)
	return logger.Check(zap.InfoLevel, "test message").ShouldSample(nil)
}