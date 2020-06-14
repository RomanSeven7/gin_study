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
var Db *gorm.DB
func initDatabase()   {
	var err error
	Db, err = gorm.Open("mysql", viper.GetString("mysql.uri"))

	if err != nil {
		logrus.Fatalln("init db err:", err)
	}
	// 设置连接池
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	Db.DB().SetMaxIdleConns(viper.GetInt("mysql.maxIdleConns"))

	// SetMaxOpenCons 设置数据库的最大连接数量。
	Db.DB().SetMaxOpenConns(viper.GetInt("mysql.maxOpenConns"))

	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	Db.DB().SetConnMaxLifetime(time.Hour)

	// 启用Logger，显示详细日志
	Db.LogMode(true)
	Db.SetLogger(log.New(os.Stdout, "\r\n", 0))

	// 全局禁用表名复数
	Db.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响
}
func registerModels() {
	Db.Table("admin_users").AutoMigrate(&UserModel{})
	Db.Table("users").AutoMigrate(&UserModel{})
}
func Init() {
	initDatabase()
	registerModels()
}
