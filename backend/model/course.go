package model

import (
	"errors"
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	Number string `json:"number" form:"number" gorm:"index" example:"0121"` // 课号
	Name   string `json:"name" form:"name" example:"数据库原理"`                 // 课名
	Credit uint8  `json:"credit" form:"credit" example:"4"`                 // 学分

	DepartmentID uint `json:"department_id" form:"department_id"` // 院系号
}

// CourseRes ...
type CourseRes struct {
	ID          uint   `json:"id"`
	Score       int    `json:"score"`        // 成绩
	Number      string `json:"number"`       // 课号
	Name        string `json:"name"`         // 课名
	Credit      uint8  `json:"credit"`       // 学分
	Department  string `json:"department"`   // 所属院系
	Term        string `json:"term"`         // 学期
	TeacherName string `json:"teacher_name"` // 老师名
	StudentName string `json:"student_name"` // 学生名
}

// CreateCoursesExample 创建课程样例
func CreateCoursesExample() (courses []Course) {
	courses = []Course{
		{Number: "0121", Name: "数据库原理(1)", Credit: 4, DepartmentID: 1},
		{Number: "0122", Name: "数据库原理(2)", Credit: 4, DepartmentID: 1},
		{Number: "0830", Name: "操作系统(2)", Credit: 5, DepartmentID: 1},
		{Number: "0451", Name: "大学物理(3)", Credit: 2, DepartmentID: 2},
		{Number: "0279", Name: "算法设计与分析", Credit: 2, DepartmentID: 3},
		{Number: "0022", Name: "编译原理", Credit: 4, DepartmentID: 2},
		{Number: "0023", Name: "编译原理", Credit: 4, DepartmentID: 3},
		{Number: "0024", Name: "编译原理", Credit: 4, DepartmentID: 1},
	}
	db.Model(&Course{}).Create(&courses)
	return
}

// GetAllCourses 获取所有课程信息
func GetAllCourses() ([]*Course, error) {
	var courses []*Course
	err := db.Model(&Course{}).Find(&courses).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return courses, nil
}

// GetCourseByID 获取指定id的课程信息
func GetCourseByID(id int) (*Course, error) {
	var course Course
	err := db.Model(&Course{}).Where("id = ?", id).First(&course).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &course, nil
}

// GetCourseByAttrs 通用方法 根据条件获取所有课程信息
func GetCourseByAttrs(course Course) (*[]Course, error) {
	var courses []Course
	err := db.Model(&Course{}).Where(course).Find(&courses).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &courses, nil
}

// UpdateCourse 通用方法 修改指定id的课程信息
func UpdateCourse(id int, data map[string]interface{}) (err error) {
	err = db.Model(&Course{}).Omit("selections").Where("id = ?", id).Updates(data).Error
	return
}

// UpdateWholeCourse 通用方法 整个地修改指定id的课程信息
func UpdateWholeCourse(id int, course Course) (err error) {
	err = db.Model(&Course{}).Omit("selections").Where("id = ?", id).Updates(course).Error
	return
}

// CreateCourse 创建课程实例
func CreateCourse(course Course) (*Course, error) {
	err := db.Omit("selections").Create(&course).Error
	if err != nil {
		return &Course{}, err
	}
	return &course, nil
}

// DeleteCourse 删除指定id课程
func DeleteCourse(id int) (err error) {
	err = db.Model(&Course{}).Where("id = ?", id).Delete(&Course{}).Error
	return
}
