package database

import (
	"fmt"
	"kevinku/go-forum/app/model"
	. "kevinku/go-forum/config"
	. "kevinku/go-forum/lib/logger"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var RDB *redis.Client

func init() {
	initMySQL()
	initRedis()
}

func initMySQL() {
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
	Logger.Info("Connected to MySQL")

	// migrate tables from models
	if err = DB.Table("user").AutoMigrate(&model.User{}); err != nil {
		Logger.Panic(
			"database auto migrate failed",
			zap.Any("error", err))
	}
	Logger.Info("MySQL migrate finished")
}

func initRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", Cfg.Redis.Host, Cfg.Redis.Port),
		Password: Cfg.Redis.Password,
		DB:       Cfg.Redis.DB,
	})
	Logger.Info("Connected to Redis")
}
