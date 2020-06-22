package g_rediscache

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"reflect"
	"time"
)

type SetOptions struct {
	Options
}

func SetAop(options *SetOptions, fallback func() ([]interface{}, error)) ([]interface{}, bool, error) {
	err := options.validate()
	if err != nil {
		return nil, false, err
	}
	cacheVs, err := GetRedisClient().SMembers(options.Key).Result()

	var result []interface{}
	// 从cache里取到值
	if len(cacheVs) > 0 {
		if len(cacheVs) == 1 && cacheVs[0] == EmptyFlag {
			return result, true, nil
		}
		for _, cacheV := range cacheVs {
			if cacheV == EmptyFlag {
				continue
			}
			rtv := reflect.New(options.Rt)
			rv := rtv.Interface()
			err := json.Unmarshal([]byte(cacheV), rv)
			if err != nil {
				return nil, false, err
			}
			result = append(result, reflect.ValueOf(rv).Elem().Interface())
		}
		return result, true, err
	}
	logrus.Warn("[REDIS][SET] cant get value from redis cache, maybe load from db!")
	result, err = fallback()
	if err != nil {
		return nil, false, err
	}
	// 回填
	rewriteCount := 0
	if result != nil && len(result) > 0 {
	 var	val []string
		for _, item := range result {
			cacheV, isEmpty, err := GetCacheValueItem(item)
			if err != nil {
				logrus.Warn("[REDIS][SET] GetCacheValueItem error!", err)
				continue
			}
			if !isEmpty {
				val = append(val, cacheV)
				rewriteCount++
			}
		}

		if rewriteCount > 0 {
			GetRedisClient().SAdd(options.Key, val)
			if options.Expires == 0 {
				options.Expires = defaultExpire
			}
			GetRedisClient().Expire(options.Key, options.Expires)
		}
	}

	// 空值回填
	if rewriteCount == 0 && options.EmptyExpires > 0 {
		GetRedisClient().SAdd(options.Key, EmptyFlag)
		GetRedisClient().Expire(options.Key, options.EmptyExpires)
		logrus.Warn("[REDIS][SET] cache empty value, key:", options.Key)
	}

	return result, false, nil

}

type SetAopProxy struct {
	options SetOptions
}

func (p *SetAopProxy) WithExpires(expires time.Duration) *SetAopProxy {
	p.options.Expires = expires
	return p
}

func (p *SetAopProxy) WithEmptyExpires(emptyExpires time.Duration) *SetAopProxy {
	p.options.EmptyExpires = emptyExpires
	return p
}

func (p *SetAopProxy) Then(f func() ([]interface{}, error)) ([]interface{}, bool, error) {
	return SetAop(&p.options, f)
}

func UseSetAop(key string, rt reflect.Type) *SetAopProxy {
	return &SetAopProxy{SetOptions{Options{Key: key, Rt: rt}}}
}
