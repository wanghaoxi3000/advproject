package serializer

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 基础序列化器
type Response struct {
	Code   int         `json:"code"`
	Data   interface{} `json:"data,omitempty"`
	Result bool        `json:"result"`
	Error  string      `json:"error,omitempty"`
}

func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code:   http.StatusOK,
		Data:   data,
		Result: true,
	})
}
