package g_rediscache

import (
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"mfx/gin_study/util"
	"reflect"
	"time"
)

type HashOptions struct {
	Options

	Fields []string

	// field attribute in model
	FieldAttr string
}

func (o HashOptions) validate() error {
	err := o.Options.validate()
	if err != nil {
		return err
	}

	if len(o.Fields) == 0 {
		return errors.New("Fields must not be empty!")
	}
	if o.FieldAttr == "" {
		return errors.New("FieldAttr must not be empty!")
	}

	return nil
}

func HashAop(options *HashOptions, fallback func() ([]interface{}, error)) ([]interface{}, bool, error) {
	err := options.validate()
	if err != nil {
		return nil, false, err
	}
	cacheVs, err := GetRedisClient().HMGet(options.Key, options.Fields...).Result()
	var result []interface{}
	shouldCallback := len(cacheVs) == 0
	if !shouldCallback {
		for _, cacheV := range cacheVs {
			if cacheV == nil {
				shouldCallback = true
				logrus.Warn("[REDIS][HASH] key ", options.Key, " has nil value, values", cacheVs)
				break
			}
			rtv := reflect.New(options.Rt)
			rv := rtv.Interface()
			err := json.Unmarshal([]byte(cacheV.(string)), rv)
			if err != nil {
				return nil, false, err
			}
			result = append(result, reflect.ValueOf(rv).Elem().Interface())
		}
		if !shouldCallback {
			return result, true, nil
		}
	}

	result, err = fallback()
	if err != nil {
		return nil, false, err
	}

	// 回填
	rewriteCount := 0
	if result != nil && len(result) > 0 {
		for _, item := range result {
			cacheV, isEmpty, err := GetCacheValueItem(item)
			if err != nil {
				logrus.Warn("[REDIS][HASH] GetCacheValueItem error!", err)
				continue
			}
			// 看一下struct里面的作为field的Field是否有正确的值
			fieldV := ""
			if options.FieldAttr != "" {
				iv := reflect.ValueOf(&item)
				if reflect.TypeOf(item).Kind() == reflect.Struct {
					ivf := iv.Elem().Elem().FieldByName(options.FieldAttr)
					if ivf.IsValid() {
						vv, success := util.Primary2String(ivf.Interface(), ivf.Kind())
						if success {
							fieldV = vv
						}
					}
				}
			}
			if fieldV == "" {
				logrus.Warn("[REDIS][HASH] key ", options.Key, " value ", item, " has not valid fieldValue!!!")
			}
			if !isEmpty && fieldV != "" {
				GetRedisClient().HSet(options.Key, fieldV, cacheV)
				rewriteCount++
			}
		}
		if rewriteCount > 0 {
			if options.Expires == 0 {
				options.Expires = defaultExpire
			}
			GetRedisClient().Expire(options.Key, options.Expires)
		}
	}

	return result, false, nil
}

type HashAopProxy struct {
	options HashOptions
}

func (p *HashAopProxy) WithExpires(expires time.Duration) *HashAopProxy {
	p.options.Expires = expires
	return p
}

func (p *HashAopProxy) WithEmptyExpires(emptyExpires time.Duration) *HashAopProxy {
	p.options.EmptyExpires = emptyExpires
	return p
}

func (p *HashAopProxy) Then(f func() ([]interface{}, error)) ([]interface{}, bool, error) {
	return HashAop(&p.options, f)
}

func UseHashAop(key string, rt reflect.Type, fields []string, fieldAttr string) *HashAopProxy {
	return &HashAopProxy{options: HashOptions{Options{Key: key, Rt: rt}, fields, fieldAttr}}
}
