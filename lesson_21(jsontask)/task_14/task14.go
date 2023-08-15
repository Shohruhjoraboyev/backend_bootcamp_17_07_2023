package task14

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"task/models"
)

/*
15. Har kuni o'rtacha user qancha summa product kiritgani va chiqargani bo'yicha jadval:
    branch      o'rtacha+        o'rtacha-
1. Anvar          370 000         435 000
2. Shuhrat        60 000          875 000
...    ...     ...    ...       ...
*/

func Task14() {
	trasnsactions, _ := readTransaction("data/branch_pr_transaction.json")
	users, _ := readProducts("data/users.json")
	// branchId va name
	branchIdName := make(map[int]string)
	for _, b := range users {
		branchIdName[b.ID] = b.Name
	}

	type PlusMinus struct {
		plusQuantity  int
		minusQuantity int
	}

	// branchid created count
	plusBranchIdTimeCount := make(map[int]map[string]PlusMinus)

	for _, tr := range trasnsactions {
		if tr.Type == "plus" {
			if _, ok := plusBranchIdTimeCount[tr.BranchID]; !ok {
				plusBranchIdTimeCount[tr.BranchID] = make(map[string]PlusMinus)
			}
			v := plusBranchIdTimeCount[tr.BranchID][tr.CreatedAt[:11]]
			v.plusQuantity += tr.Quantity
			plusBranchIdTimeCount[tr.BranchID][tr.CreatedAt[:11]] = v

		} else {
			if _, ok := plusBranchIdTimeCount[tr.BranchID]; !ok {
				plusBranchIdTimeCount[tr.BranchID] = make(map[string]PlusMinus)
			}
			v := plusBranchIdTimeCount[tr.BranchID][tr.CreatedAt[:11]]
			v.minusQuantity += tr.Quantity
			plusBranchIdTimeCount[tr.BranchID][tr.CreatedAt[:11]] = v
		}
	}

	for branch_id, innerMap := range plusBranchIdTimeCount {
		plusQuantity := 0
		minusQuantity := 0
		transactionCount := 0
		for _, PlusMinus := range innerMap {
			transactionCount++
			plusQuantity += PlusMinus.plusQuantity
			minusQuantity += PlusMinus.minusQuantity
		}
		fmt.Printf("%s, plus: %d minus: %d\n", branchIdName[branch_id], plusQuantity/transactionCount, minusQuantity/transactionCount)

	}
}

// ================================READERS======================================
func readProducts(data string) ([]models.User, error) {
	var products []models.User

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
