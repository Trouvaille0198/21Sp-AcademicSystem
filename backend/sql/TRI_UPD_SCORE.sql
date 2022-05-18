-- 当平时成绩和考试成绩更新时，更新学生的总成绩
DROP TRIGGER IF EXISTS BEF_UPD_SELECTION;
DELIMITER ;;
CREATE TRIGGER BEF_UPD_SELECTION
    BEFORE UPDATE
    ON selection
    FOR EACH ROW
BEGIN
    IF ((OLD.usual_score <> NEW.usual_score OR OLD.exam_score <> NEW.exam_score)
        AND NEW.usual_score <> -1 AND NEW.exam_score <> -1)
    THEN
        SET NEW.score = (NEW.usual_score + NEW.exam_score) / 2;
    END IF;
END
;;
DELIMITER ;

-- 当删除学院时，清除与该学院相关的所有记录
DROP TRIGGER IF EXISTS AFT_DEL_DEPARTMENT;
DELIMITER ;;
CREATE TRIGGER AFT_DEL_DEPARTMENT
    AFTER DELETE
    ON department
    FOR EACH ROW
BEGIN
    -- 删除涉及该学院的选课记录
    DELETE
    FROM selection
    WHERE student_id IN (SELECT id FROM student WHERE department_id = old.id)
       OR offered_course_id IN (SELECT id
                                FROM offered_course
                                WHERE teacher_id IN (SELECT id FROM teacher WHERE department_id = old.id)
                                   OR course_id IN (SELECT id FROM course WHERE department_id = old.id));
    -- 删除涉及该学院的开课记录
    DELETE
    FROM offered_course
    WHERE teacher_id IN (SELECT id FROM teacher WHERE department_id = old.id)
       OR course_id IN (SELECT id FROM course WHERE department_id = old.id);
    -- 删除涉及该学院的学生、老师和学生
    DELETE FROM course WHERE department_id = old.id;
    DELETE FROM teacher WHERE department_id = old.id;
    DELETE FROM student WHERE department_id = old.id;
END
;;
DELIMITER ;


-- 获取指定学院的学生、老师、课程数量
DROP PROCEDURE IF EXISTS DEP_INFO;
DELIMITER ;;
CREATE PROCEDURE DEP_INFO(in dep_id integer,
                          out student_num integer, out teacher_num integer, out course_num integer)
BEGIN
    SELECT count(*) INTO student_num FROM student WHERE department_id = dep_id;
    SELECT count(*) INTO teacher_num FROM teacher WHERE department_id = dep_id;
    SELECT count(*) INTO course_num FROM course WHERE department_id = dep_id;
END
;;
DELIMITER ;