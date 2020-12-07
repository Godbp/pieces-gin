package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/godbp/piece-gin-src/log"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
)

// Params 此中间件放在
func Params() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 设置请求跟踪ID
		reqID := ctx.Writer.Header().Get("X-Request-Id")
		if len(reqID) == 0 {
			reqID = uuid.NewV4().String()
			ctx.Writer.Header().Set("X-Request-Id", reqID)
		}
		if ctx.Request.Method == http.MethodGet {
			log.InfoWithContext(ctx, "接口:[%s] method:[%s] params:[%s]", ctx.Request.URL, ctx.Request.Method, ctx.Request.PostForm)
		} else {
			data, _ := ioutil.ReadAll(ctx.Request.Body)
			log.InfoWithContext(ctx, "接口:[%s] method:[%s] params:[%s]", ctx.Request.URL, ctx.Request.Method, string(data))
			ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		}
		ctx.Next()
	}
}
