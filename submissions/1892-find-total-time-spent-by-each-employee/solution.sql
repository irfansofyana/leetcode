SELECT
    day,
    emp_id,
    SUM(total_time) AS total_time
FROM (
    SELECT emp_id, event_day AS day, (out_time-in_time) AS total_time
    FROM employees
) AS tmp
GROUP BY day, emp_id
ORDER BY day
