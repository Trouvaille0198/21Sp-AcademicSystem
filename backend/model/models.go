package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB // 全局db对象

// Setup 初始化数据库连接
func Setup() {
	var err error
	db, err = gorm.Open(sqlite.Open("example.db"), &gorm.Config{})

	if err != nil {
		log.Fatalf("model.Setup err: %v", err)
	}

	err = db.AutoMigrate(&Student{}, &Course{}, &Selection{})
	if err != nil {
		log.Fatalf("model.Setup err: %v", err)
	}
}
