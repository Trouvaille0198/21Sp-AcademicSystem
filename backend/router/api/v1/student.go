package v1

import (
	"academic-system/model"
	"academic-system/service"
	"academic-system/utils"
	"github.com/gin-gonic/gin"
)

// GetStudentByID godoc
// @Summary      根据id获取学生信息
// @Description  根据id获取学生信息
// @Tags         student
// @Param        id   path      int  true  "student ID"
// @Success      200  {object}  model.StudentRes
// @Router       /student/{id} [get]
func GetStudentByID(c *gin.Context) {
	studentID, _ := utils.String2Int(c.Param("id"))

	student, err := service.GetStudentResByID(studentID)
	if err != nil {
		model.FailWithMessage(err.Error(), c)
		return
	}
	model.OkWithData(student, c)
}

// GetAllStudents godoc
// @Summary      获取所有学生信息
// @Description  获取所有学生信息
// @Tags         student
// @Success      200  {object}  []model.StudentRes
// @Router       /student [get]
func GetAllStudents(c *gin.Context) {
	students, err := service.GetAllStudentsRes()
	if err != nil {
		model.FailWithMessage(err.Error(), c)
		return
	}
	model.OkWithData(students, c)
}

// CreateStudent godoc
// @Summary      创建学生
// @Description  创建学生
// @Tags         student
// @Param 		 student   body   model.StudentCreateReq    true   "student 实例"
// @Success      200  {string} string
// @Router       /student [post]
func CreateStudent(c *gin.Context) {
	student := model.Student{}
	if err := c.ShouldBindJSON(&student); err != nil {
		model.FailWithMessage(err.Error(), c)
		return
	}
	newStudent, err := service.CreateStudent(student)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: students.number" {
			model.FailWithMessage("学号已存在", c)
			return
		}
		model.FailWithMessage(err.Error(), c)
		return
	}
	model.OkWithData(newStudent, c)
}

// DeleteStudentByID godoc
// @Summary      根据id删除学生
// @Description  根据id删除学生
// @Tags         student
// @Param        id   path      int  true  "student ID"
// @Success      200  {string} string
// @Router       /student/{id} [delete]
func DeleteStudentByID(c *gin.Context) {
	studentID, _ := utils.String2Int(c.Param("id"))
	err := service.DeleteStudent(studentID)
	if err != nil {
		model.FailWithMessage(err.Error(), c)
		return
	}
	model.Ok(c)
}
