package service

import (
	"academic-system/model"
	"errors"
	"gorm.io/gorm"
)

// CreateSelectionsExample 创建课程关联实例
func CreateSelectionsExample() (selections []model.Selection) {
	selections = []model.Selection{
		{StudentID: 1, OfferedCourseID: 1, Score: -1},
		{StudentID: 1, OfferedCourseID: 2, Score: -1},
		{StudentID: 1, OfferedCourseID: 4, Score: 75},
		{StudentID: 1, OfferedCourseID: 3, Score: -1},
		{StudentID: 1, OfferedCourseID: 6, Score: -1},

		{StudentID: 2, OfferedCourseID: 1, Score: -1},
		{StudentID: 2, OfferedCourseID: 3, Score: -1},
		{StudentID: 2, OfferedCourseID: 4, Score: 93},
		{StudentID: 2, OfferedCourseID: 5, Score: 70},

		{StudentID: 3, OfferedCourseID: 4, Score: 68},
		{StudentID: 1, OfferedCourseID: 6, Score: -1},

		{StudentID: 5, OfferedCourseID: 4, Score: 85},
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
		return nil, errors.New("选课关系已存在！")
	}

	// 判断学生是否已选课名、学期均相同的课程
	targetOc, err := model.GetOfferedCourseByID(int(selection.OfferedCourseID))
	if err != nil {
		return nil, errors.New("不存在该课程！")
	}
	coursesByStu, err := GetSelectedCourses(int(selection.StudentID))
	if err != nil {
		return nil, err
	}
	for _, course := range *coursesByStu {
		if course.Name == targetOc.Course.Name && course.Term == targetOc.Term {
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
