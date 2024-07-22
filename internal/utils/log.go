package utils

import (
	"log"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger is a globally accessible logger
var Logger *zap.SugaredLogger

// CreateLogger creates a logger that writes to a provided path
func CreateLogger(path string) (*zap.SugaredLogger, error) {
	// Create zap production config
	config := zap.NewProductionConfig()
	config.DisableCaller = true
	config.OutputPaths = []string{path}
	config.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(time.RFC1123))
	}

	// Build logger
	logger, err := config.Build()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := logger.Sync(); err != nil {
			log.Fatalf("Failed to sync logger: %v", err)
		}
	}()

	return logger.Sugar(), nil
}

// Log is a wrapper function to handle debug messaging for bubbletea applications
func Log(logType string, msg string, keysAndValues ...interface{}) {
	// Check if Logger is uninitialized
	if Logger == nil {
		return
	}

	switch logType {
	case "Infow":
		Logger.Infow(msg, keysAndValues...)
	case "Errorw":
		Logger.Errorw(msg, keysAndValues...)
	default:
		Logger.Errorw("Unknown log type", "log type", logType)
	}
}
