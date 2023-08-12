package nireq

import (
	"Transaction/models"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func GetBranchProductQuant() {
	branchProducts, _ := readBranchProduct("data/branch_products.json")
	products, _ := readProducts("data/products.json")
	branches, _ := readBranches("data/branches.json")

	BProductMap := make(map[int]int)
	ProductMap := make(map[int]int)
	BranchMap := make(map[int]string)

	for _, v := range branches {
		BranchMap[v.ID] = v.Name
	}

	for _, v := range products {
		ProductMap[v.ID] += v.Price
	}
	for _, v := range branchProducts {
		BProductMap[v.BranchId] += v.Quantity * ProductMap[v.ProductId]
		fmt.Printf("BranchName: %s   Sum: %d\n", BranchMap[v.BranchId], BProductMap[v.BranchId])
	}

}

// ================================READERS======================================

func readBranchProduct(data string) ([]models.BranchProduct, error) {
	var branchProducts []models.BranchProduct

	d, err := os.ReadFile(data)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}
	err = json.Unmarshal(d, &branchProducts)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}
	return branchProducts, nil
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
