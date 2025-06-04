package router

import "github.com/gin-gonic/gin"

var GinRouter *gin.Engine

func init() {
	GinRouter := gin.Default()

	GinRouter.POST("/login", CheckLogin)
}
