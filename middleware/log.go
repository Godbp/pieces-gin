package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/godbp/piece-gin-src/log"
	"io/ioutil"
	"net/http"
)

// Params 此中间件放在 setHandler之后
func Params() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		switch ctx.Request.Method {
		case http.MethodGet:
			log.InfoWithContext(ctx, "接口:[%s] method:[%s] params:[%s]", ctx.Request.URL, ctx.Request.Method, ctx.Request.PostForm)
		default:
			data, err := ioutil.ReadAll(ctx.Request.Body)
			if err != nil {
				log.InfoWithContext(ctx, "获取参数失败 接口:[%s] method:[%s] err:[%s]", ctx.Request.URL, ctx.Request.Method, err)
			}
			log.InfoWithContext(ctx, "接口:[%s] method:[%s] params:[%s]", ctx.Request.URL, ctx.Request.Method, string(data))
			ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		}
		ctx.Next()
	}
}
