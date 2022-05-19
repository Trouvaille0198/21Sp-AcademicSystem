package model

// Selection 记录学生选课情况
// 三种情况: 未选 已选 已修
type Selection struct {
	ID              uint `gorm:"primarykey"`
	StudentID       uint
	OfferedCourseID uint

	Student       Student
	OfferedCourse OfferedCourse

	Score      int `gorm:"default:-1"` // 总分, -1表示未评分
	UsualScore int `gorm:"default:-1"` // 平时成绩, -1表示未评分
	ExamScore  int `gorm:"default:-1"` // 考试成绩, -1表示未评分
}

type SelectionCreateReq struct {
	StudentID       uint `form:"student_id"`
	OfferedCourseID uint `form:"offered_course_id"`
}

// SelectionUpdateReq 更新选课记录的请求结构体
type SelectionUpdateReq struct {
	StudentID       uint `form:"student_id"`
	OfferedCourseID uint `form:"offered_course_id"`
	UsualScore      int  `form:"usual_score"`
	ExamScore       int  `form:"exam_score"`
}
