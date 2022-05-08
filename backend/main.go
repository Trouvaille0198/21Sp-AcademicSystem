package main

import (
	"academic-system/model"
	"academic-system/router"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title           Academic System
// @version         1.0
// @description     教务管理系统

// @BasePath  /api/v1
func main() {
	model.Setup() // 初始化gorm数据库连接

	r := router.NewRouter()
	_ = r.Run(":8080")
}
