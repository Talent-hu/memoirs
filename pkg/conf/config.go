package conf

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"memoirs/pkg/database/mysql"
	"memoirs/pkg/database/redis"
	"memoirs/pkg/logger"
	"memoirs/pkg/mino"
)

type AppConfig struct {
	NetConf *NetConf `mapstructure:"netConf" yaml:"netConf"`
}

type NetConf struct {
	Name      string `mapstructure:"name" json:"name" yaml:"name"`
	Port      string `mapstructure:"port" json:"port" yaml:"port"`
	CacheMode string `mapstructure:"cache-mode" json:"cache-mode" yaml:"cache-mode"`
	CacheSize int    `mapstructure:"cache-size" json:"cache-size" yaml:"cache-size"`
}

func NewViper() (*viper.Viper, error) {
	var config string
	flag.StringVar(&config, "c", "", "please choose conf file")
	flag.Parse()
	if config == "" {
		config = "config.yml"
		fmt.Println("加载项目内部默认配置文件...")
	} else {
		fmt.Println("加载外部配置文件...")
	}
	v := viper.New()
	v.SetConfigFile(config)
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("%w unmarshal ReadInConfig error", err)
	}
	return v, nil
}

func NewAppConfig(v *viper.Viper) (conf *AppConfig, err error) {
	conf = new(AppConfig)
	if err = v.UnmarshalKey("netConf", &conf.NetConf); err != nil {
		err = fmt.Errorf("%w unmarshal log error", err)
	}
	return
}

func NewLoggerCfg(v *viper.Viper) (conf *logger.Options, err error) {
	conf = new(logger.Options)
	if err = v.UnmarshalKey("log", conf); err != nil {
		err = fmt.Errorf("%w unmarshal log error", err)
	}
	return
}

func NewMysqlCfg(v *viper.Viper) (conf *mysql.Config, err error) {
	conf = new(mysql.Config)
	if err = v.UnmarshalKey("mysql", conf); err != nil {
		err = fmt.Errorf("%w unmarshal mysql error", err)
	}
	return
}

func NewRedisCfg(v *viper.Viper) (conf *redis.Config, err error) {
	conf = new(redis.Config)
	if err = v.UnmarshalKey("redis", conf); err != nil {
		err = fmt.Errorf("%w unmarshal redis error", err)
	}
	return
}

func NewMinioCfg(v *viper.Viper) (conf *mino.MinioConfig, err error) {
	conf = new(mino.MinioConfig)
	if err = v.UnmarshalKey("minio", conf); err != nil {
		err = fmt.Errorf("%w unmarshal minio error", err)
	}
	return
}
