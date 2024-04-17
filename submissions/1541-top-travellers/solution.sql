SELECT t.name, t.travelled_distance
FROM (
    SELECT u.name AS name, u.id AS id, COALESCE(SUM(r.distance), 0) AS travelled_distance
    FROM Users AS u LEFT JOIN Rides AS r on u.id = r.user_id
    GROUP BY name, id
    ORDER BY travelled_distance DESC, name ASC
) AS t
