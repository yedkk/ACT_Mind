package middleware

import (
	"net/http"
	"strings"

	"act-mind-backend/config"
	"act-mind-backend/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从Header获取token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "缺少Authorization头",
			})
			c.Abort()
			return
		}

		// 检查Bearer前缀
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization头格式错误",
			})
			c.Abort()
			return
		}

		// 解析token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.AppConfig.JWTSecret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "无效的token",
			})
			c.Abort()
			return
		}

		// 提取用户信息
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			userID := claims["user_id"]
			openID := claims["openid"]
			
			c.Set("user_id", userID)
			c.Set("openid", openID)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "token解析失败",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// GetCurrentUserID 从上下文获取当前用户ID
func GetCurrentUserID(c *gin.Context) (uint, bool) {
	userID, exists := c.Get("user_id")
	if !exists {
		return 0, false
	}
	
	// 处理不同类型的userID
	switch v := userID.(type) {
	case float64:
		return uint(v), true
	case uint:
		return v, true
	case int:
		return uint(v), true
	default:
		return 0, false
	}
}

// GetCurrentOpenID 从上下文获取当前用户OpenID
func GetCurrentOpenID(c *gin.Context) (string, bool) {
	openID, exists := c.Get("openid")
	if !exists {
		return "", false
	}
	
	if str, ok := openID.(string); ok {
		return str, true
	}
	
	return "", false
}