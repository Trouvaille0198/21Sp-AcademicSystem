-- 当平时成绩和考试成绩更新时，更新学生的总成绩
CREATE TRIGGER TRI_UPD_SCORE
    AFTER UPDATE ON selection
    REFERENCING
        OLD AS OLDTUPLE
        NEW AS NEWTUPLE
    FOR EACH ROW
    WHEN ((OLDTUPLE.usual_score <> NEWTUPLE.usual_score OR OLDTUPLE.exam_score <> NEWTUPLE.exam_score)
    AND NEWTUPLE.usual_score <> -1 AND NEWTUPLE.exam_score <> -1)
    BEGIN
        UPDATE selection SET score = (NEWTUPLE.usual_score + NEWTUPLE.exam_score)/2 WHERE id = NEWTUPLE.id;
    END;
