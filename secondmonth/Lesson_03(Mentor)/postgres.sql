								  --  1.transactionlar soni bo'yicha top branches

									SELECT b.id, b.name, COUNT(bt.id) AS transaction_count
						      FROM branch AS b
						      LEFT JOIN branch_transaction AS bt ON b.id = bt.branch_id
						      GROUP BY b.id
						      ORDER BY transaction_count DESC;

								  --  2.transactionlar summasi bo'yicha top branches

									 SELECT  b.name,SUM(p.price*bt.quantity) AS total_sum
				          FROM branch AS b
				          JOIN branch_transaction AS bt ON b.id = bt.branch_id
				          JOIN product AS p ON p.id = bt.product_id
				          GROUP BY  b.id
				          ORDER BY total_sum;

								  --  3.transactionda bo'lgan top productlar
									 SELECT p.name, COUNT(bt.id) AS transaction_count
		               FROM product AS p
		               JOIN branch_transaction AS bt ON bt.product_id = p.id
		               GROUP BY p.id
		               ORDER BY transaction_count DESC;

                   