package task6

import (
	"database/sql"
	"fmt"
	"task/models"

	_ "github.com/lib/pq"
)

// 6. har bir branch nechta plus/minus transactionlar soni, plus/minus transactionlar summasini quyidagicha chiqarish:
//                     Transactions            Summ
//                     plus   minus        plus     minus
//     1. Branch1      53      20          853 000  278 000
//     2. Branch2      38      185         492 000  1 982 000

func PlusMinus() {

	db, err := sql.Open("postgres", "postgres://postgres:Muhammad@localhost:5432/json?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//============================ THIS IS MY OWN QUERY =============================
	// query := `
	// SELECT b.name, COUNT(b.id) as tr_plus_count, CAST(SUM(p.price*bt.quantity) AS int)  as plusSum
	// FROM branch_transaction bt
	// JOIN branch b ON bt.branch_id = b.id
	// JOIN product p ON p.id = bt.product_id
	// WHERE bt.type = 'plus'
	// GROUP BY b.name
	// UNION
	// SELECT b.name, COUNT(b.id) as tr_minus_count, CAST(SUM(p.price*bt.quantity) AS int)  as minusSum
	// FROM branch_transaction bt
	// JOIN branch b ON bt.branch_id = b.id
	// JOIN product p ON p.id = bt.product_id
	// WHERE bt.type = 'minus'
	// GROUP BY b.name
	// `

	//  ==========  BY CHATGPT ============= I asked if there are other ways of combining two select
	query := `
	SELECT b.name,
       COUNT(CASE WHEN bt.type = 'plus' THEN b.id END) AS tr_plus_count,
       COALESCE(SUM(CASE WHEN bt.type = 'plus' THEN p.price * bt.quantity END)::int, 0) AS plusSum,
       COUNT(CASE WHEN bt.type = 'minus' THEN b.id END) AS tr_minus_count,
       COALESCE(SUM(CASE WHEN bt.type = 'minus' THEN p.price * bt.quantity END)::int, 0) AS minusSum
	FROM branch b
	LEFT JOIN branch_transaction bt ON bt.branch_id = b.id
	LEFT JOIN product p ON p.id = bt.product_id
	GROUP BY b.name;
`

	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var t models.Task6
		err := rows.Scan(&t.BranchaName, &t.TranPlusCount, &t.TranMinusCount, &t.TranPlusSum, &t.TranMinusSum)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s: t_plus: %d t_minus: %d - s_plus: %d s_minus: %d\n",
			t.BranchaName, t.TranPlusCount, t.TranMinusCount, t.TranPlusSum, t.TranMinusSum)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

}
