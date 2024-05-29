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
	var err error

	var code int
	var users []*model.User
	if code, users, err = service.GetUsers(); err != nil {
		HandleResponse(ctx, code, err)
		return
	}

	for i, user := range users {
		users[i] = user.Safe()
	}

	HandleResponse(ctx, code, Json{"users": users})
}

// Get User By ID
// @Summary get user by id
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

	var id int64
	if id, err = strconv.ParseInt(ctx.Param("id"), 10, 64); err != nil {
		HandleResponse(ctx, http.StatusBadRequest, fmt.Errorf("invalid user id: %v", id))
		return
	}

	var code int
	var user *model.User
	if code, user, err = service.GetUserByID(id); err != nil {
		HandleResponse(ctx, code, err)
		return
	}
	HandleResponse(ctx, http.StatusOK, Json{"user": user.Safe()})
}

// Disable User Account
// @Summary disable user account
// @tags user
// @Security JWT
// @Param id path int64 true "user id"
// @Produce json
// @Success 204
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /api/v1/users/ban/{id} [patch]
func BanUser(ctx *gin.Context) {
	var err error

	var id int64
	if id, err = strconv.ParseInt(ctx.Param("id"), 10, 64); err != nil {
		HandleResponse(ctx, http.StatusBadRequest, fmt.Errorf("invalid user id: %v", id))
		return
	}

	var code int
	if code, err = service.BanUser(id); err != nil {
		HandleResponse(ctx, code, err)
	}

	HandleResponse(ctx, code, nil)
}
