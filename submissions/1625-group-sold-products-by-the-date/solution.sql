WITH T AS (
    SELECT
        sell_date,
        COUNT(DISTINCT(product)) AS num_sold
    FROM Activities
    GROUP BY sell_date
    ORDER BY sell_date ASC
), A as (
    SELECT DISTINCT *
    FROM Activities
    ORDER BY product
)

SELECT
    T.sell_date,
    T.num_sold,
    string_agg(a.product, ',' ORDER BY a.product) AS products
FROM
    A AS a
LEFT JOIN 
    T
ON a.sell_date = T.sell_date
GROUP BY T.sell_date, T.num_sold
ORDER BY T.sell_date ASC
