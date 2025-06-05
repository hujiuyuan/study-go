package router

import "github.com/gin-gonic/gin"

func GetPosts(c *gin.Context) {
	// 查询 当前登录用户的 文章
}

func CreatePost(c *gin.Context) {
	// 写一篇文章
}

func DeletePost(c *gin.Context) {
	// 删除一篇文章
}

func UpdatePost(c *gin.Context) {
	// 更新一篇文章
}
