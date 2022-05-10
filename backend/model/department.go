package model

type Department struct {
	ID     uint   `gorm:"primarykey"`
	Number string `json:"number" form:"number" gorm:"unique;index"` // 院系号
	Name   string `json:"name" form:"name"  gorm:"unique"`          // 院系名

	Students []Student `json:"students" form:"students"` // 学生
	Teachers []Teacher `json:"teachers" form:"teachers"` // 教师
	Courses  []Course  `json:"courses" form:"courses"`   // 课程
}

type DepartmentRes struct {
	Number string `json:"number"` // 院系号
	Name   string `json:"name"`   // 院系名
}

// CreateDepartmentsExample 创建学院样例
func CreateDepartmentsExample() (departments []Department) {
	departments = []Department{
		{Number: "001", Name: "计算机学院"},
		{Number: "002", Name: "通信学院"},
		{Number: "003", Name: "材料学院"},
		{Number: "004", Name: "自动化学院"},
	}
	db.Model(&Department{}).Create(&departments)
	return
}

// DeleteAllDepartments 删除所有学院
func DeleteAllDepartments() {
	db.Where("1 = 1").Delete(&Department{})
}
