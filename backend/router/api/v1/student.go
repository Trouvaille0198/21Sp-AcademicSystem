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

// GetStudents godoc
// @Summary      获取所有学生信息
// @Description  获取所有学生信息
// @Tags         student
// @Param        number query string false "学号, 可选"
// @Success      200  {object}  []model.StudentRes
// @Router       /student [get]
func GetStudents(c *gin.Context) {
	students, err := service.GetAllStudentsRes()
	if err != nil {
		model.FailWithMessage(err.Error(), c)
		return
	}

	number, ok := c.GetQuery("number")
	if ok {
		for _, stu := range *students {
			if stu.Number == number {
				model.OkWithData(&[]model.StudentRes{stu}, c)
				return
			}
		}
		model.FailWithDetailed(&[]model.StudentRes{}, "没有找到学生", c)
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
		model.FailWithMessage("学号已存在", c)
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

// GetStudentsByOc godoc
// @Summary      获取指定id的开课记录下所有学生信息及成绩
// @Description  获取指定id的开课记录下所有学生信息及成绩
// @Tags         student, offered_course
// @Param        id   path      int  true  "offered_course ID"
// @Success      200  {object}  []model.StudentWithScoreRes
// @Router       /offered_course/{id}/student [get]
func GetStudentsByOc(c *gin.Context) {
	ocID, _ := utils.String2Int(c.Param("id"))

	students, err := service.GetStudentsByOfferedCourse(ocID)
	if err != nil {
		model.FailWithMessage(err.Error(), c)
		return
	}
	model.OkWithData(students, c)
}
