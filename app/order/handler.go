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
