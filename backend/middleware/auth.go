package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"student-sys/utils"
)

// AuthMiddleware JWT鉴权中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头获取token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// 检查token格式
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// 解析token
		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// 将用户信息存储到上下文
		c.Set("userID", claims.UserID)
		c.Set("role", claims.Role)

		c.Next()
	}
}

// TeacherAuthMiddleware 教师角色鉴权中间件
func TeacherAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从上下文获取角色信息
		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// 检查角色是否为教师
		if role != "teacher" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: Only teachers can access this resource"})
			c.Abort()
			return
		}

		c.Next()
	}
}
