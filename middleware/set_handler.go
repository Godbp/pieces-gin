package middleware

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// SetHandler handler参数设置
func SetHandler() gin.HandlerFunc {
	// 设置请求跟踪ID
	return func(ctx *gin.Context) {
		reqID := ctx.Writer.Header().Get("X-Request-Id")
		if len(reqID) == 0 {
			reqID = uuid.NewV4().String()
			ctx.Writer.Header().Set("X-Request-Id", reqID)
		}
		ctx.Next()
	}
}
