package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"fmt"
	"time"
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
	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %s %d %s \"%s\" %s\"\n",
			param.Request.RemoteAddr, // [::1]:63495
			param.ClientIP, // [::1]
			param.TimeStamp.Format(time.RFC3339), //"2020-06-05T22:15:30+08:00
			param.Method, //GET
			param.Path, //  /v1/user/6
			param.Request.Proto, // =HTTP/1.1
			param.StatusCode, // 200
			param.Latency, // 5.24106793s
			param.Request.UserAgent(), // PostmanRuntime/7.6.0
			param.ErrorMessage, // ""
		)
	}))
	r.Use(gin.Recovery())
	r.Use(middleware.LoggerMiddleware())
	for _, opt := range options {
		opt(r)
	}
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return r
}
