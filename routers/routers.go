package routers

import (
	"github.com/gin-gonic/gin"
	"mfx/gin_study/middleware"
)

type Option func(*gin.Engine)

var options = []Option{}

// 注册app的路由配置
func Include(opts ...Option) {
	options = append(options, opts...)
}

// 初始化
func Init() *gin.Engine {
	r := gin.New()
	r.Use(middleware.LogerMiddleware())
	for _, opt := range options {
		opt(r)
	}
	return r
}
