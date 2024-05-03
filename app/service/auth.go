package service

import (
	"errors"
	"kevinku/go-forum/app/model"
	"net/http"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Register(registration *model.Registration) (resp *Response) {
	var result *gorm.DB
	// check if username exists
	var user model.User
	if result = db.Table("user").Where("name = ?", registration.Name).First(&user); result.Error == nil {
		// username already exists
		logger.Error("username already failed", zap.String("username", registration.Name))
		return &Response{StatusCode: http.StatusUnprocessableEntity, Error: errorf("username %s already exists", registration.Name)}
	} else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// other errors
		logger.Error("get user failed", zap.Any("error", result.Error))
		return &Response{StatusCode: http.StatusBadRequest, Error: errorf("get user failed: %v", result.Error)}
	}

	// check if password matches confirm
	if registration.Passwd != registration.Confirm {
		logger.Error("confirm doesn't matches password")
		return &Response{StatusCode: http.StatusBadRequest, Error: errorf("confirm doesn't matches password")}
	}

	return &Response{StatusCode: http.StatusOK}
}
