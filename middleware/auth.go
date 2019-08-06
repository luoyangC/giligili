package middleware

import (
	"log"
	"giligili/util"
	"giligili/serializer"

	"github.com/gin-gonic/gin"
)

// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(200, serializer.Response{
				Status: 4003,
				Msg:    "需要登录",
			})
			c.Abort()
		}
		log.Print("get token: ", token)
		j := util.NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			c.JSON(200, serializer.Response{
				Status: 5000,
				Error:  err.Error(),
			})
			c.Abort()
		}
		c.Set("claims", claims)
	}
}
