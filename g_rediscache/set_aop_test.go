package g_rediscache

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func ExampleUseSetAop() {
	RedisTestSetup()
	cacheKey := "testtest_set_" + strconv.Itoa(getRand())
	bizFunc := func() ([]interface{}, error) {
		var result []interface{}
		result = append(result, User{Id: 1}, User{Id: 2})
		return result, nil
	}
	_, fromCache, _ := UseSetAop(cacheKey, reflect.TypeOf(User{})).WithExpires(5*time.Second).Then(bizFunc)
	fmt.Println(fromCache)
	_, fromCache, _ = UseSetAop(cacheKey, reflect.TypeOf(User{})).WithExpires(5*time.Second).Then(bizFunc)
	fmt.Println(fromCache)
	// output:
	// false
	// true
}

func TestSetAop(t *testing.T) {
	RedisTestSetup()

	cacheKey := "testtest_set_" + strconv.Itoa(getRand())

	options := &SetOptions{}
	options.Key = cacheKey
	options.Rt = reflect.TypeOf("")
	options.Expires = 30 * time.Second

	GetRedisClient().Del(cacheKey)

	// 第一次保证不从cache里面取值
	val, fromCache, err := SetAop(options, func() ([]interface{}, error) {
		var result []interface{}
		result = append(result, "11", "22")
		return result, nil
	})

	if err != nil || fromCache || len(val) != 2 || val[0] != "11" || val[1] != "22" {
		t.Fatal("1. must not be from cache FAIL")
	}

	// 第二次保证从cache里面取值
	val, fromCache, err = SetAop(options, func() ([]interface{}, error) {
		var result []interface{}
		result = append(result, "11", "22")
		return result, nil
	})

	if err != nil || !fromCache || len(val) != 2 || !(val[0] != "11" || val[1] != "11") || !(val[0] != "22" || val[1] != "22") {
		t.Fatal("2. must not be from cache FAIL")
	}

	// 第三次保证保证存进去EmptyFlag
	GetRedisClient().Del(cacheKey)
	options.EmptyExpires = 30 * time.Second
	val, fromCache, err = SetAop(options, func() ([]interface{}, error) {
		return nil, nil
	})

	if err != nil || fromCache || val != nil {
		t.Fatal("3. must be empty FAIL")
	}
	vals, _ := GetRedisClient().SMembers(cacheKey).Result()
	if len(vals) != 1 && vals[0] != EmptyFlag {
		t.Fatal("3-1. must be empty FAIL")
	}

}
