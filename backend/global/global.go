// Package global 全局变量
package global

import (
	"academic-system/config"
	"gorm.io/gorm"
)

var (
	Settings config.ServerConfig
	DB       *gorm.DB
)
