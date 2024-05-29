package service

import (
	"errors"
	"kevinku/go-forum/app/model"
	"kevinku/go-forum/lib/snowflake"
	"net/http"

	ginjwt "github.com/appleboy/gin-jwt/v2"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(registration *model.Registration) (code int, id int64, err error) {
	// check if username exists
	var user *model.User = new(model.User)
	if err = db.Where("name = ?", registration.Name).First(user).Error; err == nil {
		// username already exists
		logger.Info("username already exists", zap.String("username", registration.Name))
		return http.StatusUnprocessableEntity, -1, errorf("username %s already exists", registration.Name)
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// other errors
		logger.Error("get user failed", zap.Any("error", err))
		return http.StatusBadRequest, -1, errorf("get user failed: %v", err)
	}

	// check if password matches confirm
	if registration.Password != registration.Confirm {
		logger.Error("confirm doesn't matches password")
		return http.StatusBadRequest, -1, errorf("confirm doesn't matches password")
	}

	// insert data to db
	var hashed []byte
	if hashed, err = bcrypt.GenerateFromPassword([]byte(registration.Password), 10); err != nil {
		logger.Error("encrypt password failed", zap.Any("error", err))
		return http.StatusBadRequest, -1, errorf("encrypt password failed: %v", err)
	}

	user = &model.User{
		ID:       snowflake.NewID(),
		Name:     registration.Name,
		Email:    registration.Email,
		Password: string(hashed),
		Role:     model.ROLE_NORMAL,
	}

	if err = db.Create(user).Error; err != nil {
		logger.Error("create user failed", zap.Any("error", err))
		return http.StatusBadRequest, -1, errorf("create user failed: %v", err)
	}

	return http.StatusCreated, user.ID, nil
}

func Login(login *model.Login) (user *model.User, err error) {
	user = new(model.User)
	if err = db.Where("email = ?", login.Email).First(user).Error; err != nil {
		logger.Error("get user by name failed", zap.String("name", login.Email), zap.Any("error", err))
		return nil, ginjwt.ErrFailedAuthentication
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		logger.Error("compare password failed", zap.String("name", login.Email), zap.Any("error", err))
		return nil, ginjwt.ErrFailedAuthentication
	}

	return user, nil
}
