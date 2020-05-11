package config

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Debug          bool          `yaml:"debug"`
	Client         *ClientConf   `yaml:"client"`
	Redis          *RedisConf    `yaml:"redis"`
	Mysql          *MysqlConf    `yaml:"mysql"`
	RedisKey       *RedisKeyType `yaml:"RedisKeyType"`
}



type ClientConf struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func (s *ClientConf) GetServerAddr() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}

type RedisConf struct {
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
	PW     string `yaml:"password"`
	Db     int    `yaml:"db"`
}

type MysqlConf struct {
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
	PW     string `yaml:"password"`
	User   string `yaml:"user"`
	Db     string `yaml:"db_name"`
	Conns  int    `yaml:"conns"`
	Debug  bool   `yaml:"debug"`
}

type RedisKeyType struct {
	Token     string `yaml:"token"`
	Request   string `yaml:"request"`
}

func LoadConf() *Config {
	var err error
	var port string
	var cfile string

	flag.StringVar(&cfile, "c", "", "-c /config/file/path")
	flag.StringVar(&port, "p", "", "-p 8200")

	flag.Parse()

	if cfile == "" {
		fmt.Println("缺少运行的配置文件, 使用 -c /config/file/path")
		os.Exit(1)
	}

	viper.SetConfigFile(cfile)
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("解析配置文件失败", err.Error())
		os.Exit(1)
	}

	var c Config

	err = viper.Unmarshal(&c)
	if err != nil {
		fmt.Println("反序列化配置文件失败", err.Error())
		os.Exit(1)
	}

	return &c
}
