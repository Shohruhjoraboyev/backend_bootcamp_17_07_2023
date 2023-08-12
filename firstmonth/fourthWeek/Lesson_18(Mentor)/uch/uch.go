package uch

import (
	"Transaction/models"
	"encoding/json"
	"fmt"
	"sort"

	"log"
	"os"
)

func TopTransProduct() {
	products, _ := readProducts("data/products.json")
	transactions, _ := readTransaction("data/branch_pr_transaction.json")
	productMap := make(map[int]int)
	cuantMap := make(map[int]int)

	for _, product := range products {
		productMap[product.ID] = product.Price
	}
	for _, v := range transactions {
		cuantMap[v.ProductID] += v.Quantity
	}

	type Quant struct {
		ProductId int
		Quantity  int
	}
	countSlices := make([]Quant, 0, len(cuantMap))

	for k, v := range cuantMap {
		countSlices = append(countSlices, Quant{ProductId: k, Quantity: v})
	}

	sort.Slice(countSlices, func(i, j int) bool {
		return countSlices[i].Quantity > countSlices[j].Quantity
	})
	for i, v := range countSlices {
		fmt.Printf("%d --  Product ID: %d  Transuction Sum: %d\n", i+1, v.ProductId, v.Quantity*productMap[v.ProductId])
	}

}

// ================================READERS======================================

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

func readBranches(data string) ([]models.Branch, error) {
	var branches []models.Branch

	branch, err := os.ReadFile(data)
	if err != nil {
		log.Printf("Error while Read branch: %+v", err)
		return nil, err
	}
	err = json.Unmarshal(branch, &branches)
	if err != nil {
		log.Printf("Error while Unmarshal branch: %+v", err)
		return nil, err
	}
	return branches, nil
}

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
