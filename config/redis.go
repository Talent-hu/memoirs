package config

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"memoirs/global"
)

func InitRedis() *redis.Client {
	redisCfg := global.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.Log.Error("redis connect ping failed:", zap.Error(err))
	} else {
		global.Log.Info("redis connect ping success:", zap.String("pong", pong))
	}
	return client
}
