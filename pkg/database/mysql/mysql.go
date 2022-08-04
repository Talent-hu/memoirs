package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Config struct {
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	Path         string `mapstructure:"path" json:"path" yaml:"path"`
	Port         int    `mapstructure:"port" json:"port" yaml:"port"`
	DBName       string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"`
	LogMode      string `mapstructure:"logger-mode" json:"logger-mode" yaml:"logger-mode"`
	ZapLog       bool   `mapstructure:"zap-logger" json:"zap-logger" yaml:"zap-logger"`
}

func (this *Config) Dns() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		this.Username, this.Password, this.Path, this.Port, this.DBName, this.Config)
}

func (this *Config) GetLogMode() string {
	return this.LogMode
}

func NewMysql(config *Config) *gorm.DB {
	mysqlConfig := mysql.Config{
		DSN:                       config.Dns(),
		DefaultStringSize:         191,
		SkipInitializeWithVersion: false,
	}
	db, err := gorm.Open(mysql.New(mysqlConfig), gormLogConfig(config.LogMode))
	if err != nil {
		fmt.Println("连接mysql数据库失败，err:", err)
		panic(err)
	}
	db.DisableAutomaticPing = true
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	return db
}

func gormLogConfig(logMode string) *gorm.Config {
	config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	_default := logger.New(NewWriter(log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      true,
	})
	switch logMode {
	case "silent", "Silent":
		config.Logger = _default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = _default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = _default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = _default.LogMode(logger.Info)
	default:
		config.Logger = _default.LogMode(logger.Info)
	}
	return config
}

type writer struct {
	logger.Writer
}

func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}
