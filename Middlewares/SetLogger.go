package Middlewares

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// SetLogger 设置一个携带默认requestId的logger
func SetLogger(l *zap.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logger := l.With(zap.String("requestId", ctx.GetString("requestId")))
		ctx.Set("logger", logger)
		logger.Info("logger设置成功")
		ctx.Next()
	}
}
