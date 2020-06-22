package g_rediscache

import (
	"errors"
	"github.com/sirupsen/logrus"
	"math/rand"
	"mfx/gin_study/util"
	"time"
)

type GlobalLockOptions struct {
	Key     string
	Timeout time.Duration
	Expire  time.Duration
}

func GlobalLock(options *GlobalLockOptions, fallback func() (interface{}, error)) (interface{}, error) {
	if options.Key == "" {
		return nil, errors.New("key must not be empty")
	}

	if options.Timeout == 0 {
		options.Timeout = defaultTimeout
	}
	if options.Expire == 0 {
		options.Expire = defaultExpire
	}
	startTime := time.Now()
	for {
		unique := util.NewObjectID().Hex()
		success, err := GetRedisClient().SetNX(options.Key, unique, options.Expire).Result()
		if err == nil && success {
			defer GetRedisClient().Eval("if redis.call('get', KEYS[1]) == ARGV[1] then return redis.call('del', KEYS[1]) else return 0 end", []string{options.Key}, unique)
			logrus.Debug("add lock success: ", options.Key)
			return fallback()
		}
		logrus.Debug("waiting lock: ", options.Key)
		// 短暂休眠，避免可能的活锁
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)

		spendTime := time.Now().Sub(startTime)

		if spendTime > options.Timeout {
			logrus.Warn("acquireLock timeout: ", options.Key)
			return nil, errors.New("acquireLock timeout " + options.Key)
		}
	}
}

type GLockProxy struct {
	options GlobalLockOptions
}

func (p *GLockProxy) WithTimeout(timeout time.Duration) *GLockProxy {
	p.options.Timeout = timeout
	return p
}
func (p *GLockProxy) WithExpire(expire time.Duration) *GLockProxy {
	p.options.Expire = expire
	return p
}

func (p *GLockProxy) Then(f func() (interface{}, error)) (interface{}, error) {
	return GlobalLock(&p.options, f)
}

func UseGLock(key string) *GLockProxy {
	return &GLockProxy{options: GlobalLockOptions{Key: key}}
}