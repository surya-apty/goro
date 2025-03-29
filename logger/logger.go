package logger

import (
	"go.uber.org/zap"
)

type Logger struct {
	*zap.Logger
}

func InitLogger() *Logger {
	var err error
	zapLog, err := zap.NewProduction()
	if err != nil {
		panic("failed to initialize zap logger: " + err.Error())
	}
	return &Logger{Logger: zapLog}
}
