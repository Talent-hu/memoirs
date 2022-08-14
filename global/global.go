package global

import (
	"github.com/minio/minio-go/v7"
	"github.com/spf13/viper"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"memoirs/pkg/conf"
	"memoirs/pkg/database/redis"
	"memoirs/pkg/logger"
)

var (
	Viper              *viper.Viper
	AppConfig          *conf.AppConfig
	Log                *logger.Logger
	DB                 *gorm.DB
	Redis              *redis.Client
	Minio              *minio.Client
	Concurrent_Control = &singleflight.Group{}
)
