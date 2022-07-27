package config

import "fmt"

type Server struct {
	System `mapstructure:"system" json:"system" yaml:"system"`
	Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}

type System struct {
	Name      string `mapstructure:"name" json:"name" yaml:"name"`
	Port      string `mapstructure:"port" json:"port" yaml:"port"`
	CacheMode string `mapstructure:"cache-mode" json:"cache-mode" yaml:"cache-mode"`
}

type Mysql struct {
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	Path         string `mapstructure:"path" json:"path" yaml:"path"`
	Port         int    `mapstructure:"port" json:"port" yaml:"port"`
	DBName       string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"`
	LogMode      string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"`
	ZapLog       bool   `mapstructure:"zap-log" json:"zap-log" yaml:"zap-log"`
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
	LogInConsole  bool   `mapstructure:"log-in-console" json:"log-in-console" yaml:"log-in-console"`
}


type JWT struct {
	SignKey    string `mapstructure:"sign-key" json:"sign-key" yaml:"sign-key"`
	AppSecret  string `mapstructure:"app-secret" json:"app-secret" yaml:"app-secret"`
	AppIss     string `mapstructure:"app-iss" json:"app-iss" yaml:"app-iss"`
	ExpireTime int64  `mapstructure:"expire-time" json:"expire-time" yaml:"expire-time"`
	BufferTime int64  `mapstructure:"buffer-time" json:"buffer-time" yaml:"buffer-time"`
}
