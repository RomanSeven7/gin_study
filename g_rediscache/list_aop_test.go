package g_rediscache

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
	"time"
)
type User struct {
	Id   int64
	Name string
}

func ExampleUseListAop() {
	RedisTestSetup()
	cacheKey := "testtest_list_" + strconv.Itoa(getRand())
	bizFunc := func() ([]interface{}, error) {
		var result []interface{}
		result = append(result, User{Id: 1}, User{Id: 2})
		return result, nil
	}
	_, fromCache, _ := UseListAop(cacheKey, reflect.TypeOf(User{})).WithStart(0).WithStop(-1).WithExpires(5*time.Second).Then(bizFunc)
	fmt.Println(fromCache)
	_, fromCache, _ = UseListAop(cacheKey, reflect.TypeOf(User{})).WithStart(0).WithStop(-1).WithExpires(5*time.Second).Then(bizFunc)
	fmt.Println(fromCache)
	// output:
	// false
	// true
}

func TestListAop(t *testing.T) {

	RedisTestSetup()
	cacheKey := "testtest_list_" + strconv.Itoa(getRand())

	GetRedisClient().Del(cacheKey)

	options := &ListOptions{}
	options.Key = cacheKey
	options.Rt = reflect.TypeOf(User{})
	options.Expires = 30 * time.Second
	options.Start = 0
	options.Stop = -1

	GetRedisClient().Del(cacheKey)

	// 第一次保证不从cache里面取值
	val, fromCache, err := ListAop(options, func() ([]interface{}, error) {
		var result []interface{}
		result = append(result, User{Id: 1}, User{Id: 2})
		return result, nil
	})

	if err != nil || fromCache || len(val) != 2 {
		t.Fatal("1. must not be from cache FAIL")
	}
	for index, vv := range val {
		u := vv.(User)
		if index == 0 && u.Id != 1 {
			t.Fatal("1. must not be from cache FAIL, get 0 val")
		}
		if index == 1 && u.Id != 2 {
			t.Fatal("1. must not be from cache FAIL, get 1 val")
		}
	}

	// 第二次保证从cache里面取值
	val, fromCache, err = ListAop(options, func() ([]interface{}, error) {
		var result []interface{}
		result = append(result, User{Id: 1}, User{Id: 2})
		return result, nil
	})

	if err != nil || !fromCache || len(val) != 2 {
		t.Fatal("2. must not be from cache FAIL")
	}
	for index, vv := range val {
		u := vv.(User)
		if index == 0 && u.Id != 1 {
			t.Fatal("2. must not be from cache FAIL, get 0 val")
		}
		if index == 1 && u.Id != 2 {
			t.Fatal("2. must not be from cache FAIL, get 1 val")
		}
	}

	// 第三次保证保证存进去EmptyFlag
	GetRedisClient().Del(cacheKey)
	options.EmptyExpires = 30 * time.Second
	val, fromCache, err = ListAop(options, func() ([]interface{}, error) {
		return nil, nil
	})

	if err != nil || fromCache || val != nil {
		t.Fatal("3. must be empty FAIL")
	}
	vals, _ := GetRedisClient().LRange(cacheKey, 0, 0).Result()
	if len(vals) != 1 && vals[0] != EmptyFlag {
		t.Fatal("3-1. must be empty FAIL")
	}

	// 第四保证取空值，并且从cache里取
	val, fromCache, err = ListAop(options, func() ([]interface{}, error) {
		return nil, nil
	})

	if err != nil || !fromCache || len(val) != 0 {
		t.Fatal("4. must be empty from cache FAIL")
	}

}
