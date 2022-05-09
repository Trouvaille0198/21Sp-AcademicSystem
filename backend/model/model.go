package model

import (
	"academic-system/global"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

// InitDB 在全局DB初始化完成后，赋值给service包的db变量
func InitDB() {
	db = global.DB
}
