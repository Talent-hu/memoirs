package cache

import (
	"context"
	"memoirs/global"
	"time"
)

type RedisCache struct {
}

func NewRedisCache() *RedisCache {
	return &RedisCache{}
}

func (r *RedisCache) Set(k string, v interface{}, expiration time.Duration) error {
	cmd := global.Redis.Set(context.Background(), k, v, expiration)
	return cmd.Err()
}

func (r *RedisCache) Get(key string) (interface{}, bool) {
	cmd := global.Redis.Get(context.Background(), key)
	if cmd.Err() != nil {
		return nil, false
	}
	return cmd.Val(), true
}

func (r *RedisCache) Del(k string) bool {
	cmd := global.Redis.Del(context.Background(), k)
	if cmd.Err() != nil {
		return false
	}
	return true
}
