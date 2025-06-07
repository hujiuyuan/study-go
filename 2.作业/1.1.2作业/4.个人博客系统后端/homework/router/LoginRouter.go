package router

import (
	"blog-backend/config"
	"blog-backend/entity"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

type Login struct {
	Username string `form:"username" json:"username" xml:"username"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password"  binding:"required"`
}

func LoginFunc(c *gin.Context) {
	var loginParam Login
	if err := c.ShouldBind(&loginParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user entity.User
	config.DB.Model(&user).Where("username = ?", loginParam.Username).Find(&user)
	checkPassword := GenerateMD5Hash(user.Salt, loginParam.Password)
	if user.Password == checkPassword {
		jwtStr := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userId":   user.ID,
			"username": user.Username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

		secretKey := "" // 替换为你的密钥
		tokenString, err := jwtStr.SignedString([]byte(secretKey))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "无法生成 JWT"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": tokenString})
		return
	}
}

func CheckLogin(c *gin.Context) {
	var token = c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "必须上传JWT签名信息"})
	}

	// 鉴权操作
	var loginParam = jwt.MapClaims{}
	jwtToken, err := jwt.ParseWithClaims(token, &loginParam, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(""), nil // 使用相同的密钥
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的 token"})
		return
	}
	if loginParam, ok := jwtToken.Claims.(*jwt.MapClaims); ok && jwtToken.Valid {
		fmt.Printf("%v", loginParam)
		c.Next()
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的 token"})
	}

	// 后续操作
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
