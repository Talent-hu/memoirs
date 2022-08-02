package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"memoirs/pkg/config"
)

var (
	Viper              *viper.Viper
	Config             *config.Server
	Log                *zap.Logger
	DB                 *gorm.DB
	Redis              *redis.Client
	Concurrent_Control = &singleflight.Group{}
	//Caches *cache.CacheManage
)
