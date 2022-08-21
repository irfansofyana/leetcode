# Write your MySQL query statement below
SELECT name
FROM SalesPerson
WHERE sales_id NOT IN (
    SELECT DISTINCT(s.sales_id)
    FROM Orders AS o
    LEFT JOIN Company AS c
    ON o.com_id = c.com_id
    LEFT JOIN SalesPerson AS s
    ON o.sales_id = s.sales_id
    WHERE c.name = 'RED'
)


