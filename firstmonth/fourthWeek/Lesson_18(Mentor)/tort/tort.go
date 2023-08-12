package tort

import (
	"Transaction/models"
	"encoding/json"
	"fmt"
	"sort"

	"log"
	"os"
)

func TopTransactCategory() {
	transactions, _ := readTransaction("data/branch_pr_transaction.json")
	categories, _ := readCategory("data/categories.json")
	cuantMap := make(map[int]int)
	catogoryMap := make(map[int]string)
	for _, category := range categories {
		catogoryMap[category.ID] = category.Name
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
		fmt.Printf("%d  Category name: %s  Transactions Quantity: %d\n", i+1, catogoryMap[v.ProductId], v.Quantity)
	}
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

func readCategory(data string) ([]models.Category, error) {
	var categories []models.Category
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
