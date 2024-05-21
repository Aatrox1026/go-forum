package controller

import (
	"kevinku/go-forum/middleware"

	"github.com/gin-gonic/gin"
)

// API Test
// @Summary api test
// @tags test
// @Security JWT
// @Produce json
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /api/v1/test [get]
func Test() gin.HandlerFunc {
	return middleware.PermissionCheck(1)
}
