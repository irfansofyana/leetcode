WITH T1 AS (
    SELECT
        contest_id, 
        COUNT(contest_id) AS cnt
    FROM
        Register
    GROUP BY contest_id
), T2 AS (
    SELECT
        T1.contest_id,
        T1.cnt,
        (SELECT COUNT(user_id) FROM Users) AS total
    FROM
        T1
)

SELECT
    contest_id,
    ROUND((CAST (cnt AS float) * 100 / CAST (total AS float)), 2) AS percentage
FROM T2
ORDER BY percentage DESC, contest_id ASC