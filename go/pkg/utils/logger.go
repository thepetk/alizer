package utils

import (
	"fmt"

	"github.com/go-logr/logr"
	"go.uber.org/zap/zapcore"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

type CLILogger struct {
	Logger    logr.Logger
	Activated bool
}

var CliLogger CLILogger

func getZapcoreLevel(level string) (zapcore.Level, error) {
	switch level {
	case "debug":
		return zapcore.DebugLevel, nil
	case "warning":
		return zapcore.WarnLevel, nil
	case "info":
		return zapcore.InfoLevel, nil
	case "":
		return zapcore.ErrorLevel, nil
	default:
		return zapcore.ErrorLevel, fmt.Errorf("log level %s does not exist", level)
	}
}

func GetOrCreateLogger() logr.Logger {
	if !CliLogger.Activated {
		GenLogger("")
	}
	return CliLogger.Logger
}

func GenLogger(logLevel string) error {
	level, err := getZapcoreLevel(logLevel)
	if err != nil {
		return err
	}
	CliLogger = CLILogger{
		Logger: zap.New(zap.UseFlagOptions(&zap.Options{
			Development: true,
			TimeEncoder: zapcore.ISO8601TimeEncoder,
			Level:       level,
		})),
		Activated: true,
	}
	return nil
}
