SELECT
    DISTINCT t1.product_id, t1.product_name
FROM (
    SELECT s.product_id, p.product_name
    FROM Sales AS s 
    LEFT JOIN product AS p
    ON s.product_id = p.product_id
    WHERE s.sale_date >= '2019-01-01' AND s.sale_date <= '2019-03-31'
) AS t1
WHERE t1.product_id NOT IN (
    SELECT t2.product_id
    FROM (
        (
            SELECT DISTINCT
                s.product_id,
                CASE 
                    WHEN (s.sale_date >= '2019-01-01' AND s.sale_date <= '2019-03-31') THEN 1
                    ELSE 0
                END AS tmp
            FROM Sales AS s
            LEFT JOIN product AS p
            ON s.product_id = p.product_id
        )
        INTERSECT
        (
        SELECT p.product_id, o.tmp
            FROM Product AS p
            CROSS JOIN (
                SELECT 0 AS tmp
            ) AS o
        )
    ) AS t2
)
