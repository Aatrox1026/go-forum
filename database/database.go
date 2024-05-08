package database

import (
	"fmt"
	"kevinku/go-forum/app/model"
	"kevinku/go-forum/config"
	l "kevinku/go-forum/lib/logger"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var RDB *redis.Client

var cfg *config.Config
var logger *zap.Logger

func Init() {
	cfg = config.Cfg
	logger = l.Logger

	initMySQL()
	initRedis()
}

func initMySQL() {
	var err error
	var dsn string = fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		cfg.MySQL.User,
		cfg.MySQL.Password,
		cfg.MySQL.Host,
		cfg.MySQL.Port,
		cfg.MySQL.DB,
	)

	// connect to database
	if DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		logger.Panic(
			"connect to database failed",
			zap.Any("error", err))
	}
	logger.Info("Connected to MySQL")

	// migrate tables from models
	if err = DB.Table("user").AutoMigrate(&model.User{}); err != nil {
		logger.Panic("auto migrate table \"user\" failed", zap.Any("error", err))
	}
	logger.Info("MySQL migrate finished")
}

func initRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
	logger.Info("Connected to Redis")
}
