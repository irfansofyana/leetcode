WITH T1 AS (
    SELECT
        query_name,
        AVG(1.0 * rating / position) AS quality
    FROM
        Queries
    WHERE 
        query_name IS NOT NULL
    GROUP BY query_name
),
T2 AS (
    SELECT
        query_name,
        COUNT(rating) AS poor_cnt
    FROM 
        Queries
    WHERE rating < 3
    GROUP BY query_name
),
T3 AS (
    SELECT
        query_name,
        COUNT(rating) as total_cnt
    FROM Queries
    GROUP BY query_name
)

-- SELECT * FROM T1
-- SELECT * FROM T2
-- SELECT * FROM T3

SELECT
    T1.query_name,
    COALESCE(ROUND(T1.quality, 2), 0.00) AS quality,
    COALESCE(ROUND(100.00 * T2.poor_cnt  / T3.total_cnt, 2), 0.00) AS poor_query_percentage
FROM T1
LEFT JOIN T2
ON T1.query_name = T2.query_name
LEFT JOIN T3
ON T2.query_name = T3.query_name
