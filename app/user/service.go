package user

import (
	"mfx/gin_study/model"
	"time"
)

type UserService struct {

}

func (u *UserService)Create(name string,age int,birthday time.Time)model.UserModel{
	user :=model.UserModel{
		Name: name,
		Age: age,
		Birthday: birthday,
	}
	user.Create()
	return user.LoadByName()

}