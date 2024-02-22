# Write your MySQL query statement below

SELECT t2.Name AS Department, e.Name AS Employee, e.Salary AS Salary
FROM Employee AS e
JOIN (
    SELECT d.Id, d.Name, t.max_salary
    FROM Department AS d
    LEFT JOIN (
        SELECT DepartmentId, Max(Salary) AS max_salary
        FROM Employee
        GROUP BY DepartmentId
    ) AS t
    ON d.Id = t.DepartmentId
) AS t2
ON e.salary = t2.max_salary AND e.DepartmentId = t2.Id
