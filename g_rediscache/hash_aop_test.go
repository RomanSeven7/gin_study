package g_rediscache

import (
	"fmt"
	"github.com/astaxie/beego"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func ExampleUseHashAop() {
	RedisTestSetup()
	beego.SetLevel(beego.LevelError)
	cacheKey := "testtest_hash_" + strconv.Itoa(getRand())
	bizFunc := func() (i []interface{}, e error) {
		var r []interface{}
		r = append(r, User{Id: 1, Name: "name1"})
		r = append(r, User{Id: 2, Name: "name2"})
		return r, nil
	}
	_, fromCaChe, _ := UseHashAop(cacheKey, reflect.TypeOf(User{}), []string{"1", "2"}, "Id").WithExpires(5 * time.Second).Then(bizFunc)
	fmt.Println(fromCaChe)
	_, fromCaChe, _ = UseHashAop(cacheKey, reflect.TypeOf(User{}), []string{"1", "2"}, "Id").WithExpires(5 * time.Second).Then(bizFunc)
	fmt.Println(fromCaChe)
	// output:
	// false
	// true
}

func TestHashAop(t *testing.T) {
	RedisTestSetup()
	cacheKey := "testtest_hash_" + strconv.Itoa(getRand())

	GetRedisClient().Del(cacheKey)

	options := &HashOptions{}
	options.Key = cacheKey
	options.Rt = reflect.TypeOf(User{})
	options.Expires = 5 * time.Second
	options.Fields = []string{"1", "2"}
	options.FieldAttr = "Id"

	// 第一次, cache里没有值，从fallback取到并回填
	val, fromCache, err := HashAop(options, func() ([]interface{}, error) {
		var r []interface{}
		r = append(r, User{Id: 1, Name: "name1"})
		r = append(r, User{Id: 2, Name: "name2"})
		return r, nil
	})
	u1 := val[0].(User)
	u2 := val[1].(User)
	if err != nil || fromCache || u1.Id != 1 || u2.Id != 2 {
		t.Fatal("1. must not be from cache FAIL")
	}
	// 查看回填情况
	cacheVs, err := GetRedisClient().HMGet(cacheKey, "1", "2").Result()
	if err != nil || cacheVs[0] == nil || cacheVs[1] == nil {
		t.Fatal("1.1 rewrite to cache FAIL")
	}

	// 第二次, cache有值，直接从cache里取值
	val, fromCache, err = HashAop(options, func() ([]interface{}, error) {
		var r []interface{}
		r = append(r, User{Id: 1, Name: "name1"})
		r = append(r, User{Id: 2, Name: "name2"})
		return r, nil
	})
	u1 = val[0].(User)
	u2 = val[1].(User)
	if err != nil || !fromCache || u1.Id != 1 || u2.Id != 2 {
		t.Fatal("2. must be from cache FAIL")
	}

	time.Sleep(5 * time.Second)
	cacheVs, err = GetRedisClient().HMGet(cacheKey, "1", "2").Result()
	if err != nil || cacheVs[0] != nil || cacheVs[1] != nil {
		t.Fatal("3. expires FAIL")
	}

}
