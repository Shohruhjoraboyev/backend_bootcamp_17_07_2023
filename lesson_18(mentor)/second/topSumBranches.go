package second

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"task/models"
)

// 2.transactionlar summasi bo'yicha top branches
// {
// 	"id": 4,
// 	"name": "Uacademy",
// 	"address": "Shahriobod ko'chasi"
//   }
// {
//     "id": 1,
//     "branch_id": 4,
//     "product_id": 3,
//     "type": "plus",
//     "quantity": 74,
//     "created_at": "2023-08-09 20:05:37"
//   },
// {
//     "id": 2,
//     "name": "orange",
//     "price": 3000,
//     "category_id": 1
//   },

func CalculateSumOfPriceTopBranches() {
	branches, _ := readBranches("data/branches.json")
	productes, _ := readProducts("data/products.json")
	transactions, _ := readTransaction("data/branch_pr_transaction.json")

	branchSums := make(map[int]int)

	for _, p := range productes {
		for _, t := range transactions {
			for _, b := range branches {
				if t.BranchID == b.ID && t.ProductID == p.Id {
					for _, transaction := range transactions {
						// sum := t.Quantity * p.Price
						branchSums[transaction.BranchID] += t.Quantity * p.Price
					}
				}
			}
		}
	}

	fmt.Println(branchSums)

	// sort.Slice(sortedBranches, func(i, j int) bool {
	// 	return sortedBranches[i].Count > sortedBranches[j].Count
	// })

	// for _, v := range sortedBranches {
	// 	fmt.Printf("Branch: %s: Total Transactions: %d\n", v.BranchName, v.Count)
	// }
}

// ================================READERS======================================

func readProducts(data string) ([]models.Products, error) {
	var products []models.Products

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
