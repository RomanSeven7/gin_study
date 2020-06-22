package g_rediscache

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"reflect"
	"time"
)

type ListOptions struct {
	Options
	// 此处一定要注意，除非要取第一条数据，否则一定要设置Start和Stop
	Start int64
	Stop  int64
}

func ListAop(options *ListOptions, fallback func() ([]interface{}, error)) ([]interface{}, bool, error) {
	err := options.validate()
	if err != nil {
		return nil, false, err
	}

	cacheVs, err := GetRedisClient().LRange(options.Key, options.Start, options.Stop).Result()
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
			logrus.Info("[Reflect] Value.", rv)
			result = append(result, reflect.ValueOf(rv).Elem().Interface())
			logrus.Info("[Reflect] Result.", result)
		}
		return result, true, err
	}
	logrus.Warn("[REDIS][LIST] cant get value from redis cache, maybe load from db!")
	result, err = fallback()
	if err != nil {
		return nil, false, err
	}
	// 回填
	rewriteCount := 0
	if result != nil && len(result) > 0 {
		var cacheVList []string
		for _, item := range result {
			cacheV, isEmpty, err := GetCacheValueItem(item)
			if err != nil {
				logrus.Warn("[REDIS][LIST] GetCacheValueItem error!", err)
				continue
			}
			if !isEmpty {
				rewriteCount++
				cacheVList = append(cacheVList, cacheV)
			}
		}

		GetRedisClient().RPush(options.Key, cacheVList)
		if rewriteCount > 0 {
			if options.Expires == 0 {
				options.Expires = defaultExpire
			}
			GetRedisClient().Expire(options.Key, options.Expires)
		}
	}

	// 空值回填
	if rewriteCount == 0 && options.EmptyExpires > 0 {
		GetRedisClient().RPush(options.Key, EmptyFlag)
		GetRedisClient().Expire(options.Key, options.EmptyExpires)
		logrus.Warn("[REDIS][LIST] cache empty value, key:", options.Key)
	}

	return result, false, nil

}

type ListAopProxy struct {
	options ListOptions
}

func (p *ListAopProxy) WithExpires(expires time.Duration) *ListAopProxy {
	p.options.Expires = expires
	return p
}

func (p *ListAopProxy) WithEmptyExpires(emptyExpires time.Duration) *ListAopProxy {
	p.options.EmptyExpires = emptyExpires
	return p
}

func (p *ListAopProxy) WithStart(start int64) *ListAopProxy {
	p.options.Start = start
	return p
}

func (p *ListAopProxy) WithStop(stop int64) *ListAopProxy {
	p.options.Stop = stop
	return p
}

func (p *ListAopProxy) Then(f func() ([]interface{}, error)) ([]interface{}, bool, error) {
	return ListAop(&p.options, f)
}

func UseListAop(key string, rt reflect.Type) *ListAopProxy {
	return &ListAopProxy{ListOptions{Options: Options{Key: key, Rt: rt}}}
}
