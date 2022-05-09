package service

import (
	"academic-system/model"
	"errors"
	"gorm.io/gorm"
)

// GetOCsByStu 获取指定学生的所有课程
func GetOCsByStu(studentID int) (*[]model.OfferedCourseRes, error) {
	var courseByStuResult []model.OfferedCourseRes
	err := db.Raw("select distinct oc.id, oc.term, c.name, c.number, c.credit, t.name as teacher_name, "+
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

// GetOCsByStuWithScore 获取指定学生的所有有成绩的课程
func GetOCsByStuWithScore(studentID int) (*[]model.OfferedCourseRes, error) {
	var courseByStuResult []model.OfferedCourseRes
	err := db.Raw("select distinct oc.id, oc.term, c.name, c.number, c.credit, t.name as teacher_name, "+
		"s.name as student_name, d.name as department, se.score "+
		"from offered_course as oc, course as c, selection as se, student as s, teacher as t, department as d "+
		"where oc.teacher_id = t.id and "+
		"oc.course_id = c.id and c.department_id = d.id and se.offered_course_id = oc.id and "+
		"se.student_id = s.id "+
		"and s.id = ? and sc.score <> -1", studentID).Scan(&courseByStuResult).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &courseByStuResult, nil
}

// GetOCsByStuWithoutScore 获取指定学生的所有无成绩的课程
func GetOCsByStuWithoutScore(studentID int) (*[]model.OfferedCourseRes, error) {
	var courseByStuResult []model.OfferedCourseRes
	err := db.Raw("select distinct oc.id, oc.term, c.name, c.number, c.credit, t.name as teacher_name, "+
		"s.name as student_name, d.name as department, se.score "+
		"from offered_course as oc, course as c, selection as se, student as s, teacher as t, department as d "+
		"where oc.teacher_id = t.id and "+
		"oc.course_id = c.id and c.department_id = d.id and se.offered_course_id = oc.id and "+
		"se.student_id = s.id "+
		"and s.id = ? and sc.score = -1", studentID).Scan(&courseByStuResult).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &courseByStuResult, nil
}
