SELECT
    t.user_id,
    MAX(t.time_stamp) AS last_stamp
FROM (
    SELECT *
    FROM 
        Logins
    WHERE
        DATE(time_stamp) >= '2020-01-01' AND DATE(time_stamp) <= '2020-12-31'
) AS t
GROUP BY t.user_id;

