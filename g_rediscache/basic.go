package g_rediscache

import (
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"math/rand"
	"mfx/gin_study/conf"
	"reflect"
	"time"
)

const (
	// 因为redis不能缓存空值，但是我们又会经常需要缓存空值防止频繁击穿cache
	// 因此用一个标识存储空的值
	EmptyFlag      = "###+--**-+###"
	defaultTimeout = 5 * time.Second
	defaultExpire  = 30 * time.Second
)

func init() {

}

type Options struct {
	// 非空
	Key string
	// 非空, 如果是集合类型，表示集合里面元素的类型
	Rt           reflect.Type
	Expires      time.Duration
	EmptyExpires time.Duration
}

func (o *Options) validate() error {
	if o.Key == "" {
		return errors.New("Key must not be empty!")
	}
	if o.Rt == nil {
		return errors.New("Rt must not be empty!")
	}
	return nil
}

// 返回的三个参数，依次是: cache值，是否空，错误信息
func GetCacheValueItem(v interface{}) (string, bool, error) {
	jsonB, err := json.Marshal(v)
	if err != nil {
		return "", true, err
	}
	cacheV := string(jsonB)
	cLength := 0
	switch v.(type) {
	// string 类型单独处理
	case string:
		cLength = len(v.(string))
	default:
		cLength = len(cacheV)
	}
	return cacheV, cLength == 0, nil
}

func getRand() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(10000000)
}

func RedisTestSetup() {
	if err := conf.InitConfig();err!=nil{
		logrus.Error(err)
	}
	BuildRedisClient()
}