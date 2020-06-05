package model

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"time"
)

type UserModel struct {
	gorm.Model
	Birthday time.Time
	Age      int
	Name     string `gorm:"size:100"` // string默认长度为255, 使用这种tag重设。
}

func (user UserModel) TableName() string {
	if user.Name == "admin" {
		return "admin_users"
	} else {
		return "users"
	}
}

func (user *UserModel) Create() {
	if err:=Db.Create(user).Error;err!=nil{
		logrus.Error("user create err:",err)
	}
}

func (user *UserModel) LoadById() (UserModel,error) {
	err:=Db.Where("id = ?", user.ID).First(&user).Error
	return *user,err
}

func (user *UserModel) LoadByName() UserModel {
	Db.Where("name = ?", user.Name).First(&user)
	return *user
}