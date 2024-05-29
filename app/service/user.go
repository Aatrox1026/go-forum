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

	if err = db.WithContext(ctx).Find(&users).Error; err != nil {
		logger.Error("get users from mysql failed", zap.Any("error", err))
		return http.StatusNotFound, nil, err
	}
	return http.StatusOK, users, nil
}

func GetUserByID(id int64) (code int, user *model.User, err error) {
	ctx, cancel := NewContextWithTimeout(TIMEOUT)
	defer cancel()

	var key string = f("user:%d", id)
	user = new(model.User)
	// try to get data from cache db
	if err = rdb.Get(ctx, key, user); err == nil {
		return http.StatusOK, user, nil
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

	if err = rdb.Set(ctx, key, user, REDIS_TTL); err != nil {
		logger.Error("set user to redis failed", zap.Any("error", err))
	}
	logger.Info("set user to redis", zap.String("key", key))

	return http.StatusOK, user, nil
}

func BanUser(id int64) (code int, err error) {
	ctx, cancel := NewContextWithTimeout(TIMEOUT)
	defer cancel()

	db.WithContext(ctx).Model(&model.User{ID: id}).Update("enabled", "false")

	return http.StatusNoContent, nil
}
