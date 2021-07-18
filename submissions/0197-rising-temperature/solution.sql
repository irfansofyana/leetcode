# Write your MySQL query statement below
SELECT w.id
FROM Weather AS w, Weather AS w2
WHERE DATEDIFF(w.recordDate, w2.recordDate) = 1 AND w.temperature > w2.temperature

