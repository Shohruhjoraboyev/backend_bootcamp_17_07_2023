package bir

import (
	"Transaction/models"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
)

func Toptrans() {
	branches, _ := readBranches("data/branches.json")
	transactions, _ := readTransaction("data/branch_pr_transaction.json")
	countMap := make(map[int]int)
	BrMap := make(map[int]string)

	for _, v := range transactions {
		countMap[v.BranchID]++
	}

	for _, b := range branches {
		BrMap[b.ID] = b.Name
	}

	type Quant struct {
		BranchId int
		Quantity int
	}
	countSlices := make([]Quant, 0, len(countMap))

	for k, v := range countMap {
		countSlices = append(countSlices, Quant{BranchId: k, Quantity: v})
	}

	sort.Slice(countSlices, func(i, j int) bool {
		return countSlices[i].Quantity > countSlices[j].Quantity
	})
	for i, v := range countSlices {
		fmt.Printf("%d - BranchName: %s    Transaction Cuant:  %d\n", i+1, BrMap[v.BranchId], v.Quantity)
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
