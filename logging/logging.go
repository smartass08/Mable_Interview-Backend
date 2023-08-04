package logging

import (
	stdlog "log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func Init(logLevel zapcore.Level) {
	stdlog.SetFlags(stdlog.LstdFlags | stdlog.Lshortfile)

	var cfg zap.Config
	var options []zap.Option

	// set development configuration options
	cfg = zap.NewDevelopmentConfig()
	cfg.Level.SetLevel(logLevel)
	cfg.DisableStacktrace = true
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	var err error
	logger, err = cfg.Build(options...)
	if err != nil {
		stdlog.Fatal("logger not initialized:", err)
	}
}

// GetLogger returns the private zap logger object. Call only after Init has been called.
func GetLogger() *zap.Logger {
	if logger == nil {
		stdlog.Fatal("logger not initialized")
	}
	return logger
}
