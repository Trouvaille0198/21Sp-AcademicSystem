package model

import (
	"gorm.io/gorm"
)

// Student 数据库学生表
type Student struct {
	ID       uint   `gorm:"primarykey"`
	Number   string `gorm:"unique;index"` // 学号
	Name     string // 姓名
	Sex      string // 性别
	Age      uint   // 年龄
	Password string `gorm:"default:123" example:"123"` // 密码

	DepartmentID uint // 所属院系号

	Selections []Selection // 选课情况
}

// StudentCreateReq 创建学生的请求格式
type StudentCreateReq struct {
	Number       string `form:"number"`
	Name         string `form:"name"`
	Sex          string `form:"sex"`
	Age          uint   `form:"age"`
	Password     string `form:"password"`
	DepartmentID uint   `form:"department_id"`
}

// StudentRes 学生信息
type StudentRes struct {
	ID             uint   `json:"id"`
	Number         string `json:"number"`
	Name           string `json:"name"`
	Sex            string `json:"sex"`
	Age            uint   `json:"age"`
	DepartmentName string `json:"department_name"`
}

type StudentWithScoreRes struct {
}

// BeforeDelete 删除学生时级联删除选课关系
func (stu *Student) BeforeDelete(tx *gorm.DB) (err error) {
	_ = tx.Where("student_id = ?", stu.ID).Unscoped().Delete(&Selection{}).RowsAffected
	return
}
