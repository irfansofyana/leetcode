WITH T1 AS (
    SELECT 
        customer_id,
        MIN(order_date) AS first_order
    FROM
        Delivery
    GROUP BY customer_id
),
T2 AS (
    SELECT
        d.customer_id,
        d.order_date,
        d.customer_pref_delivery_date,
        T1.first_order
    FROM
        Delivery AS d
    LEFT JOIN
        T1
    ON 
        d.customer_id = T1.customer_id
)

SELECT 
    COALESCE(
        ROUND(
            100.00 *
            (
                SELECT
                    COUNT(*)
                FROM 
                    T2
                WHERE
                    T2.customer_pref_delivery_date = T2.first_order
            ) / (
                SELECT COUNT(DISTINCT customer_id)
                FROM T2
            ),
            2
        )   
    ) AS immediate_percentage

