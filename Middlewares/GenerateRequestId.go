package Middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GenerateRequestId 创建一个requestId方便追踪
func GenerateRequestId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 1. 获取 or 生成返回体
		requestId := ctx.Request.Header.Get("requestId")
		if requestId == "" {
			requestId = uuid.New().String()
		}
		// 2. 将requestId写入返回体中
		ctx.Writer.Header().Set("requestId", requestId)
		ctx.Set("requestId", requestId)
		ctx.Next()
	}
}
