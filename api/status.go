package api

import (
	"advancedproject/serializer"

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
