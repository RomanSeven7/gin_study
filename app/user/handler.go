package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "load user success",
	})
}

func LoadUserById(c *gin.Context) {
	id := c.Param("id")
	firstName := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")
	c.JSON(http.StatusOK, gin.H{
		"message": struct {
			Id        string
			FirstName string
			LastName  string
		}{
			Id:        id,
			FirstName: firstName,
			LastName:  lastname,
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
