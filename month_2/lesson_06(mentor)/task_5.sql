-- 5.Block/function qiling: branchni productni sotilgan soniga qarab 50dan ko'p bo'lsa 'xit', 100dan ko'p bo'lsa 'top' aks holda 'yangi' deb chiqarsin.tartibi: top->xit->yangi
-- Example:

--  product       |     status
--  Cola          |     top
--  Fanta         |     xit
--  Rio           |     yangi
DO $$
DECLARE
    products_data RECORD;
    product_name VARCHAR;
    product_sum INT;
    product_status VARCHAR;
BEGIN
    FOR products_data IN
        SELECT DISTINCT p.name, SUM(bt.quantity) AS total_quantity
        FROM branch_transaction bt
        JOIN product p ON p.id = bt.product_id
        WHERE bt.type = 'minus'
        GROUP BY p.name
    LOOP
        product_name := products_data.name; 
        product_sum := products_data.total_quantity;

        product_status := CASE
            WHEN product_sum > 100 THEN 'top'
            WHEN product_sum > 50 THEN 'hit'
            ELSE 'yangi'
        END;
    
        RAISE NOTICE 'Product: %, Status: %', product_name, product_status;
    END LOOP;
END;
$$;