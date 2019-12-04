package api

import (
	"advancedproject/serializer"

	"github.com/gin-gonic/gin"
)

// Status 状态输出
func Status(c *gin.Context) {
	serializer.SuccessResponse(c, "OK")
}
