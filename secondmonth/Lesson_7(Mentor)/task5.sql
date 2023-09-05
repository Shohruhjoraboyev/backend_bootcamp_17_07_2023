/*
5.Block/function qiling: branchni productni sotilgan soniga qarab 50dan ko'p bo'lsa 'xit', 100dan ko'p bo'lsa 'top' aks holda 'yangi' deb chiqarsin.tartibi: top->xit->yangi
Example:

 product     |    status
 Cola           |    top
 Fanta         |     xit
 Rio             |     yangi
*/


SELECT
  p.name AS Product,
  CASE
    WHEN type = 'minus' THEN
      CASE
        WHEN SUM(bt.quantity) < 50 THEN 'Yangi'
        WHEN SUM(bt.quantity) > 50 THEN 'Xit'
        WHEN SUM(bt.quantity) > 100 THEN 'Top'
      END
  END AS Status
FROM  branch AS b
  JOIN branch_transaction AS bt ON b.id = bt.branch_id
  JOIN product AS p ON bt.product_id = p.id
GROUP BY Product, Status;

--plpgsql
Create or REPLACE FUNCTION Status()
RETURNS TABLE(Product Varchar,Status Varchar)
LANGUAGE plpgsql
AS
$$
BEGIN
RETURN QUERY
SELECT
  p.name AS Product,
  CASE
    WHEN type = 'minus' THEN
      CASE
        WHEN SUM(bt.quantity) < 50 THEN 'Yangi'
        WHEN SUM(bt.quantity) > 50 THEN 'Xit'
        WHEN SUM(bt.quantity) > 100 THEN 'Top'
      END
  END AS Status
FROM  branch AS b
  JOIN branch_transaction AS bt ON b.id = bt.branch_id
  JOIN product AS p ON bt.product_id = p.id
GROUP BY Product, Status;

END;
$$

