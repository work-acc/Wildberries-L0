-- #1
INSERT INTO Orders (
    Order_Uid,
    Data
) VALUES ($1, $2);

-- #2
SELECT
    Order_Uid,
    Data
FROM Orders;
