package logging

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log = newLogger()

func newLogger() *zap.Logger {
	level, err := zapcore.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		panic(err)
	}

	return zap.New(zapcore.NewCore(zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		NameKey:        "logger",
		StacktraceKey:  "st",
		CallerKey:      "caller",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.NanosDurationEncoder,
	}), os.Stdout, level))
}

func GetLogger() *zap.Logger {
	return log
}
