package service

import (
	"errors"
	"kevinku/go-forum/app/model"
	"net/http"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func GetUsers() (code int, users []*model.User, err error) {
	ctx, cancel := NewContextWithTimeout(TIMEOUT)
	defer cancel()

	var fields []string = []string{"id", "name", "email", "role", "created_at", "updated_at"}
	if err = db.WithContext(ctx).Select(fields).Find(&users).Error; err != nil {
		logger.Error("get users from mysql failed", zap.Any("error", err))
		return http.StatusNotFound, nil, err
	}
	return http.StatusOK, users, nil
}

func GetUserByID(id int64) (code int, user *model.User, err error) {
	ctx, cancel := NewContextWithTimeout(TIMEOUT)
	defer cancel()

	user = new(model.User)
	// try to get data from cache db
	if err = rdb.Get(ctx, f("user_%d", id), user); err == nil {
		return http.StatusOK, user.Safe(), nil
	}
	if !errors.Is(err, redis.Nil) {
		logger.Info("get user from redis failed", zap.Any("error", err))
	}

	// get data from main db
	ctx, cancel = NewContextWithTimeout(TIMEOUT)
	defer cancel()

	if err = db.WithContext(ctx).Where(&model.User{ID: id}).First(user).Error; err != nil {
		logger.Error("get user from mysql failed", zap.Any("error", err))
		return http.StatusNotFound, nil, errorf("get user failed: %v", err)
	}

	// set data to cache db
	ctx, cancel = NewContextWithTimeout(TIMEOUT)
	defer cancel()

	if err = rdb.Set(ctx, f("user_%d", id), user, REDIS_TTL); err != nil {
		logger.Error("set user to redis failed", zap.Any("error", err))
	}
	logger.Info("set user to redis", zap.String("key", f("user_%d", id)))

	return http.StatusOK, user.Safe(), nil
}
