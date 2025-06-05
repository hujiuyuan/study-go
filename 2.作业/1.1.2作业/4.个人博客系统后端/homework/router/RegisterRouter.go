package router

import (
	"blog-backend/config"
	"blog-backend/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type RegisterParam struct {
	Username string `form:"username" json:"username" xml:"username"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password"  binding:"required"`
	Email    string `form:"email" json:"email" xml:"email"  binding:"email"`
}

func Register(c *gin.Context) {
	var registerParam RegisterParam
	if err := c.ShouldBind(&registerParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user entity.User
	var sameUserCount int64
	if err := config.DB.Model(&user).Where("username = ?", registerParam.Username).Count(&sameUserCount); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "unauthorized"})
		return
	}

	if sameUserCount > 0 {
		c.JSON(http.StatusForbidden, gin.H{"msg": "已存在该账号"})
		return
	}
	salt := uuid.New().String()
	user.Username = registerParam.Username
	user.Password = GenerateMD5Hash(salt, registerParam.Password)
	user.Salt = salt
	user.Email = registerParam.Email
	config.DB.Model(&entity.User{}).Save(&user)

	c.JSON(http.StatusOK, gin.H{"msg": "success"})
	return
}
