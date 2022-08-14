package config

import (
	"fmt"
	"memoirs/global"
	"memoirs/pkg/conf"
	"memoirs/pkg/database/mysql"
	"memoirs/pkg/database/redis"
	"memoirs/pkg/logger"
	"memoirs/pkg/mino"
)

func NewApp() error {
	viper, err := conf.NewViper()
	if err != nil {
		return err
	}
	global.Viper = viper
	fmt.Println("加载viper配置文件成功")
	global.AppConfig, err = conf.NewAppConfig(viper)
	if err != nil {
		return err
	}
	fmt.Println("加载app系统配置成功")
	options, err := conf.NewLoggerCfg(viper)
	if err != nil {
		return err
	}
	global.Log, err = logger.NewLogger(options)
	if err != nil {
		return err
	}
	fmt.Println("系统日志配置成功")
	mysqlConfig, err := conf.NewMysqlCfg(viper)
	if err != nil {
		return err
	}
	global.DB = mysql.NewMysql(mysqlConfig)
	fmt.Println("连接mysql成功！")
	redisConfig, err := conf.NewRedisCfg(viper)
	if err != nil {
		return err
	}
	global.Redis, err = redis.NewRedis(redisConfig)
	if err != nil {
		return err
	}
	fmt.Println("连接redis成功")
	minioConfig, err := conf.NewMinioCfg(viper)
	if err != nil {
		return err
	}
	global.Minio, err = mino.NewMinio(minioConfig)
	if err != nil {
		return err
	}
	fmt.Println("连接minio对象存储成功")
	return nil
}
