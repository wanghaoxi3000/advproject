package server

import (
	"advancedproject/api"

	"github.com/gin-gonic/gin"
)

// SetupRouter 路由配置
func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/status/", api.Status)
	}

	return r
}
