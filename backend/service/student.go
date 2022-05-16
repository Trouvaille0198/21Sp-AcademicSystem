package service

import (
	"academic-system/model"
	"errors"
	"gorm.io/gorm"
	"strconv"
)

// CreateStudentsExample 创建学生样例
func CreateStudentsExample() (students []model.Student) {
	students = []model.Student{
		{Number: "0196", Name: "学生A", Sex: "男", Age: 21, DepartmentID: 1},
		{Number: "0197", Name: "学生B", Sex: "女", Age: 22, DepartmentID: 1},
		{Number: "0198", Name: "学生C", Sex: "男", Age: 20, DepartmentID: 1},
		{Number: "0199", Name: "学生D", Sex: "女", Age: 21, DepartmentID: 1},
		{Number: "1200", Name: "学生E", Sex: "男", Age: 21, DepartmentID: 2},
	}

	db.Model(&model.Student{}).Create(&students)
	return
}

// GetAllStudentsRes 获取所有学生信息
func GetAllStudentsRes() (*[]model.StudentRes, error) {
	var students []model.StudentRes
	err := db.Table("student").
		Select("student.id", "student.number", "student.name", "student.sex", "student.age", "department.name as department_name").
		Joins("left join department on student.department_id = department.id").
		Scan(&students).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &students, nil
}

// GetStudentResByID 获取指定id的学生信息
func GetStudentResByID(id int) (*model.StudentRes, error) {
	var student model.StudentRes
	err := db.Model(&model.Student{}).
		Select("student.id", "student.number", "student.name", "student.sex", "student.age", "department.name as department_name").
		Joins("left join department on student.department_id = department.id").
		Where("student.id = ?", id).
		Scan(&student).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &student, nil
}

// GetStudentByNumber 获取指定学号的学生记录
func GetStudentByNumber(number string) (*model.Student, error) {
	var student model.Student
	err := db.Model(&model.Student{}).Where("number = ?", number).First(&student).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &student, nil
}

// UpdateStudent 通用方法 修改指定id的学生信息
func UpdateStudent(id int, data map[string]interface{}) error {
	err := db.Model(&model.Student{}).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}

// CreateStudent 创建学生实例
func CreateStudent(student model.Student) (*model.Student, error) {
	err := db.Model(&model.Student{}).Create(&student).Error
	if err != nil {
		return &model.Student{}, err
	}
	return &student, nil
}

// DeleteStudent 删除指定id学生
func DeleteStudent(id int) error {
	res := db.Model(&model.Student{}).Delete(&model.Student{ID: uint(id)})
	if res.RowsAffected == 0 {
		return errors.New("删除失败, 不存在id为" + strconv.Itoa(id) + "的学生")
	}
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// DeleteAllStudents 删除所有学生
func DeleteAllStudents() {
	db.Where("1 = 1").Delete(&model.Student{})
}

func GetStudentsByOfferedCourse(id int) (*[]model.StudentWithScoreRes, error) {
	var studentsRes []model.StudentWithScoreRes
	err := db.Raw("select distinct s.id, s.name, s.number, s.sex,s.age, de.name as department_name, "+
		"se.score, se.usual_score, se.exam_score "+
		"from student as s, selection as se, department as de "+
		"where se.offered_course_id = ? and s.id = se.student_id and de.id = s.department_id ", id).
		Scan(&studentsRes).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &studentsRes, nil
}
