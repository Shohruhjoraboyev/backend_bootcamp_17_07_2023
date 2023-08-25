package task4

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"task/models"

	_ "github.com/lib/pq"
)

// 4.transactionda bo'lgan top categorylar
func TopTransactionCategory() {
	db, err := sql.Open("postgres", "postgres://postgres:Muhammad@localhost:5432/json?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := `
	SELECT c.name AS category_name, COUNT(p.id) AS transaction_count
	FROM branch_transaction bt
	JOIN product p ON bt.product_id = p.id
	JOIN category c ON p.category_id = c.id
	GROUP BY c.id
	ORDER BY transaction_count DESC
	`
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var t models.Task2
		err := rows.Scan(&t.BranchName, &t.Sum)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s - %d\n", t.BranchName, t.Sum)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}
}

// ================================READERS======================================
func readTransaction(data string) ([]models.Transaction, error) {
	var transactions []models.Transaction

	d, err := os.ReadFile(data)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}
	err = json.Unmarshal(d, &transactions)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}
	return transactions, nil
}

func readTCategory(data string) ([]models.ProductTop, error) {
	var categories []models.ProductTop

	d, err := os.ReadFile(data)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}
	err = json.Unmarshal(d, &categories)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}
	return categories, nil
}
