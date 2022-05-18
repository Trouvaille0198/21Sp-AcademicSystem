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
        -- UPDATE selection
        SET NEW.score = (NEW.usual_score + NEW.exam_score) / 2;
        -- WHERE id = NEW.id;
    END IF;
END
;;
DELIMITER ;

-- ----------------------------
-- Triggers structure for table department
-- ----------------------------
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