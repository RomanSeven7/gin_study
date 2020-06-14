package conf

import (
	"bytes"
	"github.com/gobuffalo/packr/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func InitConfig() (err error) {
	box := packr.New("config","./" )
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
	url:=viper.GetString("mysql.uri")
	logrus.Debug(url)
	return
}
