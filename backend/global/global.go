// Package global ćšć±ćé
package global

import (
	"academic-system/config"
	"gorm.io/gorm"
)

var (
	Settings config.ServerConfig
	DB       *gorm.DB
)
