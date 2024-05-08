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
// @Router /api/v1/auth/sign-up [post]
func Register(ctx *gin.Context) {
	var err error
	var resp *service.Response

	var registration = new(model.Registration)
	if err = ctx.ShouldBindJSON(registration); err != nil {
		logger.Error(
			"invalid request body",
			zap.Any("error", err))
		HandleResponse(ctx, http.StatusUnprocessableEntity, fmt.Errorf("invalid request body: %v", err))
		return
	}

	if resp = service.Register(registration); resp.Error != nil {
		HandleResponse(ctx, resp.StatusCode, resp.Error)
		return
	}
	HandleResponse(ctx, resp.StatusCode, resp.Data)
}
