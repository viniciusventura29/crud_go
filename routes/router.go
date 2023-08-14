package routes

import (
	"crud/controllers"

	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine) *gin.RouterGroup {
	tweetController := controllers.NewTweetController()

	v1 := router.Group("/")
	{
		v1.GET("/tweets", tweetController.FindAll)
		v1.POST("/tweets", tweetController.Create)
		v1.POST("/deleteTweets/:id", tweetController.Delete)
	}

	return v1
}
