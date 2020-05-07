package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main()  {
	router := gin.Default()
	router.Use()

	// 指定地址和端口号
	err := router.Run("localhost:9090")
	logrus.Errorf("服务启动错误【%v】", err)
}