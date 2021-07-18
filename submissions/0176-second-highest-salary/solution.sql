# Write your MySQL query statement below
SELECT MAX(t.Salary) AS "SecondHighestSalary"
FROM (
    SELECT Salary
    FROM Employee
    WHERE Salary < (SELECT MAX(Salary) FROM Employee)
) AS t
