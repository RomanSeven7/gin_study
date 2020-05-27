package order

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func LoadOrder(c *gin.Context) {
	id := c.GetInt("id")
	log.Debug(id)
	log.Warn(id)
	log.Info(id)
	c.JSON(http.StatusOK, gin.H{
		"message": "load order success",
	})
}

func CreateOrder(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "create order success",
	})
}
