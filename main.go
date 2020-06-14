package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"mfx/gin_study/app/order"
	"mfx/gin_study/app/user"
<<<<<<< HEAD
	_ "mfx/gin_study/docs"
=======
	"mfx/gin_study/conf"
	"mfx/gin_study/g_rediscache"
	"mfx/gin_study/model"
>>>>>>> master
	"mfx/gin_study/routers"
)

<<<<<<< HEAD
func initConfig() (err error) {
	box := packr.New("config", "./conf")
	configType := "yml"
	defaultConfig, _ := box.Find("default.yml")
	v := viper.New()
	v.SetConfigType(configType)
	err = v.ReadConfig(bytes.NewReader(defaultConfig))
	if err != nil {
		return
	}
	configs := v.AllSettings()
	// 将default中的配置全部以默认配置写入
	for k, v := range configs {
		viper.SetDefault(k, v)
	}
	env := os.Getenv("GO_ENV")
	// 根据配置的env读取相应的配置信息
 	if env != "" {
		envConfig, _ := box.Find(env + ".yml")
		viper.SetConfigType(configType)
		err = viper.ReadConfig(bytes.NewReader(envConfig))
		if err != nil {
			return
		}
	}
	return
}
// @title Gin swagger
// @version 1.0
// @description Gin swagger 示例项目
=======

>>>>>>> master

// @contact.name
// @contact.url https://youngxhui.top
// @contact.email youngxhui@g mail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
func main() {
	pflag.Parse()
	_ = viper.BindPFlags(pflag.CommandLine)
<<<<<<< HEAD
	err := initConfig()
=======

	err := conf.InitConfig()
>>>>>>> master
	if err != nil {
		panic(err)
	}
	// 加载多个app的路由配置
	routers.Include(order.Routers, user.Routers)
	// 初始化路由
	r := routers.Init()
	defer model.Db.Close()
	// 初始化gorm
	model.Init()
	// 初始化redis
	g_rediscache.Init()
	if err := r.Run(":8089"); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}
}
