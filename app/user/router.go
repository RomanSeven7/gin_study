package user

import (
	"github.com/gin-gonic/gin"
	"mfx/gin_study/app"
)



func Routers(e *gin.Engine) {
	user := e.Group("/v1/users")
	{
		user.POST("", app.Wrapper(CreateUser))
		// Querystring parameters
		user.GET("/:id", app.Wrapper(LoadUserById))
		user.PUT("/:id", app.Wrapper(UpdateUser))
		user.GET("/", app.Wrapper(LoadAllUsers))
		user.DELETE("/:id", app.Wrapper(DeleteUser))

	}
}
