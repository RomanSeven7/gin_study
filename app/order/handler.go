package order

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func LoadOrder(c *gin.Context) {
	id := c.GetInt("id")
	log.Debug(id)
	c.JSON(http.StatusOK, gin.H{
		"message": "load order success",
	})
}

func CreateOrder(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "create order success",
	})
}

// Parameters in path
func LoadOrderById(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, "order id is %s", id)
}

func LoadOrderByIdAndItemId(c *gin.Context) {
	id := c.Param("id")
	itemId := c.Param("itemId") //
	message := c.PostForm("message")
	vn := c.DefaultQuery("vn", "1.1")
	pkg := c.Query("pkg")
	nick := c.DefaultPostForm("nick", "anonymous")
	c.JSON(200, gin.H{
		"status":   "posted",
		"message":  message,
		"nick":     nick,
		"fullPath": c.FullPath(),
		"id":       id,
		"itemId":   itemId,
		"vn":       vn,
		"pkg":      pkg,
	})
}


// Map/array as querystring or postform parameters
func UpdateOrder(c *gin.Context) {
	idMap := c.QueryMap("idMap")
	idArr := c.QueryArray("idArr")
	nameArr := c.PostFormArray("nameArr")
	nameMap := c.PostFormMap("nameMap")
	c.JSON(200, gin.H{
		"idMap":   idMap,
		"idArr":   idArr,
		"nameArr": nameArr,
		"nameMap": nameMap,
	})
}
