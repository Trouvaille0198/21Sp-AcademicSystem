package model

// OfferedCourse 开课信息
type OfferedCourse struct {
	ID   uint   `gorm:"primarykey"`
	Term string // 学期

	Course    Course
	CourseID  uint // 课号
	TeacherID uint // 教师工号

	Selections []Selection `json:"selections"`
}

// OfferedCourseRes 老师视角下的开课信息
type OfferedCourseRes struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Number      string `json:"number"`
	Credit      uint   `json:"credit"`
	TeacherName string `json:"teacher_name"`
	Department  string `json:"department"`
	Term        string `json:"term"`
}

// SelectedCourseRes 学生视角下的选课记录
type SelectedCourseRes struct {
	OfferedCourseID uint   `json:"offered_course_id"`
	SelectionID     uint   `json:"selection_id"`
	Name            string `json:"name"`
	Number          string `json:"number"`
	Credit          uint   `json:"credit"`
	TeacherName     string `json:"teacher_name"`
	StudentName     string `json:"student_name"`
	Department      string `json:"department"`
	Term            string `json:"term"`
	UsualScore      int    `json:"usual_score"`
	ExamScore       int    `json:"exam_score"`
	Score           int    `json:"score"`
}
