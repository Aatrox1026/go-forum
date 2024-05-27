package controller

import (
	"fmt"
	"kevinku/go-forum/app/model"
	"kevinku/go-forum/app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Get All Users
// @Summary get all users
// @tags user
// @Security JWT
// @Produce json
// @Success 200 {object} []model.User
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /api/v1/users [get]
func GetUsers(ctx *gin.Context) {
	var code int
	var users []*model.User
	var err error

	if code, users, err = service.GetUsers(); err != nil {
		HandleResponse(ctx, code, err)
		return
	}
	HandleResponse(ctx, code, users)
}

// User Sign Up
// @Summary user sign up
// @tags user
// @Security JWT
// @Param id path int64 true "user id"
// @Produce json
// @Success 200 {object} model.User
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /api/v1/users/{id} [get]
func GetUserByID(ctx *gin.Context) {
	var err error
	var code int
	var user *model.User

	var id int64
	if id, err = strconv.ParseInt(ctx.Param("id"), 10, 64); err != nil {
		HandleResponse(ctx, http.StatusBadRequest, fmt.Errorf("invalid user id: %v", id))
		return
	}

	if code, user, err = service.GetUserByID(id); err != nil {
		HandleResponse(ctx, code, err)
		return
	}
	HandleResponse(ctx, http.StatusOK, user)
}
