package service

import (
	"context"
	"fmt"
	"kevinku/go-forum/database"
	l "kevinku/go-forum/lib/logger"
	"time"
)

const (
	TIMEOUT time.Duration = 10 * time.Second
)

var f = fmt.Sprintf
var db = database.DB
var rdb = database.RDB
var logger = l.Logger

type Response struct {
	StatusCode int
	Data       any
	Error      error
}

func NewContextWithTimeout(timeout time.Duration) (ctx context.Context, cancel context.CancelFunc) {
	return context.WithTimeout(context.Background(), timeout)
}
