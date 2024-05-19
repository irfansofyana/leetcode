WITH T1 AS (
    SELECT
        u.product_id,
        ROUND(1.0 * SUM(u.units * p.price) / SUM(u.units), 2) AS average_price
    FROM
        UnitsSold AS u
    LEFT JOIN 
        Prices AS p
    ON
        u.product_id = p.product_id
    AND
        u.purchase_date >= p.start_date
    AND
        u.purchase_date <= p.end_date
    GROUP BY
        u.product_id
)

SELECT 
    DISTINCT(p.product_id) AS product_id,
    COALESCE(T1.average_price, 0) AS average_price
FROM 
    Prices AS p
LEFT JOIN
    T1
ON p.product_id = T1.product_id

