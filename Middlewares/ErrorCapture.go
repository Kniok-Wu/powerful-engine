package Middlewares

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ErrorCapture 捕获异常 并自动返回
func ErrorCapture() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			logger := ctx.MustGet("logger").(*zap.Logger)
			if len(ctx.Errors) > 0 {
				fileds := make([]zap.Field, 0)
				for i := 0; i < len(ctx.Errors); i++ {
					fileds = append(fileds, zap.Error(ctx.Errors[i]))
				}
				logger.Error("请求出现异常", fileds...)
			}
		}()

		ctx.Next()
	}
}
