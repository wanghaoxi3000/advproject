package util

import (
	"advancedproject/config"
	"fmt"

	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

// BuildLogger 构造 zap logger
func BuildLogger(config config.BaseConfig) {
	var baseLogger *zap.Logger
	var err error
	if config.GetDevMode() == "develop" {
		baseLogger, err = zap.NewDevelopment()
	} else {
		baseLogger, err = zap.NewProduction()
	}
	if err != nil {
		panic(fmt.Sprintf("Build logger fail %v", err.Error()))
	}

	logger = baseLogger.Sugar()
}

// Log 返回 zap 日志对象
func Log() *zap.SugaredLogger {
	return logger
}
