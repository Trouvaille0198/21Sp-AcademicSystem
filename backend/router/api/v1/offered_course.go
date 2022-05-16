package v1

import (
	"academic-system/model"
	"academic-system/service"
	"academic-system/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

// GetOfferedCourses godoc
// @Summary      获取所有开课记录
// @Description  获取所有开课记录
// @Tags         offered_course
// @Accept       json
// @Produce      json
// @Success      200  {object}  []model.OfferedCourseRes
// @Router       /offered_course [get]
func GetOfferedCourses(c *gin.Context) {
	courses, err := service.GetOfferedCourses()
	if err != nil {
		model.FailWithMessage(err.Error(), c)
		return
	}
	model.OkWithData(courses, c)
}

// GetOfferedCourseByID godoc
// @Summary      获取指定id的开课课程
// @Description  获取指定id的开课课程
// @Tags         offered_course
// @Accept       json
// @Produce      json
// @Param 		 id   path   int   true   "offered_course ID"
// @Success      200  {object}  model.OfferedCourseRes
// @Router       /offered_course/{id} [get]
func GetOfferedCourseByID(c *gin.Context) {
	ocID, _ := utils.String2Int(c.Param("id"))
	courses, err := service.GetOfferedCourseByID(ocID)
	if err != nil {
		model.FailWithMessage(err.Error(), c)
		return
	}
	model.OkWithData(courses, c)
}

// GetSelectedCourses godoc
// @Summary      获取指定学生的所有课程
// @Description  get offered courses by student id 获取指定学生的所有课程
// @Tags         student, offered_course
// @Accept       json
// @Produce      json
// @Param 		 id   path   int   true   "student ID"
// @Param        hasScore query bool false "是否有成绩 不写即全部返回"
// @Param        term query string false "指定学期"
// @Success      200  {object}  []model.SelectedCourseRes
// @Router       /student/{id}/selected_course [get]
func GetSelectedCourses(c *gin.Context) {
	hasScore, ok := c.GetQuery("hasScore")
	studentID, _ := utils.String2Int(c.Param("id"))
	var courses *[]model.SelectedCourseRes
	var sqlErr error
	if !ok {
		courses, sqlErr = service.GetSelectedCoursesByStu(studentID)
	} else {
		hasScore, err := strconv.ParseBool(hasScore)
		if err != nil {
			model.FailWithMessage(err.Error(), c)
			return
		}
		if hasScore {
			courses, sqlErr = service.GetSelectedCoursesWithScore(studentID)
		} else {
			courses, sqlErr = service.GetSelectedCoursesWithoutScore(studentID)
		}
	}

	if sqlErr != nil {
		model.FailWithMessage(sqlErr.Error(), c)
		return
	}
	model.OkWithData(courses, c)
}

// GetOCsByTeacher godoc
// @Summary      获取指定教师的所有开课课程
// @Description  get offered courses by teacher id 获取指定教师的所有开课课程
// @Tags         teacher, offered_course
// @Accept       json
// @Produce      json
// @Param 		 id   path   int   true   "teacher ID"
// @Success      200  {object}  []model.OfferedCourseRes
// @Router       /teacher/{id}/offered_course [get]
func GetOCsByTeacher(c *gin.Context) {
	teacherID, _ := utils.String2Int(c.Param("id"))
	courses, err := service.GetOCsByTeacher(teacherID)
	if err != nil {
		model.FailWithMessage(err.Error(), c)
		return
	}
	model.OkWithData(courses, c)
}
