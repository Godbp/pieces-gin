package middleware

import "github.com/gin-gonic/gin"

func Token() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}

}