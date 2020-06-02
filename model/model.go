package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

func initDatabase()  *gorm.DB {

	db, err := gorm.Open("mysql", viper.GetString("mysql.uri"))

	if err != nil {
		logrus.Fatalln("init db err:", err)
	}
	// 设置连接池
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	db.DB().SetMaxIdleConns(viper.GetInt("mysql.maxIdleConns"))

	// SetMaxOpenCons 设置数据库的最大连接数量。
	db.DB().SetMaxOpenConns(viper.GetInt("mysql.maxOpenConns"))

	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	db.DB().SetConnMaxLifetime(time.Hour)

	// 启用Logger，显示详细日志
	db.LogMode(true)
	db.SetLogger(log.New(os.Stdout, "\r\n", 0))

	// 全局禁用表名复数
	db.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响
	return db
}
func registerModels(db *gorm.DB) {
	db.Table("admin_users").AutoMigrate(&User{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Email{})
	db.AutoMigrate(&Address{})
	db.AutoMigrate(&Language{})
	db.AutoMigrate(&CreditCard{})
}
func Init()  *gorm.DB{
	db:=initDatabase()
	registerModels(db)
	return db
}
