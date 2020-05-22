package test

import (
	"bytes"
	"fmt"
	"github.com/gobuffalo/packr/v2"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func initConfig() (err error) {
	box := packr.New("config", "../conf")
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
	fmt.Sprintln()
	return
}

func TestLoadConf(t *testing.T) {
	// todo export $GO_ENV=production
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	err := initConfig()
	if err != nil {
		panic(err)
	}
	assert.Equal(t, int64(100), viper.GetInt64("db.poolSize"))
	assert.Equal(t, int64(700), viper.GetInt64("num"))

}
