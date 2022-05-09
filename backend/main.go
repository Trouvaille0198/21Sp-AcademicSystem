package main

import (
	"academic-system/config"
	"academic-system/global"
	"academic-system/initialize"
	"academic-system/router"
	"academic-system/utils"
	"github.com/spf13/viper"
	"log"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title           Academic System
// @version         1.0
// @description     教务管理系统
// @BasePath  /api/v1

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
func main() {
	utils.InitConfig()
	initialize.Setup() // 初始化gorm数据库连接

	r := router.NewRouter()
	_ = r.Run(":8080")
}
