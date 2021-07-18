# Write your MySQL query statement below
SELECT t.class
FROM (
    SELECT COUNT(DISTINCT student) as cnt, class
    FROM courses
    GROUP BY class
) AS t
WHERE t.cnt >= 5
