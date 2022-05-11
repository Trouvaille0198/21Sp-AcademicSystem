package model

import (
	"gorm.io/gorm"
)

type Teacher struct {
	ID       uint   `gorm:"primarykey"`
	Number   string `gorm:"unique;index"` // 教师工号
	Name     string
	Sex      string
	Age      uint8
	Password string `gorm:"default:123"`

	DepartmentID uint // 所属院系
}

// TeacherCreateReq 创建老师的请求格式
type TeacherCreateReq struct {
	Number       string `form:"number"`
	Name         string `form:"name"`
	Sex          string `form:"sex"`
	Age          uint   `form:"age"`
	Password     string `form:"password"`
	DepartmentID uint   `form:"department_id"`
}

type TeacherRes struct {
	ID             uint   `json:"id"`
	Number         string `json:"number"`
	Name           string `json:"name"`
	Sex            string `json:"sex"`
	Age            uint8  `json:"age"`
	Password       string `json:"password"`
	DepartmentName string `json:"department_name"`
}

// BeforeDelete 删除老师时级联删除开课记录
func (t *Teacher) BeforeDelete(tx *gorm.DB) (err error) {
	_ = tx.Where("teacher_id = ?", t.ID).Delete(&OfferedCourse{}).RowsAffected
	return
}
