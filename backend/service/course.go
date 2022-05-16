package service

import (
	"academic-system/model"
	"errors"
	"gorm.io/gorm"
	"strconv"
)

// CreateCoursesExample 创建课程样例
func CreateCoursesExample() (courses []model.Course) {
	courses = []model.Course{
		{Number: "j0001", Name: "算法设计与分析", Credit: 4, DepartmentID: 1},
		{Number: "j0002", Name: "操作系统(1)", Credit: 5, DepartmentID: 1},
		{Number: "j1001", Name: "数据库原理(1)", Credit: 4, DepartmentID: 1},
		{Number: "j1002", Name: "操作系统(2)", Credit: 5, DepartmentID: 1},
		{Number: "j2001", Name: "数据库原理(2)", Credit: 4, DepartmentID: 1},
		{Number: "j2002", Name: "编译原理", Credit: 3, DepartmentID: 1},
		{Number: "t1001", Name: "数字电路", Credit: 5, DepartmentID: 2},
		{Number: "t1002", Name: "模拟电路", Credit: 5, DepartmentID: 2},
		{Number: "l0001", Name: "大学物理", Credit: 4, DepartmentID: 3},
		{Number: "w0001", Name: "大学英语", Credit: 3, DepartmentID: 4},
	}
	db.Model(&model.Course{}).Create(&courses)
	return
}

// GetCourses 获取所有课程信息
func GetCourses() ([]*model.CourseRes, error) {
	var courses []*model.CourseRes
	err := db.Table("course").
		Select("course.id", "course.number", "course.name", "course.credit", "department.name as department_name").
		Joins("left join department on course.department_id = department.id").
		Scan(&courses).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return courses, nil
}

// GetCourseByID 获取指定id的课程信息
func GetCourseByID(id int) (*model.CourseRes, error) {
	var course model.CourseRes
	err := db.Table("course").
		Select("course.id", "course.number", "course.name", "course.credit", "department.name as department_name").
		Joins("left join department on course.department_id = department.id").
		Where("course.id = ?", id).
		Scan(&course).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &course, nil
}

// CreateCourse 创建课程实例
func CreateCourse(course model.Course) (*model.Course, error) {
	err := db.Create(&course).Error
	if err != nil {
		return &model.Course{}, err
	}
	return &course, nil
}

// DeleteCourse 删除指定id课程
func DeleteCourse(id int) (err error) {
	res := db.Model(&model.Course{}).Delete(&model.Course{ID: uint(id)})
	if res.RowsAffected == 0 {
		return errors.New("删除失败, 不存在id为" + strconv.Itoa(id) + "的课程")
	}
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// DeleteAllCourses 删除所有课程
func DeleteAllCourses() {
	db.Where("1 = 1").Delete(&model.Course{})
}
