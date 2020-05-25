package order

import (
	"github.com/gin-gonic/gin"
	"mfx/gin_study/app"
	"net/http"
)

//type OrderController struct {
//	app.BasicController
//}

// @Summary 通过文章 id 获取单个文章内容
// @version 1.0
// @Accept application/x-json-stream
// @Param id path int true "id"
// @Success 200 object model.Result 成功后返回值
// @Router /article/{id} [get]
func LoadOrder(c *gin.Context) error {
	basicHandle := app.BasicController{Ctx: c}
	basicHandle.Ok("load order success")
	return nil
}

func CreateOrder(c *gin.Context)error {
	c.JSON(http.StatusOK, gin.H{
		"message": "create order success",
	})
	return nil
}
