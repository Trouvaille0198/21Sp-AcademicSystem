package model

import (
	"errors"
	"gorm.io/gorm"
)

// Selection 记录学生选课情况
// 三种情况: 未选 已选 已修
type Selection struct {
	ID              uint `gorm:"primarykey"`
	StudentID       uint `gorm:"student_id"`
	OfferedCourseID uint `gorm:"offered_course_id"`

	Student       Student
	OfferedCourse OfferedCourse
	Score         int `gorm:"default:-1"` // 总分, -1表示未评分
	UsualScore    int `gorm:"default:-1"` // 平时成绩, -1表示未评分
	ExamScore     int `gorm:"default:-1"` // 考试成绩, -1表示未评分
}

// CreateSelectionsExample 创建课程关联实例
func CreateSelectionsExample() (selections []Selection) {
	selections = []Selection{
		{StudentID: 1, OfferedCourseID: 1, Score: -1, UsualScore: -1, ExamScore: -1},
		{StudentID: 1, OfferedCourseID: 2, Score: -1, UsualScore: -1, ExamScore: -1},
		{StudentID: 1, OfferedCourseID: 4, Score: 75, UsualScore: 75, ExamScore: 75},
		{StudentID: 1, OfferedCourseID: 3, Score: -1, UsualScore: -1, ExamScore: -1},
		{StudentID: 1, OfferedCourseID: 6, Score: -1, UsualScore: -1, ExamScore: -1},

		{StudentID: 2, OfferedCourseID: 1, Score: -1, UsualScore: -1, ExamScore: -1},
		{StudentID: 2, OfferedCourseID: 3, Score: -1, UsualScore: -1, ExamScore: -1},
		{StudentID: 2, OfferedCourseID: 4, Score: 93, UsualScore: 93, ExamScore: 93},
		{StudentID: 2, OfferedCourseID: 5, Score: 70, UsualScore: 70, ExamScore: 70},

		{StudentID: 3, OfferedCourseID: 4, Score: 68, UsualScore: 68, ExamScore: 68},
		{StudentID: 1, OfferedCourseID: 6, Score: -1, UsualScore: -1, ExamScore: -1},

		{StudentID: 5, OfferedCourseID: 4, Score: 85, UsualScore: 85, ExamScore: 85},
	}

	db.Model(&Selection{}).Create(&selections)
	return
}

// CreateSelection 创建选课记录
func CreateSelection(selection Selection) (*Selection, error) {
	// 判断数据库中是否存在重复选课
	var duplicatedResults []*Selection
	rowsAffected := db.Model(&Selection{}).Where(
		"student_id = ? AND offered_course_id = ?",
		selection.StudentID, selection.OfferedCourseID).Find(&duplicatedResults).RowsAffected
	if rowsAffected > 0 {
		return nil, errors.New("选课关系已存在！")
	}

	// TODO 判断学生是否已选课名、学期均相同的课程
	//targetCourse, err := GetOfferedCourseByID(int(selection.OfferedCourseID))
	//if err != nil {
	//	return nil, err
	//}
	//duplicatedCourses, err := GetCoursesByStudent(int(selection.StudentID))
	//if err != nil {
	//	return nil, err
	//}
	//for _, course := range *duplicatedCourses {
	//	if course.ServerName == targetCourse.ServerName && course.Term == targetCourse.Term {
	//		return nil, errors.New("学生已选课程名、学期均相同的课程！")
	//	}
	//}

	err := db.Model(&Selection{}).Create(&selection).Error
	if err != nil {
		return &Selection{}, err
	}
	return &selection, nil
}

// UpdateSelection 更新选课记录
func UpdateSelection(id int, data map[string]interface{}) (err error) {
	err = db.Model(&Selection{}).Where("id = ?", id).Updates(data).Error
	return
}

// DeleteSelection 删除指定id选课记录
func DeleteSelection(selection Selection) error {
	result := db.Model(&Selection{}).Where(
		"student_id = ? AND offered_course_id = ?", selection.StudentID, selection.OfferedCourseID).Unscoped().Delete(&Selection{})
	err := result.Error
	rowsAffected := result.RowsAffected
	if err != nil {
		return err
	} else if rowsAffected == 0 {
		return errors.New("没有此选课记录！")
	}
	return nil
}

// GetSelectionsByStudentID 获取指定学生的选课记录
func GetSelectionsByStudentID(studentID int) ([]*Selection, error) {
	var selections []*Selection
	err := db.Model(&Selection{}).Where("student_id = ?", studentID).Find(&selections).Error
	if err != nil {
		return []*Selection{}, err
	}
	return selections, nil
}

// GetSelection 根据学生id和开课课程id获取选课记录
func GetSelection(studentID, offeredCourseID int) (*Selection, error) {
	var selection Selection
	err := db.Model(&Selection{}).Where(
		"student_id = ? AND offered_course_id = ?", studentID, offeredCourseID).Find(&selection).Error
	if err != nil || errors.Is(err, gorm.ErrRecordNotFound) {
		return &Selection{}, err
	}
	return &selection, nil
}

// UpdateSelectionScore 更新课程成绩
func UpdateSelectionScore(id uint, score int) error {
	if score < 1 || score > 100 {
		return errors.New("分数不合法！")
	}

	err := db.Model(&Selection{}).Where("id = ?", id).Updates(Selection{Score: score}).Error
	return err
}

// DeleteAllSelections 删除所有选课记录
func DeleteAllSelections() {
	db.Where("1 = 1").Delete(&Selection{})
}
