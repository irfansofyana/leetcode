SELECT 
    p.product_name,
    t.unit
FROM
(
    SELECT
        product_id,
        DATE_FORMAT(order_date, '%Y-%m') AS month_order,
        SUM(unit) AS unit
    FROM Orders
    WHERE DATE_FORMAT(order_date, '%Y-%m') = '2020-02'
    GROUP BY product_id, month_order
    HAVING unit >= 100
) AS t
LEFT JOIN Products AS p
on t.product_id = p.product_id
