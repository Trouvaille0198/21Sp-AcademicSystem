package v1

import (
	"academic-system/model"
	"academic-system/service"
	"academic-system/utils"
	"github.com/gin-gonic/gin"
)

// GetCourseByID godoc
// @Summary      根据id获取课程信息
// @Description  根据id获取课程信息
// @Tags         course
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "course ID"
// @Success      200  {object}  model.CourseRes
// @Router       /course/{id} [get]
func GetCourseByID(c *gin.Context) {
	courseID, _ := utils.String2Int(c.Param("id"))
	course, err := service.GetCourseByID(courseID)
	if err != nil {
		model.FailWithMessage(err.Error(), c)
		return
	}
	model.OkWithData(course, c)
}

// CreateCourse godoc
// @Summary      创建课程
// @Description  创建课程
// @Tags         course
// @Param 		 course   body   model.CourseCreateReq  true   "course 实例"
// @Success      200  {string} string
// @Router       /course [post]
func CreateCourse(c *gin.Context) {
	course := model.Course{}
	if err := c.ShouldBindJSON(&course); err != nil {
		model.FailWithMessage(err.Error(), c)
		return
	}
	newCourse, err := service.CreateCourse(course)
	if err != nil {
		model.FailWithMessage(err.Error(), c)
		return
	}
	model.OkWithData(newCourse, c)
}

// GetCourses godoc
// @Summary      获取所有课程
// @Description  获取所有课程
// @Tags         course
// @Success      200  {object}  []model.CourseRes
// @Router       /course [get]
func GetCourses(c *gin.Context) {
	coursesResult, err := service.GetCourses()
	if err != nil {
		model.FailWithMessage(err.Error(), c)
		return
	}
	model.OkWithData(coursesResult, c)
}

// DeleteCourseByID godoc
// @Summary      根据id删除课程
// @Description  根据id删除课程
// @Tags         course
// @Param        id   path      int  true  "course ID"
// @Success      200  {string} string
// @Router       /course/{id} [delete]
func DeleteCourseByID(c *gin.Context) {
	courseID, _ := utils.String2Int(c.Param("id"))
	err := service.DeleteCourse(courseID)
	if err != nil {
		model.FailWithMessage(err.Error(), c)
		return
	}
	model.Ok(c)
}
