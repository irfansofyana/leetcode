# Write your MySQL query statement below


SELECT DISTINCT(p.Email)
FROM Person AS p
WHERE p.Id NOT IN
(
    SELECT Id FROM(
        SELECT MIN(Id) as Id, Email
        FROM Person
        GROUP BY Email
    ) AS T1
)
