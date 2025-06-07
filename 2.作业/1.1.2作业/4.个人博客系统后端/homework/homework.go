package main

import (
	"blog-backend/router"
	"blog-backend/router/business"
	"github.com/gin-gonic/gin"
)

var GinRouter *gin.Engine

func main() {
	GinRouter := gin.Default()

	GinRouter.POST("/login", router.LoginFunc)
	GinRouter.POST("/register", router.Register)

	group := GinRouter.Group("/business")
	group.Use(router.CheckLogin)

	group.GET("/post/:userId", business.GetPosts)
	GinRouter.Run()
}
