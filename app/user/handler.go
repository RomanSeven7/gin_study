package user

import (
	"github.com/gin-gonic/gin"
	"mfx/gin_study/model"
	"net/http"
)

func LoadUser(c *gin.Context) error{
	c.JSON(http.StatusOK, gin.H{
		"message": "load user success",
	})
	return nil
}

func CreateUser(c *gin.Context) error{
	name:=c.GetString("name")
	if name == "" {
		return model.ServerError()
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "create user success",
	})
	return nil
}
