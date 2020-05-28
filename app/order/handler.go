package order

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadOrder(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "load order success",
	})
}

func CreateOrder(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "create order success",
	})
}

// Parameters in pathsss
func LoadOrderById(c *gin.Context) {
	id:=c.Param("id")
	c.String(http.StatusOK, "order id is %s", id)
}

func LoadOrderByIdAndItemId(c *gin.Context) {
	id := c.Param("id")
	itemId := c.Param("itemId") //
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")

	c.JSON(200, gin.H{
		"status":  "posted",
		"message": message,
		"nick":    nick,
		"fullPath":c.FullPath(),
		"id":id,
		"itemId":itemId,
	})
}
