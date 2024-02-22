WITH T AS (
  SELECT user_id, action, COUNT(action) AS cnt
  FROM Confirmations
  GROUP BY user_id, action
), C AS (
  SELECT T.user_id AS user_id, T.cnt AS cnt
  FROM T
  WHERE T.action = 'confirmed'
), Ti AS (
  SELECT T.user_id AS user_id, T.cnt AS cnt
  FROM T
  WHERE T.action = 'timeout'
), Calc AS (
  SELECT
    s.user_id AS user_id,
    CASE
      WHEN C.cnt IS NULL THEN 0
      ELSE C.cnt
    END AS confirmed,
    CASE
      WHEN Ti.cnt IS NULL THEN 0
      ELSE Ti.cnt
    END AS timeout
  FROM Signups AS s
  LEFT JOIN C ON s.user_id = C.user_id
  LEFT JOIN Ti ON s.user_id = Ti.user_id
)

SELECT
  user_id,
  CASE
    WHEN timeout = 0 AND confirmed = 0 THEN 0
    ELSE ROUND(
      CAST (confirmed AS FLOAT) / CAST ((confirmed + timeout) AS FLOAT),
      2
    )
  END AS confirmation_rate
FROM Calc
