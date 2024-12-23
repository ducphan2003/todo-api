package services

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

func InitializeLogger() {
	Logger = sugarLog()
}

func sugarLog() *zap.SugaredLogger {
	config := zap.Config{
		Encoding:    "console",
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		OutputPaths: []string{"stderr"},

		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "message",
			TimeKey:      "time",
			LevelKey:     "level",
			CallerKey:    "caller",
			EncodeCaller: zapcore.FullCallerEncoder,
			EncodeLevel:  customLevelEncoder,
			EncodeTime:   syslogTimeEncoder,
		},
	}

	logger, _ := config.Build()
	return logger.Sugar()
}

func customLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

func syslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}
