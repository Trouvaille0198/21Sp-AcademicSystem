package v1

import (
	"academic-system/model"
	"academic-system/service"
	"academic-system/utils"
	"github.com/gin-gonic/gin"
)

// GetTeacherByID godoc
// @Summary      根据id获取老师信息
// @Description  根据id获取老师信息
// @Tags         teacher
// @Param        id   path      int  true  "teacher ID"
// @Success      200  {object}  model.TeacherRes
// @Router       /teacher/{id} [get]
func GetTeacherByID(c *gin.Context) {
	teacherID, _ := utils.String2Int(c.Param("id"))

	teacher, err := service.GetTeacherResByID(teacherID)
	if err != nil {
		model.FailWithMessage(err.Error(), c)
		return
	}
	model.OkWithData(teacher, c)
}

// GetAllTeachers godoc
// @Summary      获取所有老师信息
// @Description  获取所有老师信息
// @Tags         teacher
// @Success      200  {object}  []model.TeacherRes
// @Router       /teacher [get]
func GetAllTeachers(c *gin.Context) {
	teachers, err := service.GetAllTeachersRes()
	if err != nil {
		model.FailWithMessage(err.Error(), c)
		return
	}
	model.OkWithData(teachers, c)
}

// CreateTeacher godoc
// @Summary      创建老师
// @Description  创建老师
// @Tags         teacher
// @Param 		 teacher   body   model.TeacherCreateReq    true   "teacher 实例"
// @Success      200  {string} string
// @Router       /teacher [post]
func CreateTeacher(c *gin.Context) {
	teacher := model.Teacher{}
	if err := c.ShouldBindJSON(&teacher); err != nil {
		model.FailWithMessage(err.Error(), c)
		return
	}
	newTeacher, err := service.CreateTeacher(teacher)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: teachers.number" {
			model.FailWithMessage("工号已存在", c)
			return
		}
		model.FailWithMessage(err.Error(), c)
		return
	}
	model.OkWithData(newTeacher, c)
}

// DeleteTeacherByID godoc
// @Summary      根据id删除老师
// @Description  根据id删除老师
// @Tags         teacher
// @Param        id   path      int  true  "teacher ID"
// @Success      200  {string} string
// @Router       /teacher/{id} [delete]
func DeleteTeacherByID(c *gin.Context) {
	teacherID, _ := utils.String2Int(c.Param("id"))
	err := service.DeleteTeacher(teacherID)
	if err != nil {
		model.FailWithMessage(err.Error(), c)
		return
	}
	model.Ok(c)
}
