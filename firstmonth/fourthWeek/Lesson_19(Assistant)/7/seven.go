package seventreq

import (
	"Transaction/models"
	"encoding/json"
	"fmt"
	"sort"

	"log"
	"os"
)

func CalculateDailyProduct() {
	transactions, _ := readTransaction("data/branch_pr_transaction.json")
	plusMap := make(map[string]int)

	for _, v := range transactions {
		if v.Type == "plus" {
			plusMap[v.CreatedAt] += v.Quantity
		}
	}

	type CountSlice struct {
		Kun  string
		soni int
	}
	countSlices := make([]CountSlice, 0, len(plusMap))
	for k, v := range plusMap {
		countSlices = append(countSlices, CountSlice{Kun: k, soni: v})
	}
	sort.Slice(countSlices, func(i, j int) bool {
		return countSlices[i].soni > countSlices[j].soni
	})

	for i, v := range countSlices {
		fmt.Printf("%d -  Kun: %s   soni: %d\n", i+1, v.Kun, v.soni)
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
