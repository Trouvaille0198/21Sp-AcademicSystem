package v1

import (
	"academic-system/global"
	"academic-system/model"
	"academic-system/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// Ping godoc
// @Summary      ping!
// @Description  pong me!
// @Tags         test
// @Success      200  {string} string
// @Router       /test/ping [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// GenerateExamples godoc
// @Summary      生成样例数据
// @Description  生成样例数据
// @Tags         test
// @Success      200  {string} string
// @Router       /test/generate-examples [get]
func GenerateExamples(c *gin.Context) {
	// 创建顺序有讲究 不能乱
	model.CreateDepartmentsExample()
	service.CreateStudentsExample()
	service.CreateTeachersExample()
	service.CreateCoursesExample()
	model.CreateOfferedCoursesExample()
	service.CreateSelectionsExample()

	model.OkWithMessage("生成样例数据成功", c)
}

// DeleteAll godoc
// @Summary      删除全部数据
// @Description  删除全部数据
// @Tags         test
// @Success      200  {string} string
// @Router       /test/delete-all [delete]
func DeleteAll(c *gin.Context) {
	// model.DeleteAllDepartments()
	// service.DeleteAllStudents()
	// service.DeleteAllTeachers()
	// service.DeleteAllCourses()
	// model.DeleteAllOfferedCourse()
	// model.DeleteAllSelections()
	err := global.DB.Migrator().DropTable(&model.Department{}, &model.Course{}, &model.Teacher{}, &model.OfferedCourse{},
		&model.Student{}, &model.Selection{})
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = global.DB.AutoMigrate(
		&model.Department{}, &model.Course{}, &model.Teacher{}, &model.OfferedCourse{},
		&model.Student{}, &model.Selection{})
	if err != nil {
		log.Fatalf(err.Error())
	}

	model.OkWithMessage("数据库清空完成", c)
}
