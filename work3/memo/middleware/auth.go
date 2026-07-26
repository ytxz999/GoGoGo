package middleware

import (
	utils "memo/utils/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{
				"msg": "没有token",
			})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(401, gin.H{
				"msg": "token格式错误",
			})
			c.Abort()
			return
		}
		token := parts[1]
		userId, err := utils.ParseToken(token)
		if err != nil {
			c.JSON(401, gin.H{
				"msg": "token错误",
			})
			c.Abort()
			return
		}
		c.Set("userId", userId)
		c.Next()
	}
}
