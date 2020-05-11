package models

import (
	"fmt"
	"github.com/pieces-gin/config"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

var (
	maxConns = 20
)

var Mysql  *config.MysqlConf

type GormGin struct {
	DB        *gorm.DB
}

func (gm *GormGin)Init() *gorm.DB {
	mysqlUrl := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Mysql.User, Mysql.PW, Mysql.Host, Mysql.Port, Mysql.Db)
	db, err := gorm.Open("mysql", mysqlUrl)
	if err != nil{
		logrus.Errorf("连接【mysql】【%s】失败【%v】",mysqlUrl, err)
	}
	// 设置mysql 默认连接数
	if Mysql.Conns == 0 {
		Mysql.Conns = maxConns
	}
	db.DB().SetMaxIdleConns(Mysql.Conns)
	db.DB().SetMaxOpenConns(Mysql.Conns)
	gm.DB = db
	db.LogMode(true)
	return db
}