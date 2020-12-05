package middleware

import "github.com/gin-gonic/gin"
import "github.com/satori/go.uuid"

func Token() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置请求跟踪ID
		reqID := c.Writer.Header().Get("X-Request-Id")
		if len(reqID) == 0 {
			reqID = uuid.NewV4().String()
			c.Writer.Header().Set("X-Request-Id", reqID)
		}
		c.Next()
	}
}
