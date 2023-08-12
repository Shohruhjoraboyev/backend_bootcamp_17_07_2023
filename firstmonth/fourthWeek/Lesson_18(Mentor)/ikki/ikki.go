package ikki

import (
	"Transaction/models"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
)

func TopTransactBranch() {
	branches, _ := readBranches("data/branches.json")
	products, _ := readProducts("data/products.json")
	transactions, _ := readTransaction("data/branch_pr_transaction.json")
	productMap := make(map[int]int)
	branchMap := make(map[int]string)
	cuantMap := make(map[int]int)

	for _, b := range branches {
		branchMap[b.ID] = b.Name
	}

	for _, product := range products {
		productMap[product.ID] = product.Price
	}
	for _, v := range transactions {
		cuantMap[v.BranchID] += v.Quantity
	}

	type Quant struct {
		BranchId int
		Quantity int
	}
	countSlices := make([]Quant, 0, len(cuantMap))

	for k, v := range cuantMap {
		countSlices = append(countSlices, Quant{BranchId: k, Quantity: v})
	}

	sort.Slice(countSlices, func(i, j int) bool {
		return countSlices[i].Quantity > countSlices[j].Quantity
	})

	for i, v := range countSlices {
		fmt.Printf("%d --  BranchName: %s   Sum:  %d\n ", i+1, branchMap[v.BranchId], v.Quantity*productMap[v.BranchId])

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
