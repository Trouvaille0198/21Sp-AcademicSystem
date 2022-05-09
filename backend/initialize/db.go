package initialize

import (
	"academic-system/global"
	"academic-system/model"
	"academic-system/service"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var db *gorm.DB

// Setup 初始化数据库连接
func Setup() {
	var err error
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			global.Settings.GormInfo.Name,
			global.Settings.GormInfo.Password,
			global.Settings.GormInfo.Host,
			global.Settings.GormInfo.Port,
			global.Settings.GormInfo.DBName), // DSN data source name
		DefaultStringSize:        171,
		DisableDatetimePrecision: true,
		DontSupportRenameIndex:   true,
	}), &gorm.Config{
		SkipDefaultTransaction:                   false,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatalf("model.Setup err: %v", err)
	}

	err = db.AutoMigrate(
		&model.Department{}, &model.Course{}, &model.Teacher{}, &model.OfferedCourse{},
		&model.Student{}, &model.Selection{})
	if err != nil {
		log.Fatalf("model.Setup err: %v", err)
	}
	global.DB = db
	service.InitDB()
	model.InitDB()
}
