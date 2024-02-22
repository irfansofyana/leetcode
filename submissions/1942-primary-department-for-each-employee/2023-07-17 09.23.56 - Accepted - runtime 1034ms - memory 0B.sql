WITH count_table AS (
    SELECT employee_id, COUNT(employee_id) AS cnt
    FROM Employee
    GROUP BY employee_id
), merge_table AS (
    SELECT E.employee_id, department_id, C.cnt AS cnt, primary_flag
    FROM Employee AS E LEFT JOIN count_table AS C
    ON E.employee_id = C.employee_id
)

SELECT
    employee_id, department_id
FROM merge_table AS m
WHERE m.cnt = 1
UNION
SELECT
    employee_id, department_id
FROM merge_table AS m
WHERE m.cnt > 1 AND m.primary_flag = 'Y'