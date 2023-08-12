package eightreq

import (
	"Transaction/models"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func GetProductInputOutput() {
	products, _ := readProducts("data/products.json")
	transactions, _ := readTransaction("data/branch_pr_transaction.json")
	plusMap := make(map[int]int)
	minusMap := make(map[int]int)

	for _, v := range transactions {
		if v.Type == "plus" {
			plusMap[v.ProductID] += v.Quantity
		} else {
			minusMap[v.ProductID] += v.Quantity

		}
	}

	for product_id, name := range products {
		fmt.Printf("%s- Kiritilgan: %d    Chiqarilgan: %d\n", name.Name, plusMap[product_id], minusMap[product_id])
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

func readProducts(data string) ([]models.Product, error) {
	var products []models.Product

	p, err := os.ReadFile(data)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}
	err = json.Unmarshal(p, &products)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}
	return products, nil
}
