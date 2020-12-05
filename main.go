package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pieces-gin/models"
	"github.com/sirupsen/logrus"
	"github.com/pieces-gin/middleware"
	"github.com/pieces-gin/config"
)

func main()  {

	// init config
	conf := config.LoadConf()
	models.Mysql = conf.Mysql

	router := gin.Default()
	// 跨域请求
	router.Use(middleware.Cors())
	// token 认证
	router.Use(middleware.Token())
	// 参数日志打印
	router.Use(middleware.Params())

	// 指定地址和端口号
	err := router.Run("localhost:9090")
	if err != nil{
		logrus.Errorf("服务启动错误【%v】", err)
	}

}