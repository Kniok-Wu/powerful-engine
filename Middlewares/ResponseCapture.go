package Middlewares

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
)

type ResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// Write 重写response 同时写入响应体以及日志byte
func (rw *ResponseWriter) Write(b []byte) (int, error) {
	log.Println("写入: ", string(b))
	rw.body.Write(b)
	return rw.ResponseWriter.Write(b)
}

// ResponseCapture 重写response的write 方便日志输出
func ResponseCapture() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logger := ctx.MustGet("logger").(*zap.Logger)

		// 1. 重写返回体对象
		rw := &ResponseWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: ctx.Writer,
		}
		ctx.Writer = rw

		ctx.Next()

		// 2. 打印返回体数据
		logger.Info(fmt.Sprintf("响应头为: %s", rw.Header()))
		logger.Info(fmt.Sprintf("响应体为: %s", rw.body.String()))
	}
}
