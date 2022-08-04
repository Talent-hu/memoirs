package config

import (
	"memoirs/global"
	"memoirs/pkg/conf"
	"memoirs/pkg/database/mysql"
	"memoirs/pkg/database/redis"
	"memoirs/pkg/logger"
)

func NewApp() error {
	viper, err := conf.NewViper()
	if err != nil {
		return err
	}
	appConfig, err := conf.NewAppConfig(viper)
	if err != nil {
		return err
	}
	global.AppConfig = appConfig
	options, err := conf.NewLoggerCfg(viper)
	if err != nil {
		return err
	}
	log, err := logger.NewLogger(options)
	if err != nil {
		return err
	}
	global.Log = log
	mysqlConfig, err := conf.NewMysqlCfg(viper)
	if err != nil {
		return err
	}
	mysqlClient := mysql.NewMysql(mysqlConfig)
	redisConfig, err := conf.NewRedisCfg(viper)
	if err != nil {
		return err
	}
	global.DB = mysqlClient
	redisClient, err := redis.NewRedis(redisConfig)
	if err != nil {
		return err
	}
	global.Redis = redisClient
	return nil
}
