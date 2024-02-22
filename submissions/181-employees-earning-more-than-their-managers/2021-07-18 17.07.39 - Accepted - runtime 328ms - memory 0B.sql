# Write your MySQL query statement below
SELECT E1.Name AS "Employee"
FROM Employee AS E1, (
    SELECT *
    FROM Employee
) AS E2
WHERE E1.ManagerId = E2.Id AND E1.Salary > E2.Salary
