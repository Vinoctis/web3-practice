package middleware

import (
	"github.com/gin-gonic/gin"
	//"blog/internal/utils"
	"blog/internal/response"
	"blog/internal/utils"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Fail(c, "请提供认证token呗", nil)
			c.Abort()
			return
		}
		//检查 Bearer 前缀
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			response.Fail(c, "token格式错误", nil)
			c.Abort()
			return
		}

		//解析token
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			response.Fail(c, "token无效", nil)
			c.Abort()
			return
		}

		//将claims 存储到context中
		c.Set("claims", claims)

		c.Next()
	}
}