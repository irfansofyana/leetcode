SELECT DISTINCT
    customer_id
FROM Customer
WHERE customer_id NOT IN (
    SELECT DISTINCT
        T.customer_id
    FROM
        (SELECT DISTINCT customer_id, Product.product_key AS product_key
        FROM Customer
        CROSS JOIN Product) AS T
    LEFT JOIN 
        (
            SELECT *
            FROM Customer
        ) AS c
    ON T.customer_id = c.customer_id AND T.product_key = c.product_key
    WHERE c.customer_id IS NULL AND c.product_key IS NULL
)

