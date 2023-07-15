SELECT
    manager.name AS name
FROM
   Employee LEFT JOIN (
       SELECT name, id
       FROM Employee
   ) AS manager
   ON Employee.managerId = manager.id
GROUP BY manager.id
HAVING COUNT(manager.id) >= 5
