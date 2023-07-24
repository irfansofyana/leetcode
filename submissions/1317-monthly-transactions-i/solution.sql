WITH T1 AS (
    SELECT
        t.*,
        FORMAT(t.trans_date, 'yyyy-MM') AS month   
    FROM Transactions AS t
), T2 AS (
    SELECT 
        month,
        country,
        COUNT(id) AS trans_count
    FROM T1
    GROUP BY month, country
), T3 AS (
    SELECT 
        month,
        country,
        COUNT(id) AS approved_count
    FROM T1
    WHERE state = 'approved'
    GROUP BY month, country
), T4 AS (
    SELECT 
        month,
        country,
        SUM(amount) AS trans_total_amount
    FROM T1
    GROUP BY month, country
), T5 AS (
    SELECT 
        month,
        country,
        SUM(amount) AS approved_total_amount
    FROM T1
    WHERE state = 'approved'
    GROUP BY month, country
)

select
    T2.month,
    T2.country,
    COALESCE(T2.trans_count, 0) AS trans_count,
    COALESCE(T3.approved_count, 0) AS approved_count,
    COALESCE(T4.trans_total_amount, 0) AS trans_total_amount,
    COALESCE(T5.approved_total_amount, 0) AS approved_total_amount
FROM
    T2
LEFT JOIN T3
ON T2.month = T3.month AND T2.country = T3.country
LEFT JOIN T4
ON T2.month = T4.month and T2.country = T4.country
LEFT JOIN T5
ON T2.month = T5.month AND T2.country = T5.country
