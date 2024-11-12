package Middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
)

// RequestCapture 默认打印请求体和请求头
func RequestCapture() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logger := ctx.MustGet("logger").(*zap.Logger)
		// 1. 打印请求头数
		if headerStr, err := json.Marshal(ctx.Request.Header); err != nil {
			logger.Error("读取请求头失败", zap.Error(err), zap.Any("header", ctx.Request.Header))
		} else {
			logger.Info(fmt.Sprintf("请求头: %s", headerStr))
		}
		// 2. 打印请求体数据
		bodyByte, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			logger.Error("读取请求体失败", zap.Error(err), zap.Any("body", bodyByte))
		} else {
			logger.Info(fmt.Sprintf("请求体: %s", string(bodyByte)))
		}
		// 重置请求体 以便后续读取
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyByte))

		ctx.Next()
	}
}
