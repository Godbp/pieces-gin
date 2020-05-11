package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func Params() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodGet{
			logrus.Infof("接口：【%v】 method：【%v】 params： 【%v】",c.Request.URL, c.Request.Method, c.Request.PostForm)
		} else {
			data, _ := ioutil.ReadAll(c.Request.Body)
			logrus.Infof("接口：【%v】 method：【%v】 params： 【%v】",c.Request.URL, c.Request.Method, string(data))
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		}
		c.Next()
	}
}