package g_rediscache

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"reflect"
	"strings"
	"sync"
)

func Init() {
	BuildRedisClient()
}

var (
	mu   sync.Mutex
	mode string
	// NOTICE!!!
	RedisClient *redis.ClusterClient
	single      *redis.Client
	sentinel    *redis.Client
)

func BuildRedisClient() {
	mu.Lock()
	defer mu.Unlock()
	mode = viper.GetString("redis.mode")

	c := GetRedisClient()
	vc := reflect.ValueOf(c)
	if !vc.IsNil() {
		return
	}

	addrs := viper.GetString("redis.conn")
	if addrs == "" {
		logrus.Info("there is not redis.conn config, skip build redis client")
		return
	}
	password := viper.GetString("redis.password")
	diaTimeout := viper.GetDuration("redis.conn_timeout") * 1e6
	readTimeout := viper.GetDuration("redis.so_timeout") * 1e6
	minIdleConns := viper.GetInt("redis.min_idle_conns")
	maxRetries := viper.GetInt("redis.max_retries")

	switch mode {
	case "single":
		options := &redis.Options{
			Addr:         addrs,
			Password:     password,
			DialTimeout:  diaTimeout,
			ReadTimeout:  readTimeout,
			WriteTimeout: readTimeout,
			MinIdleConns: minIdleConns,
			MaxRetries:   maxRetries,
		}
		single = redis.NewClient(options)

	case "sentinel":
		masterName := viper.GetString("redis.sentinel.master")
		if len(masterName) == 0 {
			panic("redis.sentinel.master must not empty when redis.mode=sentinel")
		}
		option := &redis.FailoverOptions{
			MasterName:    masterName,
			SentinelAddrs: strings.Split(addrs, ","),
			Password:      password,
			DialTimeout:   diaTimeout,
			ReadTimeout:   readTimeout,
			WriteTimeout:  readTimeout,
			MinIdleConns:  minIdleConns,
			MaxRetries:    maxRetries,
		}
		sentinel = redis.NewFailoverClient(option)
	default:
		RedisClient = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:        strings.Split(addrs, ","),
			Password:     password,
			DialTimeout:  diaTimeout,
			ReadTimeout:  readTimeout,
			WriteTimeout: readTimeout,
			MinIdleConns: minIdleConns,
			MaxRetries:   maxRetries,
		})
	}

	_, err := GetRedisClient().Ping().Result()

	if err != nil {
		logrus.Error("Redis Connection error. :", err)
		RedisClient = nil
	}
}

func GetRedisClient() redis.Cmdable {
	switch mode {
	case "single":
		return single
	case "sentinel":
		return sentinel
	default:
		return RedisClient
	}
}
