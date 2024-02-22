# Write your MySQL query statement below


SELECT
    tmp.stock_name,
    SUM(tmp.gain) AS capital_gain_loss
FROM (    
        SELECT 
                stock_name,
                CASE
                    WHEN operation='Sell' THEN price
                    ELSE -1*price
                END AS gain
        FROM Stocks
) AS tmp
GROUP BY tmp.stock_name
