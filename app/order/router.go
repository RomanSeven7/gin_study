package order

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	order := e.Group("/v1/order")
	{
		order.GET("", LoadOrder)
		order.POST("", CreateOrder)
		// This handler will match /user/john but will not match /user/ or /user
		order.GET("/:id", LoadOrderById)
		// load multipart/urlencoded Form
		order.POST("/:id/:itemId", LoadOrderByIdAndItemId)
		order.PUT("/:id/", UpdateOrder)
	}
}
