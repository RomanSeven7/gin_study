package user

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"mfx/gin_study/g_rediscache"
	"mfx/gin_study/model"
	"reflect"
	"time"
)

type UserService struct {
}

func (u *UserService) Create(name string, age int, birthday time.Time) model.UserModel {
	user := model.UserModel{
		Name:     name,
		Age:      age,
		Birthday: birthday,
	}
	user.Create()
	return user.LoadByName()
}

func (u *UserService) LoadById(uid int64) model.UserModel {
	var err error
	user := model.UserModel{
	}
	user.ID = uint(uid)
	val, isCache, _ := g_rediscache.UseSimpleAop(fmt.Sprintf("load_by_userId:%d", uid), reflect.TypeOf(model.UserModel{})).
		WithExpires(time.Hour * 1).
		Then(func() (interface{}, error) {
			if user, err = user.LoadById(); err != nil {
				logrus.Error(err)
				panic("LoadById error")
			}
			return user, nil
		})
	logrus.Info("val：", val)
	logrus.Info("isCache:	", isCache)
	return val.(model.UserModel)
}

func (u *UserService) UpdateUserById(req model.UpdateUserReq) model.UpdateUserResp {

	userModel := req.UpdateById()
	refreshUserRDB(userModel.ID)
	return model.UpdateUserResp{
		ID:       userModel.ID,
		Birthday: userModel.Birthday,
		Age:      userModel.Age,
		Name:     userModel.Name,
	}
}

func (u *UserService) LoadAllUser() []model.UserModel {
	user := model.UserModel{
	}
	return user.LoadAllUsers()
}

func (u *UserService) DeleteUser(uid int) interface{} {
	user := model.UserModel{}
	user.ID = uint(uid)
	user.DeleteById()
	refreshUserRDB(user.ID)
	return nil
}

func refreshUserRDB(userId uint){
	_, _ = g_rediscache.RedisClient.Del(fmt.Sprintf("load_by_userId:%d", userId)).Result()
}