package service

import (
	"errors"
	"fmt"
	"kevinku/go-forum/app/model"
	"net/http"

	"gorm.io/gorm"
)

func Register(registration *model.Registration) (resp *Response) {
	var result *gorm.DB
	// check if username exists
	var user *model.User
	if result = db.Where("name = ?", registration.Name).First(user); result.Error == nil {
		// username already exists
		return &Response{}
	} else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {

	}

	// check if password matches confirm
	if registration.Passwd != registration.Confirm {
		logger.Error("confirm doesn't matches password")
		return &Response{StatusCode: http.StatusBadRequest, Error: fmt.Errorf("confirm doesn't matches password")}
	}

	return nil
}
