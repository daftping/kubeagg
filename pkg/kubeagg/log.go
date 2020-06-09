package kubeagg

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugar *zap.SugaredLogger

// getZapLevel returns zapcore.Level for provided --loglevel string.
// caseinsensitive
func getZapLevel(level string) zapcore.Level {
	switch {
	case strings.EqualFold(level, "Info"):
		return zapcore.InfoLevel
	case strings.EqualFold(level, "Warn"):
		return zapcore.WarnLevel
	case strings.EqualFold(level, "Debug"):
		return zapcore.DebugLevel
	case strings.EqualFold(level, "Error"):
		return zapcore.ErrorLevel
	case strings.EqualFold(level, "Fatal"):
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

// initZapLog set sugar global variable which can be used package wide for logging
func initZapLog() {
	if len(globalConfigVar.LogLevel) == 0 {
		globalConfigVar.LogLevel = "Info"
	}
	config := zap.NewDevelopmentConfig()
	config.Level = zap.NewAtomicLevelAt(getZapLevel(globalConfigVar.LogLevel))
	logger, _ := config.Build()
	defer logger.Sync()
	logger.Debug("logger construction succeeded")
	sugar = logger.Sugar()
}
