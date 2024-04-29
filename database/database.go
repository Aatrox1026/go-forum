package database

import (
	"fmt"
	. "kevinku/go-forum/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var dsn string = fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		Cfg.MySQL.User,
		Cfg.MySQL.Password,
		Cfg.MySQL.Host,
		Cfg.MySQL.Port,
		Cfg.MySQL.DB,
	)

	var err error
	if DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		panic(fmt.Errorf("connect to database failed: %v", err))
	}
}
