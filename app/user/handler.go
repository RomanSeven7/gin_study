package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)
var userService UserService
func LoadUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "load user success",
	})
}

func LoadUserById(c *gin.Context) {
	id := c.Param("id")
	if intId,err:=strconv.Atoi(id);err!=nil{
		logrus.Error(err)
		panic("translate err")
	}else {
		c.JSON(http.StatusOK, gin.H{
			"message": userService.LoadById(int64(intId)),
		})
	}

}
func CreateUser(c *gin.Context) {
	name:=c.Query("name")
	age:=c.Query("age")
	ageInt,_:=strconv.Atoi(age)
	c.JSON(http.StatusOK, gin.H{
		"message": userService.Create(name,ageInt,time.Now()),
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
