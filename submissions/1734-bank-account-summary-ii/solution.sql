# Write your MySQL query statement below

SELECT u.name, l.balance
FROM
Users AS u LEFT JOIN
(
    SELECT account, SUM(amount) AS balance
    FROM Transactions
    GROUP BY account
) AS l
ON u.account = l.account
WHERE l.balance > 10000
