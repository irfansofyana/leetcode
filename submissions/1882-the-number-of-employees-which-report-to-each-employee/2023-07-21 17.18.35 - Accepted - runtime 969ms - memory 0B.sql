SELECT DISTINCT
    reports_to AS employee_id,
    t.name AS name,
    COUNT(reports_to) AS reports_count,
    ROUND(AVG(age)) AS average_age
FROM
    Employees AS e
LEFT JOIN (
    SELECT employee_id, name
    FROM Employees
) AS t
ON
    e.reports_to = t.employee_id
WHERE 
    reports_to IS NOT NULL
GROUP BY
    reports_to
ORDER BY
    employee_id
