package service

import (
	"context"
	"fmt"
	"kevinku/go-forum/database"
	l "kevinku/go-forum/lib/logger"
	"kevinku/go-forum/lib/redis"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	DELAY     time.Duration = 2 * time.Second  // delay for double delete cache
	TIMEOUT   time.Duration = 10 * time.Second // timeout for database context
	REDIS_TTL time.Duration = 60 * time.Second
)

var f = fmt.Sprintf
var errorf = fmt.Errorf

var db *gorm.DB
var rdb *redis.Client
var logger *zap.Logger

type Json = map[string]any
type Response struct {
	Code  int
	Data  any
	Error error
}

func Init() {
	db = database.DB
	rdb = &redis.Client{
		Client: *database.RDB,
	}
	logger = l.Logger
}

func NewContextWithTimeout(timeout time.Duration) (ctx context.Context, cancel context.CancelFunc) {
	return context.WithTimeout(context.Background(), timeout)
}

func DoubleDelete(key string) {
	var err error
	time.Sleep(DELAY)

	ctx, cancel := NewContextWithTimeout(TIMEOUT)
	defer cancel()

	if err = rdb.Del(ctx, key).Err(); err != nil {
		logger.Error("double delete data failed", zap.Any("error", err))
		return
	}
	logger.Info("double delete data from redis", zap.String("key", key))
}
