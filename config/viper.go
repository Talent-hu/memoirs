package config

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"memoirs/global"
)

func InitLoadConfig() {
	var config string
	flag.StringVar(&config, "c", "", "please choose config file")
	flag.Parse()
	if config == "" {
		config = "config.yml"
		fmt.Println("加载项目内部默认配置文件...")
	} else {
		fmt.Println("加载外部配置文件...")
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("读取配置文件%失败！！", config))
	}
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		if err = v.Unmarshal(&global.Config); err != nil {
			fmt.Println(fmt.Errorf("修改配置文件异常：%s", err))
		}
	})
	if err = v.Unmarshal(&global.Config); err != nil {
		fmt.Println("配置文件序列化失败：", err)
	}
	global.Viper = v
}
