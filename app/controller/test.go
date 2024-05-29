package controller

import (
	"log"
	"time"

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
	return func(ctx *gin.Context) {
		go func() {
			time.Sleep(5 * time.Second)
			log.Println("second delete")
		}()
		log.Println("first delete")
		ctx.JSON(200, "success")
	}
}
