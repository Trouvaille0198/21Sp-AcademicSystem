package service

import (
	"academic-system/model"
	"errors"
	"gorm.io/gorm"
	"strconv"
)

// CreateTeachersExample 创建教师样例
func CreateTeachersExample() (teachers []model.Teacher) {
	teachers = []model.Teacher{
		{Number: "0596", Name: "老师A", Sex: "男", Age: 41, DepartmentID: 1},
		{Number: "0597", Name: "老师B", Sex: "女", Age: 32, DepartmentID: 2},
		{Number: "0598", Name: "老师C", Sex: "男", Age: 50, DepartmentID: 3},
		{Number: "0599", Name: "老师D", Sex: "女", Age: 44, DepartmentID: 4},
		{Number: "0600", Name: "老师E", Sex: "男", Age: 38, DepartmentID: 1},
	}

	db.Model(&model.Teacher{}).Create(&teachers)
	return
}

// GetAllTeachersRes 获取所有教师信息
func GetAllTeachersRes() (*[]model.TeacherRes, error) {
	var teachers []model.TeacherRes
	err := db.Table("teacher").
		Select("teacher.id", "teacher.number", "teacher.name", "teacher.sex", "teacher.age", "department.name as department_name").
		Joins("left join department on teacher.department_id = department.id").
		Scan(&teachers).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &teachers, nil
}

// GetTeacherResByID 获取指定id的教师信息
func GetTeacherResByID(id int) (*model.TeacherRes, error) {
	var teacher model.TeacherRes
	err := db.Table("teacher").
		Select("teacher.id", "teacher.number", "teacher.name", "teacher.sex", "teacher.age", "department.name as department_name").
		Joins("left join department on teacher.department_id = department.id").
		Where("teacher.id = ?", id).
		Scan(&teacher).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &teacher, nil
}

// GetTeacherByNumber 获取指定工号的教师信息
func GetTeacherByNumber(number string) (*model.Teacher, error) {
	var teacher model.Teacher
	err := db.Model(&model.Teacher{}).Where("number = ?", number).First(&teacher).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &teacher, nil
}

// UpdateTeacher 通用方法 修改指定id的教师信息
func UpdateTeacher(id int, data map[string]interface{}) error {
	err := db.Model(&model.Teacher{}).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}

// CreateTeacher 创建教师实例
func CreateTeacher(teacher model.Teacher) (*model.Teacher, error) {
	err := db.Omit("selections").Create(&teacher).Error
	if err != nil {
		return &model.Teacher{}, err
	}
	return &teacher, nil
}

// DeleteTeacher 删除指定id教师
func DeleteTeacher(id int) error {
	res := db.Model(&model.Teacher{}).Delete(&model.Teacher{ID: uint(id)})
	if res.RowsAffected == 0 {
		return errors.New("删除失败, 不存在id为" + strconv.Itoa(id) + "的学生")
	}
	if res.Error != nil {
		return res.Error
	}
	return nil
}
