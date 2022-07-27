package config

import (
	"go.uber.org/zap"
	"memoirs/global"
)

func init() {
	// 加载配置文件
	InitLoadConfig()
	// 加载zap日志配置
	global.Log = InitZap()
	zap.ReplaceGlobals(global.Log)
	// 加载redis
	global.Redis = InitRedis()
	// 加载数据库连接
	global.DB = Gorm()
}
