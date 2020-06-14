package user

import (
	"github.com/gin-gonic/gin"
	"mfx/gin_study/app"
)



func Routers(e *gin.Engine) {
	user := e.Group("/v1/user")
	{
<<<<<<< HEAD
		order.GET("", app.Wrapper(LoadUser))
		order.POST("", app.Wrapper(CreateUser))
=======
		user.GET("", LoadUser)
		// Querystring parameters
		user.GET("/:id", LoadUserById)
		user.POST("", CreateUser)
		// will math /v1/user/  /v1/user/*
		// c.FullPath() == "/v1/user/:name/*action"
		//user.POST("/:name/*action", LoadUserByName)

>>>>>>> master
	}
}
