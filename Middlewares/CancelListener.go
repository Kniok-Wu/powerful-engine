package Middlewares

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CancelListener 捕获异常 并自动返回
func CancelListener(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		select {
		case <-ctx.Done():
			// 如果全局上下文被取消，则终止请求
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"error": "server is shutting down",
			})
			c.Abort()
		default:
			c.Next()
		}
	}
}
