package Middlewares

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"martini/Response"
	"net/http"
)

// PanicCapture 捕获异常 并自动返回
func PanicCapture() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			logger := ctx.MustGet("logger").(*zap.Logger)
			if err := recover(); err != nil {
				logger.Error("请求异常: ", zap.Any("error", err), zap.String("requestId", ctx.GetString("requestId")))
				ctx.JSON(http.StatusOK, Response.NewStandardError())
				// 停止后续的处理
				ctx.Abort()
			}
		}()

		ctx.Next()
	}
}
