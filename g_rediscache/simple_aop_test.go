package g_rediscache

import (
"fmt"
	"github.com/sirupsen/logrus"
	"reflect"
"strconv"
"testing"
"time"
)

func ExampleUseSimpleAop() {
	RedisTestSetup()
	logrus.SetLevel(logrus.ErrorLevel)

	cacheKey := "testtest_simple_" + strconv.Itoa(getRand())
	bizFunc := func() (interface{}, error) {

		return "abc", nil
	}
	v1, fromCache, _ := UseSimpleAop(cacheKey, reflect.TypeOf("")).WithExpires(5*time.Second).Then(bizFunc)
	fmt.Println(fromCache)
	v2, fromCache, _ := UseSimpleAop(cacheKey, reflect.TypeOf("")).WithExpires(5*time.Second).Then(bizFunc)
	fmt.Println(fromCache)
	fmt.Println(v1 == v2)
	fmt.Println(v1)
	// output:
	// false
	// true
	// true
	// abc
}

func TestSimpleAop(t *testing.T) {

	RedisTestSetup()
	//ctx := context.TODO()

	cacheKey := "testtest" + strconv.Itoa(getRand())

	GetRedisClient().Del(cacheKey)

	options := &SimpleOptions{}
	options.Key = cacheKey
	options.Rt = reflect.TypeOf("")
	options.Expires = 30 * time.Second

	// 第一次保证cache里面取不到值
	val, fromCache, err := SimpleAop(options, func() (interface{}, error) {
		return "123", nil
	})

	if fromCache || err != nil || val != "123" {
		t.Fatal("make sure not from cache error!", err)
	}

	// 第二次，肯定从cache里面取值
	val, fromCache, err = SimpleAop(options, func() (interface{}, error) {
		return "123", nil
	})

	if !fromCache || err != nil || val != "123" {
		t.Fatal("make sure not from cache error!", err)
	}

	// 第三次，缓存空值
	GetRedisClient().Del(cacheKey)
	options.EmptyExpires = 10 * time.Second
	val, fromCache, err = SimpleAop(options, func() (interface{}, error) {
		return "", nil
	})

	if fromCache || err != nil || val != "" {
		t.Fatal("make sure not from cache error!", err)
	}

	cacheV, _ := GetRedisClient().Get(cacheKey).Result()
	if cacheV != EmptyFlag {
		t.Error("empty cache fail")
	}

	// 第四次，空值从cache里面取出
	val, fromCache, err = SimpleAop(options, func() (interface{}, error) {
		return "123", nil
	})

	if !fromCache || err != nil || val != nil {
		t.Fatal("make sure from empty cache error!", err)
	}

}
