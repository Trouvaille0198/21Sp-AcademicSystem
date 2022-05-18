package service

import (
	"academic-system/model"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

// GetDepartmentByID 根据id获取课程信息
func GetDepartmentByID(id int) (*model.DepartmentRes, error) {
	var department model.DepartmentRes
	var (
		studentNum int
		teacherNum int
		courseNum  int
	)
	_ = db.Raw("SELECT count(*) FROM student WHERE department_id = ?", id).Scan(&studentNum).Error
	_ = db.Raw("SELECT count(*) FROM teacher WHERE department_id = ?", id).Scan(&teacherNum).Error
	_ = db.Raw("SELECT count(*) FROM course WHERE department_id = ?", id).Scan(&courseNum).Error
	_ = db.Model(&model.Department{}).Where("id = ?", id).Scan(&department)
	department.StudentNum = studentNum
	department.TeacherNum = teacherNum
	department.CourseNum = courseNum
	fmt.Printf("%+v", department)
	return &department, nil
}

// GetDepartments 获取所有课程信息
func GetDepartments() (*[]model.DepartmentRes, error) {
	var oriDep []model.Department
	err := db.Model(&model.Department{}).Scan(&oriDep).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	var departments []model.DepartmentRes
	for _, v := range oriDep {
		var (
			studentNum int
			teacherNum int
			courseNum  int
		)
		_ = db.Raw("SELECT count(*) FROM student WHERE department_id = ?", v.ID).Scan(&studentNum).Error
		_ = db.Raw("SELECT count(*) FROM teacher WHERE department_id = ?", v.ID).Scan(&teacherNum).Error
		_ = db.Raw("SELECT count(*) FROM course WHERE department_id = ?", v.ID).Scan(&courseNum).Error
		departments = append(departments, model.DepartmentRes{
			ID:         v.ID,
			Name:       v.Name,
			StudentNum: studentNum,
			TeacherNum: teacherNum,
			CourseNum:  courseNum,
		})
	}
	return &departments, nil
}
