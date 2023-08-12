package besh

import (
	"Transaction/models"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func CategoryTrans() {

	products, _ := readProducts("data/products.json")
	transactions, _ := readTransaction("data/branch_pr_transaction.json")
	categories, _ := readCategoires("data/categories.json")

	catMap := make(map[int]models.Category)
	productCategoryMap := make(map[int]int)

	for _, category := range categories {
		catMap[category.ID] = category
	}

	for _, product := range products {
		productCategoryMap[product.ID] = product.CategoryID
	}

	countMap := make(map[int]map[int]int)

	for _, tr := range transactions {
		if countMap[tr.BranchID] == nil {
			countMap[tr.BranchID] = make(map[int]int)
		}
		countMap[tr.BranchID][productCategoryMap[tr.ProductID]]++
	}
	branchMap := make(map[int]models.Branch)
	for branchID, catCount := range countMap {
		branch := branchMap[branchID]

		for catID, count := range catCount {
			category := catMap[catID]
			fmt.Printf("Branch: %s, Category: %s, Count: %d\n", branch.Name, category.Name, count)
		}

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

func readCategoires(data string) ([]models.Category, error) {
	var categories []models.Category

	p, err := os.ReadFile(data)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}
	err = json.Unmarshal(p, &categories)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}
	return categories, nil
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

// countmap2:=make(map[int]countSUm)
// for _, v := range transactions {
// 	v:=countmap2[tr.BranchId]
// 	if tr.Type=="plus"{
