package v1

import (
	"academic-system/model"
	"academic-system/service"
	"github.com/gin-gonic/gin"
)

// LoginAsStu godoc
// @Summary      以学生身份登录
// @Description  以学生身份登录
// @Tags         login
// @Accept       mpfd
// @Produce      json
// @Param        number   formData   string  true  "学号"
// @Param        password   formData   string  true  "密码"
// @Success      200  {object}  []model.Student
// @Router       /login/student [post]
func LoginAsStu(c *gin.Context) {
	number := c.PostForm("number")
	password := c.PostForm("password")

	student, err := service.GetStudentByNumber(number)
	if err != nil || student.ID == 0 {
		model.FailWithMessage("找不到此账号", c)
		return
	}
	// 简单地验证一下密码
	if student.Password != password {
		model.FailWithMessage("密码错误", c)
		return
	}
	model.OkWithDetailed(student, "登陆成功", c)

}

// LoginAsAdmin godoc
// @Summary      以管理员身份登录
// @Description  以管理员身份登录
// @Tags         login
// @Param        number   formData   string  true  "账号"
// @Param        password   formData   string  true  "密码"
// @Accept       mpfd
// @Produce      json
// @Success      200
// @Router       /login/admin [post]
func LoginAsAdmin(c *gin.Context) {
	number := c.PostForm("number")
	password := c.PostForm("password")

	// 简单地验证一下账号和密码
	if number != "123" || password != "123" {
		model.FailWithMessage("账号和密码错误", c)
		return
	}

	model.OkWithMessage("登陆成功", c)

}

// LoginAsTeacher godoc
// @Summary      以教师身份登录
// @Description  以教师身份登录
// @Tags         login
// @Param        number   formData   string  true  "账号"
// @Param        password   formData   string  true  "密码"
// @Accept       mpfd
// @Produce      json
// @Success      200  {object}  []model.Teacher
// @Router       /login/teacher [post]
func LoginAsTeacher(c *gin.Context) {
	number := c.PostForm("number")
	password := c.PostForm("password")

	teacher, err := service.GetTeacherByNumber(number)
	if err != nil || teacher.ID == 0 {
		model.FailWithMessage("找不到此账号", c)
		return
	}
	// 简单地验证一下密码
	if teacher.Password != password {
		model.FailWithMessage("密码错误", c)
		return
	}
	model.OkWithDetailed(teacher, "登陆成功", c)
}
