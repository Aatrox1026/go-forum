package controller

import (
	"github.com/gin-gonic/gin"
)

// CreateHost creates host
// @Summary create host
// @tags waf
// @Accept  json
// @Param request body model.User true "request data"
// @Produce json
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /v1/waf/host [post]
func SignUp(ctx *gin.Context) {
	// ctx.ShouldBindJSON()
}
