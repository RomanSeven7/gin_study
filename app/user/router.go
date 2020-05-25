package user

import (
	"github.com/gin-gonic/gin"
	"mfx/gin_study/app"
)



func Routers(e *gin.Engine) {
	order := e.Group("/v1/user")
	{
		order.GET("", app.Wrapper(LoadUser))
		order.POST("", app.Wrapper(CreateUser))
	}
}
