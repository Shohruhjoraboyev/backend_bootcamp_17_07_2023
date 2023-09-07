-- 1. Top branches by transaction count
CREATE OR REPLACE FUNCTION get_top_branches_by_transaction_count()
  RETURNS TABLE (branch_id INT, branch_name TEXT, transaction_count INT) 
  LANGUAGE plpgsql
  AS $$
BEGIN
  RETURN QUERY
  SELECT b.id, b.name, COUNT(bt.id) AS transaction_count
  FROM branch AS b
  LEFT JOIN branch_transaction AS bt ON b.id = bt.branch_id
  GROUP BY b.id, b.name
  ORDER BY transaction_count DESC;
END;
$$;

-- 2. Top branches by transaction sum
CREATE OR REPLACE FUNCTION get_top_branches_by_transaction_sum()
  RETURNS TABLE (branch_name TEXT, total_sum NUMERIC) 
  LANGUAGE plpgsql
  AS $$
BEGIN
  RETURN QUERY
  SELECT b.name, SUM(p.price * bt.quantity) AS total_sum
  FROM branch AS b
  JOIN branch_transaction AS bt ON b.id = bt.branch_id
  JOIN product AS p ON p.id = bt.product_id
  GROUP BY b.id, b.name
  ORDER BY total_sum DESC;
END;
$$;

-- 3. Top products by transaction count
CREATE OR REPLACE FUNCTION get_top_products_by_transaction_count()
  RETURNS TABLE (product_name TEXT, transaction_count INT) 
  LANGUAGE plpgsql
  AS $$
BEGIN
  RETURN QUERY
  SELECT p.name, COUNT(bt.id) AS transaction_count
  FROM product AS p
  JOIN branch_transaction AS bt ON bt.product_id = p.id
  GROUP BY p.id, p.name
  ORDER BY transaction_count DESC;
END;
$$ ;

-- 4. Top categories by transaction count
CREATE OR REPLACE FUNCTION get_top_categories_by_transaction_count()
  RETURNS TABLE (category_name TEXT, transaction_count INT) 
  LANGUAGE plpgsql
  AS $$
BEGIN
  RETURN QUERY
  SELECT c.name, COUNT(c.id) AS transaction_count
  FROM branch_transaction AS bt
  JOIN product AS p ON bt.product_id = p.id
  JOIN category AS c ON p.category_id = c.id
  GROUP BY c.id, c.name
  ORDER BY transaction_count DESC;
END;
$$;

-- 5. Number of transactions per category in each branch
CREATE OR REPLACE FUNCTION get_transactions_per_category_per_branch()
  RETURNS TABLE (branch_name TEXT, category_name TEXT, transaction_count INT)
  LANGUAGE plpgsql
   AS $$
BEGIN
  RETURN QUERY
  SELECT b.name, c.name, COUNT(bt.id) AS transaction_count
  FROM branch AS b
  JOIN branch_transaction AS bt ON b.id = bt.branch_id
  JOIN product AS p ON bt.product_id = p.id
  JOIN category AS c ON p.category_id = c.id
  GROUP BY b.id, b.name, c.id, c.name
  ORDER BY b.name, c.name;
END;
$$ ;

/*
	6. har bir branch nechta plus/minus transactionlar soni, plus/minus transactionlar summasini quyidagicha chiqarish:
	                    Transactions            Summ
	                    plus   minus        plus     minus
	    1. Branch1      53      20          853 000  278 000
	    2. Branch2      38      185         492 000  1 982 000
*/

CREATE OR REPLACE FUNCTION get_transactions_quantity_sum_plusMinus()
RETURNS TABLE(branch_name TEXT,trpluscount INT,trminuscount INT,TrsSumPlus INT,TrsSumMinus INT)
LANGUAGE  plpgsql
AS $$
BEGIN
 RETURN QUERY
 SELECT 
    b.name AS branch_name,
    COUNT(CASE WHEN t.type = 'plus' THEN t.id END) AS trpluscount,
    COUNT(CASE WHEN t.type = 'minus' THEN t.id END) AS trminuscount,
    SUM(CASE WHEN t.type = 'plus' THEN t.quantity * p.price END) AS TrSumPlus,
    SUM(CASE WHEN t.type = 'minus' THEN t.quantity * p.price END) AS TrSumMinus
FROM 
    branch AS b
JOIN 
    BrTransaction AS t ON b.id = t.branch_id
JOIN 
    product AS p ON t.product_id = p.id
GROUP BY 
    b.name;
 END;
 $$; 

