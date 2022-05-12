package v1

import (
	"academic-system/model"
	"academic-system/service"
	"github.com/gin-gonic/gin"
)

// UpdateSelectionScore godoc
// @Summary      更新课程成绩
// @Description  更新课程成绩
// @Tags         selection, admin
// @Param 		 selection   body   model.SelectionUpdateReq    true   "选课信息"
// @Success      200  {string} string
// @Router       /selection [put]
func UpdateSelectionScore(c *gin.Context) {
	var selectionUpdateReq model.SelectionUpdateReq
	if err := c.ShouldBindJSON(&selectionUpdateReq); err != nil {
		model.FailWithMessage(err.Error(), c)
		return
	}

	err := service.UpdateSelectionScore(selectionUpdateReq)
	if err != nil {
		model.FailWithMessage(err.Error(), c)
		return
	}
	model.Ok(c)
}

// CreateSelection godoc
// @Summary      创建选课
// @Description  创建选课
// @Tags         selection
// @Param 		 selection   body   model.SelectionCreateReq   true   "选课情况"
// @Success      200  {string} string
// @Router       /selection [post]
func CreateSelection(c *gin.Context) {
	var selection model.Selection
	if err := c.ShouldBindJSON(&selection); err != nil {
		model.FailWithMessage(err.Error(), c)
		return
	}

	newSelection, err := service.CreateSelection(selection)
	if err != nil {
		model.FailWithMessage(err.Error(), c)
		return
	}
	model.OkWithData(newSelection, c)
}

// DeleteSelection godoc
// @Summary      删除选课
// @Description  删除选课
// @Tags         selection
// @Param 		 selection   body   model.SelectionCreateReq   true   "选课情况"
// @Success      200  {string} string
// @Router       /selection [delete]
func DeleteSelection(c *gin.Context) {
	var selection model.Selection
	if err := c.ShouldBindJSON(&selection); err != nil {
		model.FailWithMessage(err.Error(), c)
		return
	}

	err := service.DeleteSelection(selection)
	if err != nil {
		model.FailWithMessage(err.Error(), c)
		return
	}
	model.Ok(c)
}
