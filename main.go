package main

import (
	"bytes"
	"fmt"
	"github.com/gobuffalo/packr/v2"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"mfx/gin_study/app/order"
	"mfx/gin_study/app/user"
	"mfx/gin_study/model"
	"mfx/gin_study/routers"
	"os"
)

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

func main() {
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	err := initConfig()
	if err != nil {
		panic(err)
	}
	// 加载多个app的路由配置
	routers.Include(order.Routers, user.Routers)
	// 初始化路由
	r := routers.Init()
	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
	model.Init()
}
