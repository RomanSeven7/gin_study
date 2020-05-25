package app

import (
	"github.com/gin-gonic/gin"
	"mfx/gin_study/model"
	"net/http"
	"time"
)

type BasicController struct {
	Ctx *gin.Context
}

type HandlerFunc func(c *gin.Context) error

// 统一返回值
type ResponseData struct {
	Ret        int         `json:"ret"`
	Message    string      `json:"message"`
	Result     interface{} `json:"result"`
	ServerTime int64       `json:"serverTime"`
}

func (t *BasicController) Ok(d interface{}) {
	rd := &ResponseData{
		Ret:        200,
		Message:    "ok",
		Result:     d,
		ServerTime: time.Now().UnixNano() / 1000000,
	}
	t.Ctx.JSONP(http.StatusOK, rd)
	return
}

func Wrapper(handler func(c *gin.Context) error) (func(c *gin.Context)) {
	return func(c *gin.Context) {
		var (
			err error
		)
		handler(c)
		if err != nil {
			var apiException *model.APIException
			if h,ok := err.(*model.APIException); ok {
				apiException = h
			}else {
				apiException = model.ServerError()
			}
			apiException.Request = c.Request.Method + " "+ c.Request.URL.String()
			c.JSON(apiException.Code,apiException)
			return
		}

	}
}
