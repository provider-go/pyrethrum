package models

import (
	"github.com/provider-go/pkg/go-logger"
	"github.com/provider-go/pyrethrum/core/global"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysql() {
	var err error
	dsn := viper.GetString("mysql.dsn")
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	sqlDB, _ := global.DB.DB()
	if viper.GetInt("mysql.maxIdleConns") != 0 {
		sqlDB.SetMaxIdleConns(viper.GetInt("mysql.maxIdleConns")) // 设置最大空闲数
	}
	if viper.GetInt("mysql.maxOpenConns") != 0 {
		sqlDB.SetMaxOpenConns(viper.GetInt("mysql.maxOpenConns")) // 设置最大连接数
	}

	if err != nil {
		logger.Error("MySQL connection failed.", "err", err.Error())
	} else {
		logger.Info("MySQL connection successful!")
	}
}
