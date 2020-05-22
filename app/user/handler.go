package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "load user success",
	})
}

func CreateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "create user success",
	})
}
