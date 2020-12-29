package middleware

import "github.com/gin-gonic/gin"

// Token Authorization
func Token() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sid := ctx.GetHeader("Authorization")
		if len(sid) == 0 {
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
