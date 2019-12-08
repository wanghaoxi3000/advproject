package config

import (
	"github.com/joho/godotenv"
)

// Init 加载环境变量配置
func init() {
	godotenv.Load()
}
