package sixreq

import (
	"Transaction/models"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func CalculateTransSum() {
	branches, _ := readBranches("data/branches.json")
	products, _ := readProducts("data/products.json")

	transactions, _ := readTransaction("data/branch_pr_transaction.json")

	plusBranchIdtransCount := make(map[int]int)
	minusBranchIdtransCount := make(map[int]int)

	plusProductSum := make(map[int]int)
	minusProductSum := make(map[int]int)
	branchMap := make(map[int]string)
	ProductMap := make(map[int]int)

	for _, b := range branches {
		branchMap[b.ID] = b.Name
	}
	for _, p := range products {
		ProductMap[p.ID] = p.Price
	}

	for _, tr := range transactions {
		if tr.Type == "plus" {
			plusBranchIdtransCount[tr.BranchID]++
			plusProductSum[tr.BranchID] += tr.Quantity * ProductMap[tr.ProductID]
		} else {
			minusBranchIdtransCount[tr.BranchID]++
			minusProductSum[tr.BranchID] += tr.Quantity * ProductMap[tr.ProductID]
		}
	}

	for branchId, branchName := range branchMap {
		fmt.Printf("%s- TransactPlus: %d  TransactMinus: %d  SumPLus:%d   SumMinus: %d\n",
			branchName, plusBranchIdtransCount[branchId], minusBranchIdtransCount[branchId], plusProductSum[branchId], minusProductSum[branchId],
		)
	}

}

// ================================READERS======================================

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
