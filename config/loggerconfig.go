package config

import (
	"os"
	"strings"
)

type LoggerConfig interface {
	GetLoggerLevel() string
	GetLoggerEncoding() string
	GetOutputPaths() []string
	GetErrorOutputPaths() []string
}

type envLoggerConfig struct {
	config map[string]string
}

// GetLoggerLevel 激活
func (c *envLoggerConfig) GetLoggerLevel() string {
	return c.config["LOG_LEVEL"]
}

// GetLoggerEncoding host 主机名
func (c *envLoggerConfig) GetLoggerEncoding() string {
	return c.config["LOG_ENCODING"]
}

// GetOutputPaths 连接端口
func (c *envLoggerConfig) GetOutputPaths() []string {
	return strings.Split(c.config["LOG_OUTPUT_PATHS"], ",")
}

// ErrorOutputPaths 数据库名称
func (c *envLoggerConfig) GetErrorOutputPaths() []string {
	return strings.Split(c.config["LOG_ERROUTPUT_PATHS"], ",")
}

func defaultEnvLoggerConfig() *envLoggerConfig {
	return &envLoggerConfig{
		config: map[string]string{
			"LOG_LEVEL":           "info",
			"LOG_ENCODING":        "json",
			"LOG_OUTPUT_PATHS":    "stdout",
			"LOG_ERROUTPUT_PATHS": "stderr",
		},
	}
}

func GetLoggerConfig() LoggerConfig {
	config := defaultEnvLoggerConfig()
	for envName := range config.config {
		if "" != os.Getenv(envName) {
			config.config[envName] = os.Getenv(envName)
		}
	}
	return config
}
