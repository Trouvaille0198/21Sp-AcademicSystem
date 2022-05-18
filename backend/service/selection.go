package service

import (
	"academic-system/model"
	"errors"
	"gorm.io/gorm"
	"log"
)

// CreateSelectionsExample 创建课程关联实例
func CreateSelectionsExample() (selections []model.Selection) {
	selections = []model.Selection{
		{StudentID: 1, OfferedCourseID: 1, UsualScore: 70, ExamScore: 80, Score: 75},
		{StudentID: 1, OfferedCourseID: 2, UsualScore: 88, ExamScore: 90, Score: 89},
		{StudentID: 1, OfferedCourseID: 3, UsualScore: 60, ExamScore: 70, Score: 65},
		{StudentID: 1, OfferedCourseID: 4, UsualScore: 90, ExamScore: 88, Score: 89},
		{StudentID: 1, OfferedCourseID: 6, UsualScore: -1, ExamScore: -1, Score: -1},
		{StudentID: 1, OfferedCourseID: 8, UsualScore: 90, ExamScore: -1, Score: -1},
		{StudentID: 1, OfferedCourseID: 11, UsualScore: -1, ExamScore: -1, Score: -1},
		{StudentID: 1, OfferedCourseID: 13, UsualScore: -1, ExamScore: -1, Score: -1},

		{StudentID: 2, OfferedCourseID: 1, UsualScore: 66, ExamScore: 70, Score: 68},
		{StudentID: 2, OfferedCourseID: 2, UsualScore: 93, ExamScore: 87, Score: 90},
		{StudentID: 2, OfferedCourseID: 4, UsualScore: 78, ExamScore: 68, Score: 73},
		{StudentID: 2, OfferedCourseID: 6, UsualScore: -1, ExamScore: -1, Score: -1},
		{StudentID: 2, OfferedCourseID: 8, UsualScore: -1, ExamScore: -1, Score: -1},
		{StudentID: 2, OfferedCourseID: 13, UsualScore: -1, ExamScore: -1, Score: -1},

		{StudentID: 3, OfferedCourseID: 1, UsualScore: 76, ExamScore: 70, Score: 73},
		{StudentID: 3, OfferedCourseID: 2, UsualScore: 84, ExamScore: 87, Score: 86},
		{StudentID: 3, OfferedCourseID: 3, UsualScore: 68, ExamScore: 68, Score: 68},
		{StudentID: 3, OfferedCourseID: 5, UsualScore: 78, ExamScore: 68, Score: 73},
		{StudentID: 3, OfferedCourseID: 7, UsualScore: -1, ExamScore: -1, Score: -1},
		{StudentID: 3, OfferedCourseID: 9, UsualScore: -1, ExamScore: -1, Score: -1},
		{StudentID: 3, OfferedCourseID: 12, UsualScore: -1, ExamScore: -1, Score: -1},
		{StudentID: 3, OfferedCourseID: 14, UsualScore: -1, ExamScore: -1, Score: -1},

		{StudentID: 4, OfferedCourseID: 1, UsualScore: 86, ExamScore: 67, Score: 77},
		{StudentID: 4, OfferedCourseID: 2, UsualScore: 88, ExamScore: 78, Score: 83},
		{StudentID: 4, OfferedCourseID: 5, UsualScore: 79, ExamScore: 73, Score: 76},
		{StudentID: 4, OfferedCourseID: 7, UsualScore: -1, ExamScore: -1, Score: -1},
		{StudentID: 4, OfferedCourseID: 9, UsualScore: -1, ExamScore: -1, Score: -1},

		{StudentID: 5, OfferedCourseID: 1, UsualScore: 77, ExamScore: 83, Score: 80},
		{StudentID: 5, OfferedCourseID: 2, UsualScore: 84, ExamScore: 72, Score: 78},
		{StudentID: 5, OfferedCourseID: 10, UsualScore: 86, ExamScore: -1, Score: -1},
		{StudentID: 5, OfferedCourseID: 15, UsualScore: -1, ExamScore: -1, Score: -1},
	}

	db.Model(&model.Selection{}).Create(&selections)
	return
}

// CreateSelection 创建选课记录
func CreateSelection(selection model.Selection) (*model.Selection, error) {
	// 判断数据库中是否存在重复选课
	var duplicatedResults []*model.Selection
	rowsAffected := db.Model(&model.Selection{}).Where(
		"student_id = ? AND offered_course_id = ?",
		selection.StudentID, selection.OfferedCourseID).Find(&duplicatedResults).RowsAffected
	if rowsAffected > 0 {
		log.Printf("%+v \n", duplicatedResults[0])
		return nil, errors.New("选课关系已存在！")
	}

	// 判断学生是否已选课名、学期均相同的课程
	targetOc, err := GetOfferedCourseByID(int(selection.OfferedCourseID))
	if err != nil {
		return nil, errors.New("不存在该课程！")
	}
	coursesByStu, err := GetSelectedCoursesByStu(int(selection.StudentID))
	if err != nil {
		return nil, err
	}
	for _, course := range *coursesByStu {
		if course.Name == targetOc.Name && course.Term == targetOc.Term {
			return nil, errors.New("学生已选课程名、学期均相同的课程！")
		}
	}

	err = db.Model(&model.Selection{}).Create(&selection).Error
	if err != nil {
		return &model.Selection{}, err
	}
	return &selection, nil
}

// UpdateSelection 更新选课记录
func UpdateSelection(id int, data map[string]interface{}) (err error) {
	err = db.Model(&model.Selection{}).Where("id = ?", id).Updates(data).Error
	return
}

// DeleteSelection 删除指定id选课记录
func DeleteSelection(selection model.Selection) error {
	result := db.Model(&model.Selection{}).Where(
		"student_id = ? AND offered_course_id = ?", selection.StudentID, selection.OfferedCourseID).Unscoped().Delete(&model.Selection{})
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
func GetSelectionsByStudentID(studentID int) ([]*model.Selection, error) {
	var selections []*model.Selection
	err := db.Model(&model.Selection{}).Where("student_id = ?", studentID).Find(&selections).Error
	if err != nil {
		return []*model.Selection{}, err
	}
	return selections, nil
}

// GetSelection 根据学生id和开课课程id获取选课记录
func GetSelection(studentID, offeredCourseID int) (*model.Selection, error) {
	var selection model.Selection
	err := db.Model(&model.Selection{}).Where(
		"student_id = ? AND offered_course_id = ?", studentID, offeredCourseID).Find(&selection).Error
	if err != nil || errors.Is(err, gorm.ErrRecordNotFound) {
		return &model.Selection{}, err
	}
	return &selection, nil
}

// UpdateSelectionScore 更新课程成绩
func UpdateSelectionScore(selectionUpdateReq model.SelectionUpdateReq) error {
	selection, err := GetSelection(int(selectionUpdateReq.StudentID), int(selectionUpdateReq.OfferedCourseID))
	if err != nil || selection.ID == 0 {
		return errors.New("没有此选课记录！")
	}
	usualScore := selectionUpdateReq.UsualScore
	examScore := selectionUpdateReq.ExamScore
	if usualScore < 1 || usualScore > 100 || examScore < 1 || examScore > 100 {
		return errors.New("分数不合法！")
	}

	err = db.Model(&model.Selection{}).
		Where("id = ?", selection.ID).
		Updates(model.Selection{UsualScore: usualScore, ExamScore: examScore}).Error
	return err
}

// DeleteAllSelections 删除所有选课记录
func DeleteAllSelections() {
	db.Where("1 = 1").Delete(&model.Selection{})
}
