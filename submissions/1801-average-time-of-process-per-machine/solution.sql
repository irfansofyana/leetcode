WITH st AS (
  SELECT *
  FROM Activity
  WHERE activity_type = 'start'
), et AS (
  SELECT *
  FROM Activity
  WHERE activity_type = 'end'
), m AS (
  SELECT DISTINCT
    A.machine_id,
    A.process_id,
    et.timestamp - st.timestamp AS duration
  FROM
    Activity AS A
    LEFT JOIN st ON A.machine_id = st.machine_id AND A.process_id = st.process_id
    LEFT JOIN et ON A.machine_id = et.machine_id AND A.process_id = et.process_id
)

SELECT
  machine_id,
  ROUND(AVG(duration), 3) AS processing_time
FROM m
GROUP BY machine_id
