package router

import (
	"blog-backend/config"
	"blog-backend/entity"
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

type Login struct {
	Username string `form:"username" json:"username" xml:"username"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password"  binding:"required"`
}

func CheckLogin(c *gin.Context) {
	var loginParam Login
	if err := c.ShouldBind(&loginParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user entity.User
	if err := config.DB.Model(&user).Where("username = ?", loginParam.Username).Find(&user); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}
	checkPassword := GenerateMD5Hash(user.Salt, loginParam.Password)
	if user.Password == checkPassword {
		jwtStr := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userId":   user.ID,
			"username": user.Username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})
		c.JSON(http.StatusOK, gin.H{"data": jwtStr})
		return
	}
}

func GenerateMD5Hash(salt string, password string) string {
	// 将盐值和密码拼接成一个字符串
	data := salt + password

	// 创建一个 MD5 哈希对象
	hash := md5.New()

	// 将拼接后的字符串写入哈希对象
	hash.Write([]byte(data))

	// 获取哈希结果的字节切片
	hashBytes := hash.Sum(nil)

	// 将哈希结果的字节切片转换为十六进制字符串
	hashString := hex.EncodeToString(hashBytes)

	return hashString
}
