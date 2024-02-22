SELECT DISTINCT
    p.product_id,
    (CASE
        WHEN
            t3.min_diff IS NULL THEN 10
        ELSE
            t3.new_price
    END) AS price
FROM Products AS p
LEFT JOIN (
    SELECT 
        t1.product_id,
        t1.min_diff,
        t2.new_price
    FROM (
        SELECT
            t.product_id,
            MIN(t.diff) AS min_diff
        FROM 
            (SELECT
                Products.*,
                DATEDIFF('2019-08-16', change_date) AS diff
            FROM 
                Products
            WHERE DATEDIFF('2019-08-16', change_date) >= 0) AS t
        GROUP BY t.product_id
    ) AS t1
    LEFT JOIN (
        SELECT
            Products.*,
            DATEDIFF('2019-08-16', change_date) AS diff
        FROM 
            Products
        WHERE DATEDIFF('2019-08-16', change_date) >= 0
    ) AS t2
    ON t1.product_id = t2.product_id AND t1.min_diff = t2.diff
) AS t3
ON p.product_id = t3.product_id

