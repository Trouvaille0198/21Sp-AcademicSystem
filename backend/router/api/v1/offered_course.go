package v1

import (
	"academic-system/model"
	"academic-system/service"
	"academic-system/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

// GetOCsByStu godoc
// @Summary      获取指定学生的所有课程
// @Description  get offered courses by student id 获取指定学生的所有课程
// @Tags         student, course
// @Accept       json
// @Produce      json
// @Param 		 id   path   int   true   "student ID"
// @Param        hasScore query bool false "是否有成绩 不写即全部返回"
// @Success      200  {object}  []model.OfferedCourseRes
// @Router       /student/{id}/offered_course [get]
func GetOCsByStu(c *gin.Context) {
	hasScore, ok := c.GetQuery("hasScore")
	studentID, _ := utils.String2Int(c.Param("id"))
	var courses *[]model.OfferedCourseRes
	var sqlErr error
	if !ok {
		courses, sqlErr = service.GetOCsByStu(studentID)
	} else {
		hasScore, err := strconv.ParseBool(hasScore)
		if err != nil {
			model.FailWithMessage(err.Error(), c)
			return
		}
		if hasScore {
			courses, sqlErr = service.GetOCsByStuWithScore(studentID)
		} else {
			courses, sqlErr = service.GetOCsByStuWithoutScore(studentID)
		}
	}

	if sqlErr != nil {
		model.FailWithMessage(sqlErr.Error(), c)
		return
	}
	model.OkWithData(courses, c)
}
