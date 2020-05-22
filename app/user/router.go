package user

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine){
	order := e.Group("/v1/user")
	{
		order.GET("", LoadUser)
		order.POST("", CreateUser)
	}
}