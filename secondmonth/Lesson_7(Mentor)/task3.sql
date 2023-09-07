/*
 3.Masala: json data bo'yicha branchdan product_id lik productdan n quantity sotish kerak.
 Shunday plpgsql procedure/function yozing agar branchda product yetarli bo'lsa, 
transaction create qilib branchdagi productlardan n quantity kamaytirsin, 
 agar kam bo'lsa 'product yetarli emas' xabari chiqsin
 */

CREATE OR REPLACE FUNCTION sellProduct(
  branchId int,
  productId int,
  req_quantity int
)
RETURNS branch_products
LANGUAGE plpgsql
AS $$
DECLARE 
  product_quantity int;
BEGIN
  -- Product malumotlarini olish
  SELECT quantity INTO product_quantity
  FROM branch_products
  WHERE branch_id = branchId AND product_id = productId;

  IF product_quantity IS NOT NULL AND product_quantity >= req_quantity THEN
    UPDATE branch_products
    SET quantity = quantity - req_quantity
    WHERE branch_id = branchId AND product_id = productId;
    COMMIT;
    RAISE NOTICE 'Muvaffaqiyatli sotildi, Qolgan Miqdor %', quantity;
      RETURN (SELECT * FROM branch_products WHERE branch_id = branchId AND product_id = productId);

  ELSE 
    RAISE NOTICE 'Maxsulot yetarli emas';
    ROLLBACK;
  END IF;
  
END;
$$;

