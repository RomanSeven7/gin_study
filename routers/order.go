package routers

import (
	"github.com/gin-gonic/gin"
	"mfx/gin_study/controller"
)
func LoadOrder(e *gin.Engine){
	 order := e.Group("/v1/order")
	{
		order.GET("", controller.LoadOrder)
		order.POST("", controller.CreateOrder)
	}
}