package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"mfx/gin_study/model"
	"net/http"
)

func LoadUser(c *gin.Context) {
	var commonParams model.CommonParams
	if err:=c.ShouldBind(&commonParams);err!=nil{
		logrus.Error(err)
	}


	c.JSON(http.StatusOK, gin.H{
		"message": "load user success",
		"commonParams":commonParams,
	})
}

func LoadUserById(c *gin.Context) {
	id := c.Param("id")
	firstName := c.DefaultQuery("firstName", "Guest")
	lastName := c.Query("lastName") // shortcut for c.Request.URL.Query().Get("lastname")
	c.JSON(http.StatusOK, gin.H{
		"message": struct {
			Id        string
			FirstName string
			LastName  string
		}{
			Id:        id,
			FirstName: firstName,
			LastName:  lastName,
		},
	})
}
func CreateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "create user success",
	})
}

func LoadUserByName(c *gin.Context) {
	res := c.FullPath() == "/v1/user/:name/*action"
	fmt.Sprintln(res) // true

	name := c.Param("name")
	action := c.Param("action") // the action will add /
	message := name + " is " + action
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
