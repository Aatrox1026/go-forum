package controller

import (
	"fmt"
	"kevinku/go-forum/app/model"
	"kevinku/go-forum/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// User Sign Up
// @Summary user sign up
// @tags auth
// @Accept  json
// @Param request body model.Registration true "registration data"
// @Produce json
// @Success 201 {object} string
// @Failure 400 {object} string
// @Failure 422 {object} string
// @Router /auth/sign-up [post]
func Register(ctx *gin.Context) {
	var err error
	var code int
	var data Json

	var registration = new(model.Registration)
	if err = ctx.ShouldBindJSON(registration); err != nil {
		logger.Error(
			"invalid request body",
			zap.Any("error", err))
		HandleResponse(ctx, http.StatusUnprocessableEntity, fmt.Errorf("invalid request body: %v", err))
		return
	}

	if code, data, err = service.Register(registration); err != nil {
		HandleResponse(ctx, code, err)
		return
	}
	HandleResponse(ctx, code, data)
}
