package routers

import (
	"github.com/gin-gonic/gin"
	"mfx/gin_study/controller"
)

func LoadUser(e *gin.Engine){
	order := e.Group("/v1/user")
	{
		order.GET("", controller.LoadUser)
		order.POST("", controller.CreateUser)
	}
}