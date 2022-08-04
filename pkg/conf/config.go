package conf

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"memoirs/pkg/database/mysql"
	"memoirs/pkg/database/redis"
	"memoirs/pkg/logger"
)

type AppConfig struct {
	NetConf *NetConf `mapstructure:"netConf" yaml:"netConf"`
	/*Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`*/
}

type NetConf struct {
	Name      string `mapstructure:"name" json:"name" yaml:"name"`
	Port      string `mapstructure:"port" json:"port" yaml:"port"`
	CacheMode string `mapstructure:"cache-mode" json:"cache-mode" yaml:"cache-mode"`
	CacheSize int    `mapstructure:"cache-size" json:"cache-size" yaml:"cache-size"`
}

/*
type Mysql struct {
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	Path         string `mapstructure:"path" json:"path" yaml:"path"`
	Port         int    `mapstructure:"port" json:"port" yaml:"port"`
	DBName       string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`
	Config       string `mapstructure:"conf" json:"conf" yaml:"conf"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"`
	LogMode      string `mapstructure:"logger-mode" json:"logger-mode" yaml:"logger-mode"`
	ZapLog       bool   `mapstructure:"zap-logger" json:"zap-logger" yaml:"zap-logger"`
}

func (this *Mysql) Dns() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		this.Username, this.Password, this.Path, this.Port, this.DBName, this.Config)
}

func (this *Mysql) GetLogMode() string {
	return this.LogMode
}

type Redis struct {
	Addr       string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Password   string `mapstructure:"password" json:"password" yaml:"password"`
	DB         int    `mapstructure:"db" json:"db" yaml:"db"`
	MaxIdle    int    `mapstructure:"max-idle" json:"max-idle" yaml:"max-idle"`
	MaxActive  int    `mapstructure:"max-active" json:"max-active" yaml:"max-active"`
	MaxTimeout int    `mapstructure:"max-timeout" json:"max-timeout" yaml:"max-timeout"`
}

type Zap struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`
	Format        string `mapstructure:"format" json:"format" yaml:"format"`
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	Director      string `mapstructure:"director" json:"director" yaml:"director"`
	ShowLine      bool   `mapstructure:"show-line" json:"show-line" yaml:"show-line"`
	EncodeLevel   string `mapstructure:"encode-level" json:"encode-level" yaml:"encode-level"`
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktrace-key" yaml:"stacktrace-key"`
	LogInConsole  bool   `mapstructure:"logger-in-console" json:"logger-in-console" yaml:"logger-in-console"`
}

type JWT struct {
	SignKey    string `mapstructure:"sign-key" json:"sign-key" yaml:"sign-key"`
	AppSecret  string `mapstructure:"app-secret" json:"app-secret" yaml:"app-secret"`
	AppIss     string `mapstructure:"app-iss" json:"app-iss" yaml:"app-iss"`
	ExpireTime int64  `mapstructure:"expire-time" json:"expire-time" yaml:"expire-time"`
	BufferTime int64  `mapstructure:"buffer-time" json:"buffer-time" yaml:"buffer-time"`
}*/

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
