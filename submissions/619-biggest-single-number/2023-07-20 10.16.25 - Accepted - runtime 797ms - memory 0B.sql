SELECT
    IFNULL(( SELECT
        num
    FROM (
        SELECT num, COUNT(*) AS cnt
        FROM MyNumbers 
        GROUP BY num
    ) AS t
    WHERE t.cnt = 1
    ORDER BY num DESC
    LIMIT 1), NULL) AS num