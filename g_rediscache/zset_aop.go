package g_rediscache

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"mfx/gin_study/util"
	"reflect"
	"strconv"
	"time"
)

// ZsetAop 中，EmptyExpires 使用无效
type ZSetOptions struct {
	Options

	// 如果为true，表示返回map，
	// key是member，value是score
	// NOTICE!!! 并且返回结果的map的value一定是float64类型
	IsMap bool

	Desc    bool
	ByScore bool

	// 如果是struct类型，需要指定一下，使用哪个字段作为score，否则score默认是0
	ScoreField string

	// 如果是ByScore，必须需要指定Min和Max
	Min    int
	Max    int
	Offset int64
	Count  int64

	// 如果不是ByScore, 需要指定Start, Stop
	Start int64
	Stop  int64
}

// NOTICE!!! 如果fallback返回结果的map的value一定是float64类型
func ZSetAop(options *ZSetOptions, fallback func() (interface{}, error)) (interface{}, bool, error) {
	zrangeBy := redis.ZRangeBy{
		Min:    strconv.Itoa(options.Min),
		Max:    strconv.Itoa(options.Max),
		Offset: options.Offset,
		Count:  options.Count,
	}
	if options.Stop == 0 {
		options.Stop = -1
	}
	var cacheVs []redis.Z
	var err error = nil
	if options.Desc && options.ByScore {
		cacheVs, err = GetRedisClient().ZRevRangeByScoreWithScores(options.Key, zrangeBy).Result()
	} else if options.Desc {
		cacheVs, err = GetRedisClient().ZRevRangeWithScores(options.Key, options.Start, options.Stop).Result()
	} else if options.ByScore {
		cacheVs, err = GetRedisClient().ZRangeByScoreWithScores(options.Key, zrangeBy).Result()
	} else {
		cacheVs, err = GetRedisClient().ZRangeWithScores(options.Key, options.Start, options.Stop).Result()
	}
	if err != nil {
		return nil, false, err
	}
	var result []interface{}
	var mapResult = make(map[interface{}]float64)
	if len(cacheVs) > 0 {
		if len(cacheVs) == 1 && cacheVs[0].Member == EmptyFlag {
			if options.IsMap {
				return mapResult, true, nil
			} else {
				return result, true, nil
			}
		}
		for _, cacheV := range cacheVs {
			rtv := reflect.New(options.Rt)
			rv := rtv.Interface()
			err := json.Unmarshal([]byte(cacheV.Member.(string)), rv)
			if err != nil {
				return nil, false, err
			}
			vv := reflect.ValueOf(rv).Elem().Interface()

			if options.IsMap {
				mapResult[vv] = cacheV.Score
			} else {
				result = append(result, vv)
			}
		}
		if options.IsMap {
			return mapResult, true, nil
		} else {
			return result, true, nil
		}
	}

	logrus.Warn("[REDIS][ZSET] cant get value from redis cache, maybe load from db!")

	fResult, err := fallback()
	if err != nil {
		return nil, false, err
	}
	if fResult == nil {
		return nil, false, nil
	}
	rewriteCount := 0
	var members []redis.Z
	if options.IsMap {

		for k, score := range fResult.(map[interface{}]float64) {
			cacheV, isEmpty, err := GetCacheValueItem(k)
			if err != nil {
				logrus.Warn("[REDIS][ZSET] GetCacheValueItem error!", err)
				continue
			}
			if !isEmpty {
				members=append(members,redis.Z{Member: cacheV, Score: score} )
				rewriteCount++
			}
		}
	} else {
		for _, resultItem := range fResult.([]interface{}) {
			var score float64 = 0
			// 看一下struct里面的作为score的field是否有正确的值
			if options.ScoreField != "" {
				iv := reflect.ValueOf(&resultItem)
				if reflect.TypeOf(resultItem).Kind() == reflect.Struct {
					ivf := iv.Elem().Elem().FieldByName(options.ScoreField)
					if ivf.IsValid() {
						vv, success := util.Number2Float64(ivf.Interface(), ivf.Kind())
						if success {
							score = vv
						}
					}
				}
			}
			cacheV, isEmpty, err := GetCacheValueItem(resultItem)
			if err != nil {
				logrus.Warn("[REDIS][ZSET] GetCacheValueItem error!", err)
				continue
			}
			if !isEmpty {
				members=append(members,redis.Z{Member: cacheV, Score: score} )
				rewriteCount++
			}
		}

	}
	if len(members)>0 {
		GetRedisClient().ZAdd(options.Key,members... )
	}
	if rewriteCount > 0 {
		GetRedisClient().Expire(options.Key, options.Expires)
	} else {
		// 空值回填
		if rewriteCount == 0 && options.EmptyExpires > 0 {
			GetRedisClient().ZAdd(options.Key, redis.Z{Member: EmptyFlag})
			GetRedisClient().Expire(options.Key, options.EmptyExpires)
			logrus.Warn("[REDIS][ZSET] cache empty value, key:", options.Key)
		}
	}

	return fResult, false, nil
}

type ZSetAopProxy struct {
	options ZSetOptions
}

func (p *ZSetAopProxy) WithExpires(expires time.Duration) *ZSetAopProxy {
	p.options.Expires = expires
	return p
}

func (p *ZSetAopProxy) WithEmptyExpires(emptyExpires time.Duration) *ZSetAopProxy {
	p.options.EmptyExpires = emptyExpires
	return p
}

func (p *ZSetAopProxy) WithIsMap(isMap bool) *ZSetAopProxy {
	p.options.IsMap = isMap
	return p
}

func (p *ZSetAopProxy) WithDesc(desc bool) *ZSetAopProxy {
	p.options.Desc = desc
	return p
}

func (p *ZSetAopProxy) WithByScore(byScore bool) *ZSetAopProxy {
	p.options.ByScore = byScore
	return p
}

func (p *ZSetAopProxy) WithScoreField(scoreField string) *ZSetAopProxy {
	p.options.ScoreField = scoreField
	return p
}

func (p *ZSetAopProxy) WithMin(min int) *ZSetAopProxy {
	p.options.Min = min
	return p
}

func (p *ZSetAopProxy) WithMax(max int) *ZSetAopProxy {
	p.options.Max = max
	return p
}

func (p *ZSetAopProxy) WithOffset(offset int64) *ZSetAopProxy {
	p.options.Offset = offset
	return p
}

func (p *ZSetAopProxy) WithCount(count int64) *ZSetAopProxy {
	p.options.Count = count
	return p
}
func (p *ZSetAopProxy) WithStart(start int64) *ZSetAopProxy {
	p.options.Start = start
	return p
}

func (p *ZSetAopProxy) WithStop(stop int64) *ZSetAopProxy {
	p.options.Stop = stop
	return p
}

func (p *ZSetAopProxy) Then(f func() (interface{}, error)) (interface{}, bool, error) {
	return ZSetAop(&p.options, f)
}

func UseZSetAop(key string, rt reflect.Type) *ZSetAopProxy {
	return &ZSetAopProxy{ZSetOptions{Options: Options{Key: key, Rt: rt}}}
}
