package utils

import (
	"context"
	"gin_api_02/global"
	"time"

	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

func NewCaptchaRedisStore() *RedisStore {
	return &RedisStore{
		Expiration: time.Second * 180,
		PreKey:     "CAPTCHA_",
	}
}

type RedisStore struct {
	Expiration time.Duration
	PreKey     string
	Context    context.Context
}

func (rs *RedisStore) UseWithCtx(ctx context.Context) base64Captcha.Store {
	rs.Context = ctx
	return rs
}

func (rs *RedisStore) Set(id string, value string) error {
	err := global.SYS_REDIS.Set(rs.Context, rs.PreKey+id, value, rs.Expiration).Err()
	global.SYS_LOG.Info(id + "---" + value)
	if err != nil {
		global.SYS_LOG.Error("RedisStoreSetError!", zap.Error(err))
	}
	return err
}

func (rs *RedisStore) Get(key string, clear bool) string {
	val, err := global.SYS_REDIS.Get(rs.Context, key).Result()
	if err != nil {
		global.SYS_LOG.Error("RedisStoreGetError!", zap.Error(err))
		return ""
	}
	if clear {
		err := global.SYS_REDIS.Del(rs.Context, key).Err()
		if err != nil {
			global.SYS_LOG.Error("RedisStoreClearError!", zap.Error(err))
			return ""
		}
	}
	return val
}

func (rs *RedisStore) Verify(id, answer string, clear bool) bool {
	key := rs.PreKey + id
	v := rs.Get(key, clear)
	return v == answer
}
