package controller

import (
	"kevinku/go-forum/app/model"
	"kevinku/go-forum/app/service"
	"kevinku/go-forum/database"

	"github.com/gin-gonic/gin"
)

// User Sign Up
// @Summary user sign up
// @tags auth
// @Accept  json
// @Param request body model.Registration true "registration data"
// @Produce json
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /v1/waf/host [post]
func Register(ctx *gin.Context) {
	_ = database.DB

	var registration = new(model.Registration)
	ctx.ShouldBindJSON(registration)

	service.Register()

}
