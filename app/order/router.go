package order

import (
	"github.com/gin-gonic/gin"
	"mfx/gin_study/app"
)

func Routers(e *gin.Engine) {
	order := e.Group("/v1/order")
	{
		order.GET("", app.Wrapper(LoadOrder))
		order.POST("", app.Wrapper(CreateOrder))
		// This handler will match /user/john but will not match /user/ or /user
		order.GET("/:id", LoadOrderById)
		// load multipart/urlencoded Form
		order.POST("/:id/:itemId", LoadOrderByIdAndItemId)
		order.PUT("/:id/", UpdateOrder)
	}
}
