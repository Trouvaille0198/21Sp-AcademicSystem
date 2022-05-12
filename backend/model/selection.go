package model

// Selection 记录学生选课情况
// 三种情况: 未选 已选 已修
type Selection struct {
	ID              uint `gorm:"primarykey"`
	StudentID       uint `gorm:"student_id"`
	OfferedCourseID uint `gorm:"offered_course_id"`

	Student       Student
	OfferedCourse OfferedCourse
	Score         int `gorm:"default:-1"` // 分数, -1表示未评分
}

// SelectionUpdateReq 更新选课记录的请求结构体
type SelectionUpdateReq struct {
	StudentID uint `json:"student_id"`
	CourseID  uint `json:"course_id"`
}
