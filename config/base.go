package config

import "os"

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

func GetBaseConfig() BaseConfig {
	config := defaultBaseConfig()
	for envName := range config.config {
		if "" != os.Getenv(envName) {
			config.config[envName] = os.Getenv(envName)
		}
	}
	return config
}
