package main

import (
	"github.com/br3akerX/clerk"
	"github.com/br3akerX/clerk/pkg/encoder"
	"github.com/br3akerX/martini/Engine"
	"github.com/br3akerX/martini/Middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	engine := Engine.NewEngine("127.0.0.1", 7980, nil)
	engine.SetLogger(clerk.NewDefaultLogger(encoder.DefaultJsonEncoder(), os.Stdout))
	engine.Use(
		Middlewares.CancelListener(engine.GetContext()),
		Middlewares.GenerateRequestId(),
		Middlewares.SetLogger(engine.GetLogger()),
		Middlewares.PanicCapture(),
		Middlewares.RequestCapture(),
		Middlewares.ErrorCapture(),
		Middlewares.ResponseCapture(),
	)
	engine.SetHandlers([]Engine.HandlerGroup{
		{
			GroupPrefix: "",
			GroupHandlerItems: []Engine.HandlerItem{
				{
					Method: http.MethodGet,
					Path:   "/test",
					Handler: func(c *gin.Context) {
						c.JSON(200, gin.H{
							"msg": "成功",
						})
					},
				},
			},
		},
	})
	engine.Exec()
}
