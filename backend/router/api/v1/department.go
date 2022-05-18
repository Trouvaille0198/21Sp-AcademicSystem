package v1

import (
	"academic-system/model"
	"academic-system/service"
	"academic-system/utils"
	"github.com/gin-gonic/gin"
)

// GetDepartmentByID godoc
// @Summary      获根据id获取学院信息
// @Description  获根据id获取学院信息
// @Tags         department
// @Param        id   path      int  true  "department ID"
// @Success      200  {object}  model.DepartmentRes
// @Router       /department/{id} [get]
func GetDepartmentByID(c *gin.Context) {
	departmentID, _ := utils.String2Int(c.Param("id"))

	department, err := service.GetDepartmentByID(departmentID)
	if err != nil {
		model.FailWithMessage(err.Error(), c)
		return
	}
	model.OkWithData(department, c)
}

// GetDepartments godoc
// @Summary      获取所有学院信息
// @Description  获取所有学院信息
// @Tags         department
// @Success      200  {object}  []model.DepartmentRes
// @Router       /department [get]
func GetDepartments(c *gin.Context) {

	departments, err := service.GetDepartments()
	if err != nil {
		model.FailWithMessage(err.Error(), c)
		return
	}
	model.OkWithData(departments, c)
}
