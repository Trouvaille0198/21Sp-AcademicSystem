package service

import (
	"academic-system/model"
	"errors"
	"gorm.io/gorm"
)

// GetSelectedCourses 获取指定学生的所有课程
func GetSelectedCourses(studentID int) (*[]model.SelectedCourseRes, error) {
	var courseByStuResult []model.SelectedCourseRes
	err := db.Raw("select distinct oc.id as offered_course_id, se.id as selection_id, oc.term, c.name, c.number, c.credit, t.name as teacher_name, "+
		"s.name as student_name, d.name as department, se.score "+
		"from offered_course as oc, course as c, selection as se, student as s, teacher as t, department as d "+
		"where oc.teacher_id = t.id and "+
		"oc.course_id = c.id and c.department_id = d.id and se.offered_course_id = oc.id and "+
		"se.student_id = s.id "+
		"and s.id = ?", studentID).Scan(&courseByStuResult).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &courseByStuResult, nil
}

// GetSelectedCoursesWithScore 获取指定学生的所有有成绩的课程
func GetSelectedCoursesWithScore(studentID int) (*[]model.SelectedCourseRes, error) {
	var courseByStuResult []model.SelectedCourseRes
	err := db.Raw("select distinct oc.id as offered_course_id, se.id as selection_id, oc.term, c.name, c.number, c.credit, t.name as teacher_name, "+
		"s.name as student_name, d.name as department, se.score "+
		"from offered_course as oc, course as c, selection as se, student as s, teacher as t, department as d "+
		"where oc.teacher_id = t.id and "+
		"oc.course_id = c.id and c.department_id = d.id and se.offered_course_id = oc.id and "+
		"se.student_id = s.id "+
		"and s.id = ? and se.score <> -1", studentID).Scan(&courseByStuResult).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &courseByStuResult, nil
}

// GetSelectedCoursesWithoutScore 获取指定学生的所有无成绩的课程
func GetSelectedCoursesWithoutScore(studentID int) (*[]model.SelectedCourseRes, error) {
	var courseByStuResult []model.SelectedCourseRes
	err := db.Raw("select distinct oc.id as offered_course_id, se.id as selection_id, oc.term, c.name, c.number, c.credit, t.name as teacher_name, "+
		"s.name as student_name, d.name as department, se.score "+
		"from offered_course as oc, course as c, selection as se, student as s, teacher as t, department as d "+
		"where oc.teacher_id = t.id and "+
		"oc.course_id = c.id and c.department_id = d.id and se.offered_course_id = oc.id and "+
		"se.student_id = s.id "+
		"and s.id = ? and se.score = -1", studentID).Scan(&courseByStuResult).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &courseByStuResult, nil
}

// GetOCsByTeacher 获取指定教师的所有开课课程
func GetOCsByTeacher(teacherID int) (*[]model.OfferedCourseRes, error) {
	var offeredCourseRes []model.OfferedCourseRes
	err := db.Raw("select distinct oc.id, oc.term, c.name, c.number, c.credit, t.name as teacher_name, "+
		"d.name as department "+
		"from offered_course as oc, course as c, teacher as t, department as d "+
		"where oc.teacher_id = t.id and "+
		"oc.course_id = c.id and c.department_id = d.id "+
		"and t.id = ?", teacherID).Scan(&offeredCourseRes).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &offeredCourseRes, nil
}
