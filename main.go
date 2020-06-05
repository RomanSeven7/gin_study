package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"mfx/gin_study/app/order"
	"mfx/gin_study/app/user"
	"mfx/gin_study/conf"
	"mfx/gin_study/g_rediscache"
	"mfx/gin_study/model"
	"mfx/gin_study/routers"
)



func main() {
	pflag.Parse()
	_ = viper.BindPFlags(pflag.CommandLine)

	err := conf.InitConfig()
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
