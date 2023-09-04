-- start a transaction
BEGIN;

-- create branch_pr_transactions
INSERT INTO branch_pr_transactions(branch_id,user_id,product_id,type,quantity)
VALUES(2,1,3,'plus',45);

-- update branch products
UPDATE branch_products
SET quantity = quantity + 45
WHERE product_id = 3,branch_id=2;

-- commit the transaction
COMMIT;