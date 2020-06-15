package model

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"time"
)

type UserModel struct {
	gorm.Model
	Birthday time.Time // 生日
	Age      int       // 年龄
	Name     string    `gorm:"size:100"` // 姓名
}

type UpdateUserReq struct {
	ID   uint   `form:"id"`                      // Id
	Age  int    `form:"age" binding:"required"`  // 年龄
	Name string `form:"name" binding:"required"` // 姓名
}

type UpdateUserResp struct {
	ID       uint      `json:"id"`       // Id
	Birthday time.Time `json:"birthday"` // 生日
	Age      int       `json:"age"`      // 年龄
	Name     string    `json:"name"`     // 姓名
}

func (user UserModel) TableName() string {
	if user.Name == "admin" {
		return "admin_users"
	} else {
		return "users"
	}
}

func (user *UserModel) Create() {
	if err := Db.Create(user).Error; err != nil {
		logrus.Error("user create err:", err)
	}
}

func (user *UserModel) LoadById() (UserModel, error) {
	err := Db.Where("id = ?", user.ID).First(&user).Error
	return *user, err
}

func (user *UserModel) LoadByName() UserModel {
	Db.Where("name = ?", user.Name).First(&user)
	return *user
}

// 使用 struct 更新多个属性，只会更新其中有变化且为非零值的字段
func (user *UpdateUserReq) UpdateById() UserModel {

	userModel := UserModel{}
	userModel.ID = user.ID
	_ = Db.Model(&userModel).Updates(UserModel{Name: user.Name, Age: user.Age}).Value

	return userModel

}

// 查询所有的记录
func (user *UserModel) LoadAllUsers() []UserModel {
	userModelList := []UserModel{}
	Db.Find(&userModelList)
	return userModelList
}

// 根据id 删除用户信息
func (user *UserModel)DeleteById() {
	// 删除现有记录
	value:=Db.Delete(&user).Value
	//// DELETE from user where id=10;
	logrus.Debug(value)
}
