package g_rediscache

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func ExampleUseZSetAop() {
	RedisTestSetup()

	cacheKey := "testtest_zset_" + strconv.Itoa(getRand())
	bizFunc := func() (interface{}, error) {
		var r []interface{}
		r = append(r, User{Id: 1, Name: "name1"})
		r = append(r, User{Id: 2, Name: "name2"})
		return r, nil
	}
	_, fromCache, _ := UseZSetAop(cacheKey, reflect.TypeOf(User{})).WithScoreField("Id").WithStart(0).WithStop(-1).WithExpires(5*time.Second).Then(bizFunc)
	fmt.Println(fromCache)
	_, fromCache, _ = UseZSetAop(cacheKey, reflect.TypeOf(User{})).WithScoreField("Id").WithStart(0).WithStop(-1).WithExpires(5*time.Second).Then(bizFunc)
	fmt.Println(fromCache)
	// output:
	// false
	// true
}

func TestZSetAop(t *testing.T) {
	RedisTestSetup()


	cacheKey := "testtest_zset_" + strconv.Itoa(getRand())

	GetRedisClient().Del(cacheKey)

	options := &ZSetOptions{}
	options.Key = cacheKey
	options.Rt = reflect.TypeOf("")
	options.Start = 0
	options.Stop = -1
	options.Expires = 30 * time.Second

	// 第一次, map 类型 cache里没有值，从fallback取到并回填
	options.IsMap = true
	val, fromCache, err := ZSetAop(options, func() (interface{}, error) {
		r := make(map[interface{}]float64)
		r["1"] = 1
		r["2"] = 2
		return r, nil
	})
	valm := val.(map[interface{}]float64)
	if err != nil || fromCache || valm["1"] != 1 || valm["2"] != 2 {
		t.Fatal("1. must not be from cache FAIL")
	}
	// 第二次, map 类型 cache里有值
	val, fromCache, err = ZSetAop(options, func() (interface{}, error) {
		r := make(map[interface{}]float64)
		r["1"] = 1
		r["2"] = 2
		return r, nil
	})
	valm = val.(map[interface{}]float64)
	if err != nil || !fromCache || valm["1"] != 1 || valm["2"] != 2 {
		t.Fatal("2. must be from cache FAIL")
	}

	GetRedisClient().Del(cacheKey)
	// 第三次, 非map 类型 cache没值
	options.IsMap = false
	options.ScoreField = "Id"
	options.Rt = reflect.TypeOf(User{})
	val, fromCache, err = ZSetAop(options, func() (interface{}, error) {
		var r []interface{}
		r = append(r, User{Id: 1, Name: "name1"})
		r = append(r, User{Id: 2, Name: "name2"})
		return r, nil
	})
	valarr := val.([]interface{})
	u1 := valarr[0].(User)
	u2 := valarr[1].(User)
	if err != nil || fromCache || len(valarr) != 2 || u1.Name != "name1" || u1.Id != 1 || u2.Name != "name2" || u2.Id != 2 {
		t.Fatal("3. must not be from cache FAIL")
	}

	// 第四次, 非map 类型 cache里有值
	val, fromCache, err = ZSetAop(options, func() (interface{}, error) {
		var r []interface{}
		r = append(r, User{Id: 1, Name: "name1"})
		r = append(r, User{Id: 2, Name: "name2"})
		return r, nil
	})
	valarr = val.([]interface{})
	u1 = valarr[0].(User)
	u2 = valarr[1].(User)
	if err != nil || !fromCache || len(valarr) != 2 || u1.Name != "name1" || u1.Id != 1 || u2.Name != "name2" || u2.Id != 2 {
		t.Fatal("4. must not be from cache FAIL")
	}
}
