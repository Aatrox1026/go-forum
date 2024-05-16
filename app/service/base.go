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

const TIMEOUT time.Duration = 10 * time.Second

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
