package g_rediscache

import (
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"sync"
	"testing"
	"time"
)

func ExampleUseGLock() {
	RedisTestSetup()
	cacheKey := "testtest_global_lock_" + strconv.Itoa(getRand())
	res, err := UseGLock(cacheKey).WithTimeout(-1).WithExpire(10 * time.Second).Then(func() (i interface{}, e error) {
		return 1, e
	})

	if err != nil {
		panic(err )
	}

	fmt.Println(res)
	// output:
	// 1

}

func TestUseGLock_BizResult(t *testing.T) {
	RedisTestSetup()

	cacheKey := "testtest_global_lock_" + strconv.Itoa(getRand())
	t.Log("cacheKey:", cacheKey)

	expected := "ha"
	res, err := UseGLock(cacheKey).WithTimeout(0).WithExpire(10 * time.Second).Then(func() (i interface{}, e error) {
		return expected, nil
	})
	if err != nil {
		t.Fatal("acquire global lock fail", err)
	}

	t.Log("res:", res)
	if res != expected {
		t.Error("res is not expected", res)
	}

}

func TestUseGLock_Release(t *testing.T) {
	RedisTestSetup()

	cacheKey := "testtest_global_lock_" + strconv.Itoa(getRand())
	t.Log("cacheKey:", cacheKey)
	go func() {
		_, err := UseGLock(cacheKey).WithTimeout(-1).WithExpire(10 * time.Second).Then(func() (i interface{}, e error) {
			time.Sleep(5 * time.Second)
			return nil, nil
		})
		if err != nil {
			t.Fatal("acquire global lock fail", err)
		}
	}()
	time.Sleep(1 * time.Second)

	res, err := GetRedisClient().Get(cacheKey).Result()
	if err == redis.Nil {
		t.Error("acquire global lock fail", res)
		t.Fail()
	} else if err != nil {
		t.Fatal(err)
	}

	time.Sleep(5 * time.Second)

	res, err = GetRedisClient().Get(cacheKey).Result()
	if err == redis.Nil {
	} else if err != nil {
		t.Fatal(err)
	}
}

func TestGlobalLock(t *testing.T) {
	RedisTestSetup()

	cacheKey := "testtest_global_lock_" + strconv.Itoa(getRand())

	options := &GlobalLockOptions{
		Key:     cacheKey,
		Expire:  5 * time.Second,
		Timeout: 10 * time.Second,
	}

	GetRedisClient().Del(cacheKey)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		_, err := GlobalLock(options, func() (interface{}, error) {
			t.Log("hahaha1")
			// 休眠3秒
			time.Sleep(3 * time.Second)
			return nil, nil
		})

		if err != nil {
			t.Fatal("1. acquire global lock FAIL")
		}
	}()

	// 确保上面的go先执行
	time.Sleep(300 * time.Millisecond)
	go func() {
		defer wg.Done()
		startTime := time.Now()
		_, err := GlobalLock(options, func() (interface{}, error) {
			t.Log("hahaha2")
			return nil, nil
		})
		if err != nil {
			t.Fatal("2. acquire global lock FAIL")
		}
		if time.Now().Sub(startTime) < 2*time.Second {
			t.Fatal("2.1. should not be happened, last operation hold lock 3s!")
		}
	}()

	wg.Wait()
}
