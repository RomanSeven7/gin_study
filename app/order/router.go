package order

import (
	"github.com/gin-gonic/gin"
)
func Routers(e *gin.Engine){
	 order := e.Group("/v1/order")
	{
		order.GET("", LoadOrder)
		order.POST("", CreateOrder)
	}
}