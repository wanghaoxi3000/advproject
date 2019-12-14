package config

import "os"

// BaseConfig 项目基础配置
type BaseConfig interface {
	GetDevMode() string
}

type envBaseConfig struct {
	config map[string]string
}

func (c *envBaseConfig) GetDevMode() string {
	return c.config["RUN_MODE"]
}

func defaultBaseConfig() *envBaseConfig {
	return &envBaseConfig{
		config: map[string]string{
			"RUN_MODE": "develop",
		},
	}
}

// GetBaseConfig 获取基础配置Map
func GetBaseConfig() BaseConfig {
	config := defaultBaseConfig()
	for envName := range config.config {
		if "" != os.Getenv(envName) {
			config.config[envName] = os.Getenv(envName)
		}
	}
	return config
}
