package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
	"mfx/gin_study/app"
	"mfx/gin_study/model"
	"strconv"
	"time"
)

var userService UserService

// @Summary 通过用户 id 获取用户信息
// @Tags 用户模块
// @version 1.0
// @Accept application/x-www-form-urlencoded
// @Param id path int true "id"
// @Success 200 object model.UserModel 成功后返回值
// @Router  /v1/users/{id} [get]
func LoadUserById(c *gin.Context) {
	id := c.Param("id")
	basicHandle := app.BasicController{Ctx: c}
	if intId, err := strconv.Atoi(id); err != nil {
		logrus.Error(err)
		panic("translate err")
	} else {
		basicHandle.Ok(userService.LoadById(int64(intId)))
	}

}

// @Summary 创建用户
// @Tags 用户模块
// @version 1.0
// @Accept application/x-www-form-urlencoded
// @Param name query string true "name"
// @Param age query int true "age"
// @Success 200 object model.UserModel 成功后返回值
// @Router  /v1/users [post]
func CreateUser(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	ageInt, _ := strconv.Atoi(age)
	basicHandle := app.BasicController{Ctx: c}
	basicHandle.Ok(userService.Create(name, ageInt, time.Now()))
}

// @Summary 更新用户
// @Tags 用户模块
// @version 1.0
// @Accept application/x-www-form-urlencoded
// @Param id path int true "id"
// @Param age formData int false "age"
// @Param name formData string false "name"
// @Success 200 object model.UpdateUserResp 成功后返回值
// @Router  /v1/users/{id} [put]
func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	intId, _ := strconv.Atoi(id)
	// you can bind multipart form with explicit binding declaration:
	// c.ShouldBindWith(&form, binding.Form)
	// or you can simply use autobinding with ShouldBind method:
	var user model.UpdateUserReq
	// in this case proper binding will be automatically selected
	logrus.Info(user)
	if err := c.ShouldBindWith(&user, binding.FormPost); err != nil {
		logrus.Panic(err)
	}
	user.ID = uint(intId)
	basicHandle := app.BasicController{Ctx: c}

	basicHandle.Ok(userService.UpdateUserById(user))
}

func LoadUserByName(c *gin.Context) {
	res := c.FullPath() == "/v1/user/:name/*action"
	fmt.Sprintln(res) // true

	name := c.Param("name")
	action := c.Param("action") // the action will add /
	message := name + " is " + action
	basicHandle := app.BasicController{Ctx: c}
	basicHandle.Ok(message)
}

// @Summary 获取所有用户
// @Tags 用户模块
// @version 1.0
// @Accept application/x-www-form-urlencoded
// @Success 200 object []model.UpdateUserResp 成功后返回值
// @Router  /v1/users [get]
func LoadAllUsers(c *gin.Context) {
	basicHandle := app.BasicController{Ctx: c}
	basicHandle.Ok(userService.LoadAllUser())
}
