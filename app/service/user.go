package service

import (
	"errors"
	"kevinku/go-forum/app/model"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func GetUserByID(id int64) (user *model.User, err error) {
	ctx, cancel := NewContextWithTimeout(TIMEOUT)
	defer cancel()

	user = new(model.User)
	// try to get data from cache db
	if err = rdb.Get(ctx, f("user_%d", id), user); err == nil {
		user.Password = ""
	}
	if !errors.Is(err, redis.Nil) {
		logger.Info("get user from redis failed", zap.Any("error", err))
	}

	// get data from main db
	ctx, cancel = NewContextWithTimeout(TIMEOUT)
	defer cancel()

	if err = db.WithContext(ctx).Where(&model.User{ID: id}).First(user).Error; err != nil {
		logger.Error("get user from mysql failed", zap.Any("error", err))
		return nil, errorf("get user failed: %v", err)
	}

	// set data to cache db
	ctx, cancel = NewContextWithTimeout(TIMEOUT)
	defer cancel()

	if err = rdb.Set(ctx, f("user_%d", id), user, REDIS_TTL); err != nil {
		logger.Error("set user to redis failed", zap.Any("error", err))
	}

	return user, nil
}
