package utils

import (
	"academic-system/config"
	"academic-system/global"
	"github.com/spf13/viper"
	"log"
)

func InitConfig() {
	// 实例化viper
	v := viper.New()
	// 设置配置文件的路径
	v.SetConfigFile("./setting.yml")
	err := v.ReadInConfig()
	if err != nil {
		log.Println(err)
	}
	serverConfig := config.ServerConfig{}
	// 给serverConfig初始值
	err = v.Unmarshal(&serverConfig)
	if err != nil {
		log.Println(err)
	}
	// 传递给全局变量
	global.Settings = serverConfig
}
