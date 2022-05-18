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
	service.CreateOfferedCoursesExample()
	service.CreateSelectionsExample()
	global.DB.Exec("-- 当平时成绩和考试成绩更新时，更新学生的总成绩\nDROP TRIGGER IF EXISTS BEF_UPD_SELECTION;\nDELIMITER ;;\nCREATE TRIGGER BEF_UPD_SELECTION\n    BEFORE UPDATE\n    ON selection\n    FOR EACH ROW\nBEGIN\n    IF ((OLD.usual_score <> NEW.usual_score OR OLD.exam_score <> NEW.exam_score)\n        AND NEW.usual_score <> -1 AND NEW.exam_score <> -1)\n    THEN\n        -- UPDATE selection\n        SET NEW.score = (NEW.usual_score + NEW.exam_score) / 2;\n        -- WHERE id = NEW.id;\n    END IF;\nEND\n;;\nDELIMITER ;\n\n-- ----------------------------\n-- Triggers structure for table department\n-- ----------------------------\nDROP TRIGGER IF EXISTS AFT_DEL_DEPARTMENT;\nDELIMITER ;;\nCREATE TRIGGER AFT_DEL_DEPARTMENT\n    AFTER DELETE\n    ON department\n    FOR EACH ROW\nBEGIN\n    -- 删除涉及该学院的选课记录\n    DELETE\n    FROM selection\n    WHERE student_id IN (SELECT id FROM student WHERE department_id = old.id)\n       OR offered_course_id IN (SELECT id\n                                FROM offered_course\n                                WHERE teacher_id IN (SELECT id FROM teacher WHERE department_id = old.id)\n                                   OR course_id IN (SELECT id FROM course WHERE department_id = old.id));\n    -- 删除涉及该学院的开课记录\n    DELETE\n    FROM offered_course\n    WHERE teacher_id IN (SELECT id FROM teacher WHERE department_id = old.id)\n       OR course_id IN (SELECT id FROM course WHERE department_id = old.id);\n    -- 删除涉及该学院的学生、老师和学生\n    DELETE FROM course WHERE department_id = old.id;\n    DELETE FROM teacher WHERE department_id = old.id;\n    DELETE FROM student WHERE department_id = old.id;\nEND\n;;\nDELIMITER ;")

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
