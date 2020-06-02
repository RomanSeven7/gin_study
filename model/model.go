package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func initDatabase()  *gorm.DB {
	url:=viper.GetString("mysql.url")

	db, err := gorm.Open("mysql", url)
	defer db.Close()
	if err != nil {
		logrus.Error("init db err:", err)
	}
	// 全局禁用表名复数
	db.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响
	return db
}
func registerModels(db *gorm.DB) {
	db.Create(&User{})
	db.Create(&Email{})
	db.Create(&Address{})
	db.Create(&Language{})
	db.Create(&CreditCard{})
}
func Init() {
	db:=initDatabase()
	registerModels(db)
}
