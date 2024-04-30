package database

import (
	"fmt"
	"kevinku/go-forum/app/model"
	. "kevinku/go-forum/config"
	. "kevinku/go-forum/lib/logger"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	var dsn string = fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		Cfg.MySQL.User,
		Cfg.MySQL.Password,
		Cfg.MySQL.Host,
		Cfg.MySQL.Port,
		Cfg.MySQL.DB,
	)

	// connect to database
	if DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		Logger.Panic(
			"connect to database failed",
			zap.Any("error", err))
	}

	// migrate tables from models
	DB.Table("user").AutoMigrate(&model.User{})
}
