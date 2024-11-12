package Engine

import (
	"context"
	"fmt"
	"github.com/br3akerX/clerk"
	clerkEncoder "github.com/br3akerX/clerk/pkg/encoder"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Engine struct {
	engine *gin.Engine
	ctx    context.Context
	cancel context.CancelFunc
	logger *zap.Logger
	db     *gorm.DB
	ip     string
	port   int
}

type HandlerGroup struct {
	GroupPrefix       string
	GroupHandlerItems []HandlerItem
}

type HandlerItem struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

func NewEngine(ip string, port int, ctx context.Context) *Engine {
	e := &Engine{
		engine: gin.New(),
		logger: nil,
		db:     nil,
		ip:     ip,
		port:   port,
	}

	// 1. 上下文管理
	if ctx != nil {
		e.ctx, e.cancel = context.WithTimeout(ctx, 5*time.Second)
	} else {
		e.ctx, e.cancel = context.WithTimeout(context.Background(), 5*time.Second)
	}

	return e
}

func (e *Engine) SetLogger(logger *zap.Logger) *Engine {
	e.logger = logger
	return e
}

func (e *Engine) GetLogger() *zap.Logger {
	return e.logger
}

func (e *Engine) SetHandlers(gh []HandlerGroup) *Engine {
	for _, g := range gh {
		group := e.engine.Group(g.GroupPrefix)
		for _, h := range g.GroupHandlerItems {
			switch h.Method {
			case http.MethodGet:
				group.GET(h.Path, h.Handler)
			case http.MethodPost:
				group.POST(h.Path, h.Handler)
			case http.MethodPut:
				group.PUT(h.Path, h.Handler)
			case http.MethodDelete:
				group.DELETE(h.Path, h.Handler)
			case http.MethodPatch:
				group.PATCH(h.Path, h.Handler)
			case http.MethodOptions:
				group.OPTIONS(h.Path, h.Handler)
			case http.MethodHead:
				group.HEAD(h.Path, h.Handler)
			default:
				log.Fatalf("未知方法: %s\n", h.Method)
			}

			msg := fmt.Sprintf("注册路由: %s:%s%s", h.Method, g.GroupPrefix, h.Path)
			if e.logger != nil {
				e.logger.Info(msg)
			} else {
				log.Println(msg)
			}
		}
	}

	return e
}

func (e *Engine) SetDB(db *gorm.DB) *Engine {
	e.db = db
	return e
}

func (e *Engine) GetDB() *gorm.DB {
	return e.db
}

func (e *Engine) Use(middleware ...gin.HandlerFunc) *Engine {
	e.engine.Use(middleware...)
	return e
}

func (e *Engine) Exec() {
	defer e.cancel()

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", e.ip, e.port),
		Handler: e.engine,
	}

	// 1. 使用默认logger 仅输出至命令行 并添加至中间件中
	if e.logger == nil {
		e.logger = clerk.NewDefaultLogger(clerkEncoder.DefaultJsonEncoder(), os.Stdout)
		e.logger.Warn("未设置logger，将使用默认logger")
	}
	// 2. 插入中间件
	e.logger.Info(fmt.Sprintf("服务器监听地址: %s", srv.Addr))

	// 2. 启动服务器
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			e.logger.Error("服务器启动失败", zap.Error(err))
			os.Exit(1)
		}
	}()
	e.logger.Info("服务器启动成功！")

	// 3. 监听退出信号
	quitS := make(chan os.Signal, 1)
	signal.Notify(quitS, syscall.SIGINT, syscall.SIGTERM)
	<-quitS
	e.logger.Info("监听到退出信号")
	if err := srv.Shutdown(e.ctx); err != nil {
		e.logger.Error("服务器启动失败", zap.Error(err))
	}
}

func (e *Engine) GetContext() context.Context {
	return e.ctx
}

// Stop 发起engine退出信号 所有监听上下文的处理均停止
func (e *Engine) Stop() {
	e.cancel()
}

func (e *Engine) Done() <-chan struct{} {
	return e.ctx.Done()
}
