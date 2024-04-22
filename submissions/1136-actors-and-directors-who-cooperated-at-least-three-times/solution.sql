SELECT
    actor_id, director_id
FROM (
    SELECT actor_id, director_id, COUNT(*) AS cnt
    FROM ActorDirector
    GROUP BY actor_id, director_id
    HAVING cnt >= 3
) AS t
