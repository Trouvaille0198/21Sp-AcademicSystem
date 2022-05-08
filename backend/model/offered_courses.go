package model

import (
	"errors"
	"gorm.io/gorm"
)

// OfferedCourse 开课信息
type OfferedCourse struct {
	gorm.Model
	Term string `json:"term" form:"term" example:"22-冬季学期"` // 学期

	CourseID  uint `json:"course_id" form:"course_id"`   // 课号
	TeacherID uint `json:"teacher_id" form:"teacher_id"` // 教师工号

	Selections []Selection `json:"selections"`
}

// CreateOfferedCoursesExample 创建课程样例
func CreateOfferedCoursesExample() (offeredCourses []OfferedCourse) {
	offeredCourses = []OfferedCourse{
		{Term: "21-冬季学期", TeacherID: 1, CourseID: 1},
		{Term: "22-春季学期", TeacherID: 3, CourseID: 3},
		{Term: "21-冬季学期", TeacherID: 2, CourseID: 4},
		{Term: "21-春季学期", TeacherID: 2, CourseID: 2},
		{Term: "21-秋季学期", TeacherID: 1, CourseID: 1},
		{Term: "22-春季学期", TeacherID: 2, CourseID: 6},
		{Term: "22-春季学期", TeacherID: 3, CourseID: 8},
		{Term: "22-春季学期", TeacherID: 1, CourseID: 7},
	}

	db.Model(&OfferedCourse{}).Create(&offeredCourses)
	return
}

func GetOfferedCourseByID(id int) (*OfferedCourse, error) {
	var offeredCourse OfferedCourse
	err := db.Model(&OfferedCourse{}).Where("id = ?", id).First(&offeredCourse).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &offeredCourse, nil
}
