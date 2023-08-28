/*
	6. har bir branch nechta plus/minus transactionlar soni, plus/minus transactionlar summasini quyidagicha chiqarish:
	                    Transactions            Summ
	                    plus   minus        plus     minus
	    1. Branch1      53      20          853 000  278 000
	    2. Branch2      38      185         492 000  1 982 000
*/
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
/*
    	7. har bir kunda kirgan productlar sonini kamayish tartibida chiqarish:
	        kun         soni
	    1. 2023-08-04   789
	    2. 2023-08-12   634
*/ 

SELECT 
    t.createdat AS kun,
    SUM(t.quantity) AS soni
FROM 
    BrTransaction AS t
WHERE 
    t.type = 'plus'
GROUP BY 
    t.createdat
ORDER BY 
    kun;

/*
    	8. Product qancha kiritilgan va chiqarilganligi jadvali:
	    Name    Kiritilgan  Chiqarilgan
	    Olma     345            847
	    Cola     374            219
	    ....     ...       ...   ....
*/
    

SELECT 
    p.name AS Name
    SUM(CASE WHEN t.type = 'plus' THEN t.quantity  END) AS Kiritilgan,
    SUM(CASE WHEN t.type = 'minus' THEN t.quantity END) AS Chiqarilgan
FROM 
    products AS p
JOIN 
    BrTransaction AS t ON t.product_id =p.id
GROUP BY 
    b.name;
 /*
 	9. Filialda qancha summalik product borligi jadvali:

	1. Branch1        853 000
	2. Branch2      1 982 000
	...    ...   ...   ...  ...
 */   

 SELECT 
    b.name AS branch_name,
    SUM( brp.quantity * p.price ) AS TrSum,
FROM 
    branch AS b
JOIN 
    Brproduct AS brp ON b.id = brp.branch_id
JOIN 
    product AS p ON brp.product_id = p.id
GROUP BY 
    b.name;

   	/*
		10. har bir user transaction qilgan summasi jadvali:

		1  Akrom   1 349 000
		3  Ilhom   2 974 000
		...  ...  ...
    */
  SELECT 
   u.name AS Name,
 SUM(u.quantity * p.price)
  FROM  
  BrTransaction AS t
  JOIN
  products AS p ON t.product_id =p.id
  JOIN
   users AS u ON u.id=t.user_id
  GROUP BY
  u.name 

 /*
 		11. har bir user kun bo'yicha nechta va necha sumlik transaction qilgani jadvali:

		1 Akrom 2023-01-01  12  382 000
		2 Suhrob 2023-03-05  32  745 000
		...   ...   ...   ...  ....
 */
   SELECT 
   u.name as Name,
   t.createdat AS CreatedAt,
   COUNT(t.quantity), 
   SUM(t.quantity * p.price)
  FROM  
  BrTransaction AS t
  JOIN
  products AS p ON t.product_id =p.id
  JOIN
   users AS u ON u.id=t.user_id
  GROUP BY
  u.name 

  /*
  		12. har bir user qancha product kiritgani va chiqargani jadvali:
		         kiritgan  chiqargan
		1 Akrom    12         84
		2 Suhrob   54         33

  */
  SELECT 
    u.name AS Name
    SUM(CASE WHEN t.type = 'plus' THEN t.quantity  END) AS Kiritilgan,
    SUM(CASE WHEN t.type = 'minus' THEN t.quantity END) AS Chiqarilgan
FROM
  BrTransaction AS t
  JOIN
  products AS p ON t.product_id =p.id
  JOIN
   users AS u ON u.id=t.user_id
  GROUP BY
  u.name 

	/*
		13. Har kuni o'rtacha qancha product kiritilgani va chiqarilgani bo'yicha jadval:
		    branch      o'rtacha+   o'rtacha-
		1. Chilonzor      73         34
		2. MGorkiy        60         75
		...    ...     ...    ...       ...
    */    