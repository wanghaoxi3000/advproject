package api

import (
	"advancedproject/serializer"
	"os"

	"github.com/gin-gonic/gin"
)

// Status 状态输出
// @Summary 状态输出
// @Tags status
// @Success 200 {object} serializer.Response 成功后返回值
// @Router /status [get]
func Status(c *gin.Context) {
	serializer.SuccessResponse(c, "OK")
}

// Hostname 主机标识
// @Summary 主机标识
// @Tags status
// @Success 200 {object} serializer.Response 成功后返回值
// @Router /hostname [get]
func Hostname(c *gin.Context) {
	hostname := os.Getenv("HOSTNAME")
	serializer.SuccessResponse(c, hostname)
}

// TODO: 状态输出控制
