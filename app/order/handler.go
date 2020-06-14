package order

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mfx/gin_study/app"
)


func LoadOrder(c *gin.Context)  {
	basicHandle := app.BasicController{Ctx: c}
	basicHandle.Ok("load order success")
	//return nil
}

func CreateOrder(c *gin.Context)  {
	//c.JSON(http.StatusOK, gin.H{
	//	"message": "create order success",
	//})
	basicHandle := app.BasicController{Ctx: c}
	basicHandle.Ok("create order success")
	//return nil
}

// Parameters in path
func LoadOrderById(c *gin.Context) {
	basicHandle := app.BasicController{Ctx: c}
	id := c.Param("id")
	basicHandle.Ok(fmt.Sprintf("order id is %s", id))
}

func LoadOrderByIdAndItemId(c *gin.Context) {
	id := c.Param("id")
	itemId := c.Param("itemId") //
	message := c.PostForm("message")
	vn := c.DefaultQuery("vn", "1.1")
	pkg := c.Query("pkg")
	nick := c.DefaultPostForm("nick", "anonymous")
	basicHandle := app.BasicController{Ctx: c}
	mapV := map[string]string{
		"status":   "posted",
		"message":  message,
		"nick":     nick,
		"fullPath": c.FullPath(),
		"id":       id,
		"itemId":   itemId,
		"vn":       vn,
		"pkg":      pkg,
	}
	basicHandle.Ok(mapV)
}

// Map/array as querystring or postform parameters
func UpdateOrder(c *gin.Context) {
	basicHandle := app.BasicController{Ctx: c}
	idMap := c.QueryMap("idMap")
	idArr := c.QueryArray("idArr")
	nameArr := c.PostFormArray("nameArr")
	nameMap := c.PostFormMap("nameMap")
	mapV := map[string]interface{}{
		"idMap":   idMap,
		"idArr":   idArr,
		"nameArr": nameArr,
		"nameMap": nameMap,
	}
	basicHandle.Ok(mapV)
}
