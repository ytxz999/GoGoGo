package middleware

import (
	"memo/common"
	utils "memo/utils/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			common.UnauFail(c, "没有token")
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			common.BadReqFail(c, "token格式错误")
			c.Abort()
			return
		}
		token := parts[1]
		userId, err := utils.ParseToken(token)
		if err != nil {
			common.UnauFail(c, "token错误")
			c.Abort()
			return
		}
		c.Set("userId", userId)
		c.Next()
	}
}
