package router

import (
	"urlShortner/controller"

	"github.com/gin-gonic/gin"
)


func Router(userHandler *controller.UserHandler) *gin.Engine {
	r := gin.Default()
	r.GET("/ping",userHandler.Ping)
	r.POST("/url",userHandler.InputURL)
	r.GET("/welcome/12345",userHandler.Welcome)
	return r
}