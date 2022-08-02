package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"memoirs/global"
	"memoirs/model"
	"os"
	"time"
)

func Gorm() *gorm.DB {
	return gormMysql()
}

func gormMysql() *gorm.DB {
	m := global.Config.Mysql
	mysqlConfig := mysql.Config{
		DSN:                       m.Dns(),
		DefaultStringSize:         191,
		SkipInitializeWithVersion: false,
	}
	db, err := gorm.Open(mysql.New(mysqlConfig), gormLogConfig())
	if err != nil {
		return nil
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(m.MaxIdleConns)
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	//RegisterTable(db)
	return db
}

func RegisterTable(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.User{},
		&model.UserRole{},
		&model.Role{},
		&model.RoleMenu{},
		&model.Menu{},
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func gormLogConfig() *gorm.Config {
	config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	_default := logger.New(NewWriter(log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      true,
	})
	switch global.Config.LogMode {
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
