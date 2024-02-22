SELECT
  tmp.student_id AS student_id,
  tmp.student_name AS student_name,
  tmp.subject_name AS subject_name,
  (CASE 
    WHEN e.student_id IS NULL THEN 0
    ELSE COUNT(*)
  END) AS attended_exams
FROM
  (
    SELECT student_id, student_name, subject_name
    FROM Students AS s CROSS JOIN Subjects AS su
  ) AS tmp LEFT JOIN Examinations AS e
  ON tmp.student_id = e.student_id AND tmp.subject_name = e.subject_name
GROUP BY tmp.student_id, tmp.student_name, tmp.subject_name
ORDER BY tmp.student_id, tmp.subject_name