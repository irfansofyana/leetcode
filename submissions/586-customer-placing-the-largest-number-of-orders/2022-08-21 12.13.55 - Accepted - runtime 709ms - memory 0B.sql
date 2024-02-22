# Write your MySQL query statement below
SELECT t.customer_number
FROM (
    SELECT COUNT(customer_number) AS cnt, customer_number
    FROM Orders
    GROUP BY customer_number
    ORDER BY cnt DESC
) AS t
LIMIT 1
