package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)


func LogerMiddleware() gin.HandlerFunc {
	// 实例化
	logger := logrus.New()
	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{})
	//logrus.SetReportCaller(true)
	logrus.SetOutput(os.Stderr)
	logrus.SetLevel(logrus.DebugLevel)
	//rotateFileHook, _ := NewRotateFileHook(*Generate(logrus.DebugLevel))
	//rotateFileHook2, _ := NewRotateFileHook(*Generate(logrus.InfoLevel))
	//rotateFileHook3, _ := NewRotateFileHook(*Generate(logrus.WarnLevel))
	//rotateFileHook4, _ := NewRotateFileHook(*Generate(logrus.ErrorLevel))
	//logrus.AddHook(rotateFileHook)
	//logrus.AddHook(rotateFileHook2)
	//logrus.AddHook(rotateFileHook3)
	//logrus.AddHook(rotateFileHook4)
	return func(c *gin.Context) {
		//开始时间
		startTime := time.Now()
		//处理请求
		c.Next()
		//结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		//请求方式
		reqMethod := c.Request.Method
		//请求路由
		reqUrl := c.Request.RequestURI
		//状态码
		statusCode := c.Writer.Status()
		//请求ip
		clientIP := c.ClientIP()

		// 日志格式
		logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUrl,
		}).Info()

	}
}


func Generate(level logrus.Level)*RotateFileConfig {
	rotateFileConfig:=  RotateFileConfig{
		Filename: fmt.Sprintf("../../log/%s-%s.log",level.String(),time.Now().Format(time.RFC3339)),
		MaxSize: 10,
		MaxBackups: 10,
		MaxAge: 7,
		Level: level,
		Formatter: &logrus.JSONFormatter{},
	}
	writer, _ := rotatelogs.New(
		rotateFileConfig.Filename+".%Y-%m-%d-%H-%M",
		rotatelogs.WithLinkName(rotateFileConfig.Filename),
		rotatelogs.WithMaxAge(time.Duration(2)*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(5)*time.Second),
	)
	logrus.SetOutput(writer)
	return &rotateFileConfig
}