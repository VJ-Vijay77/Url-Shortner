package controller

import (
	"urlShortner/model"
	"urlShortner/shared"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{
	Config *shared.Service
}
var Index = 1

func NewUserHandler(userService *shared.Service) *UserHandler{
	return &UserHandler{Config: userService}
}

func(s *UserHandler) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	  })
}

func(s *UserHandler) InputURL(c *gin.Context) {
type urlStruct struct {
	Url string `json:"url"`
}
var url urlStruct
if err := c.ShouldBindJSON(&url); err != nil {
	c.JSON(400, gin.H{
		"msg":"Unable to bing json",
		"error" : err,
	})
	return
}
var savetodb model.Url
savetodb.LongUrl = encodetobase64(url.Url)

c.JSON(200, url.Url)
}
func(s *UserHandler) Welcome(c *gin.Context) {
c.JSON(200, "Welcome Home")
}

