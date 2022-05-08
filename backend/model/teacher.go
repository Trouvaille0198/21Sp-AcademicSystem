package model

import (
	"errors"
	"gorm.io/gorm"
)

type Teacher struct {
	gorm.Model
	Number   string `json:"number" form:"number" gorm:"unique;index" example:"0198"` // 教师工号
	Name     string `json:"name" form:"name" example:"王二"`
	Sex      string `json:"sex" form:"sex" example:"男"`
	Age      uint8  `json:"age" form:"age" example:"21"`
	Password string `json:"password" form:"password" gorm:"default:123" example:"123"`

	DepartmentID uint        `json:"department_id" form:"department_id"` // 所属院系
	Selections   []Selection `json:"selections"`                         // 选课情况
}

// CreateTeachersExample 创建教师样例
func CreateTeachersExample() (teachers []Teacher) {
	teachers = []Teacher{
		{Number: "0596", Name: "老师A", Sex: "男", Age: 41, DepartmentID: 1},
		{Number: "0597", Name: "老师B", Sex: "女", Age: 32, DepartmentID: 2},
		{Number: "0598", Name: "老师C", Sex: "男", Age: 50, DepartmentID: 3},
		{Number: "0599", Name: "老师D", Sex: "女", Age: 44, DepartmentID: 4},
		{Number: "0600", Name: "老师E", Sex: "男", Age: 38, DepartmentID: 1},
	}

	db.Model(&Teacher{}).Create(&teachers)
	return
}

// GetAllTeachers 获取所有教师信息
func GetAllTeachers() (*[]Teacher, error) {
	var teachers []Teacher
	err := db.Model(&Teacher{}).Find(&teachers).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &teachers, nil
}

// GetTeacherByID 获取指定id的教师信息
func GetTeacherByID(id int) (*Teacher, error) {
	var student Teacher
	err := db.Model(&Teacher{}).Where("id = ?", id).First(&student).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &student, nil
}

// GetTeacherByNumber 获取指定工号的教师信息
func GetTeacherByNumber(number string) (*Teacher, error) {
	var teacher Teacher
	err := db.Model(&Teacher{}).Where("number = ?", number).First(&teacher).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &teacher, nil
}

// GetTeacherByAttrs 通用方法 根据条件获取所有教师信息
func GetTeacherByAttrs(student Teacher) (*[]Teacher, error) {
	var teachers []Teacher
	err := db.Model(&Teacher{}).Where(student).Find(&teachers).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &teachers, nil
}

// UpdateTeacher 通用方法 修改指定id的教师信息
func UpdateTeacher(id int, data map[string]interface{}) error {
	err := db.Model(&Teacher{}).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}

// CreateTeacher 创建教师实例
func CreateTeacher(student Teacher) (*Teacher, error) {
	err := db.Omit("selections").Create(&student).Error
	if err != nil {
		return &Teacher{}, err
	}
	return &student, nil
}

// DeleteTeacherByNumber 删除指定工号教师
func DeleteTeacherByNumber(number int) error {
	err := db.Model(&Teacher{}).Where("number = ?", number).Delete(&Teacher{}).Error
	if err != nil {
		return err
	}
	return nil
}
