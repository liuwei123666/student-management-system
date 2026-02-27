package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorResponse 错误响应结构
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
	Status  int    `json:"status"`
}

// ErrorHandler 统一错误处理中间件
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 处理请求
		c.Next()

		// 检查是否有错误
		if len(c.Errors) > 0 {
			// 获取最后一个错误
			err := c.Errors.Last()
			log.Printf("Error: %v", err.Error())

			// 根据错误类型返回不同的状态码
			statusCode := http.StatusInternalServerError
			message := "Internal server error"

			// 检查是否有自定义状态码
			if c.Writer.Status() != http.StatusOK {
				statusCode = c.Writer.Status()
			}

			// 返回错误响应
			c.JSON(statusCode, ErrorResponse{
				Error:   err.Error(),
				Message: message,
				Status:  statusCode,
			})

			// 中止后续处理
			c.Abort()
		}
	}
}
