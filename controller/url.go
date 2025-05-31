package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"urlShortner/model"
	"urlShortner/shared"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Config *shared.Service
}

var Index = 0

func NewUserHandler(userService *shared.Service) *UserHandler {
	return &UserHandler{Config: userService}
}

func (s *UserHandler) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (s *UserHandler) InputURL(c *gin.Context) {
	
	var url model.UrlStruct
	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(400, gin.H{
			"msg":   "Unable to bing json",
			"error": err,
		})
		return
	}
	var savetodb model.Url
	savetodb.LongUrl = shared.Encodetobase64(url.Url)
	Index+=1
	count := strconv.Itoa(Index)
	shortUrl := "http://localhost:8080/short" + count
	savetodb.ShortUrl = shortUrl
	if err := s.Config.Db.Create(&savetodb).Error; err != nil {
		c.JSON(502, gin.H{
			"msg": "Create db operation failed",
			"err": err,
		})
		return
	}

	c.JSON(200, gin.H{"url": shortUrl})
}

func (s *UserHandler) Welcome(c *gin.Context) {
	c.JSON(200, "Welcome Home")
}
func (s *UserHandler) Welcome1(c *gin.Context) {
	c.JSON(200, "Welcome Home1")
}
func (s *UserHandler) Welcome2(c *gin.Context) {
	c.JSON(200, "Welcome Home2")
}


func (s *UserHandler) OutputURL(c *gin.Context) {
var urlFromBody model.UrlStruct
	if err := c.ShouldBindJSON(&urlFromBody); err != nil {
		fmt.Println("Error binding data",err)
		return
	}
    var longUrl model.Url
	if err := s.Config.Db.Where("short_url=?",urlFromBody.Url).First(&longUrl).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":"Invalid URL",
			"error": err,
		})
		return
	}
	url,err := shared.Decodetobase64(longUrl.LongUrl)
	if err != nil {
		fmt.Println("Error decoding the string",err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	path := strings.TrimPrefix(url,"http://localhost:8080")
	c.Redirect(http.StatusFound,path)
}
