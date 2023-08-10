package fifth

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"task/models"
)

// 5.har bir branchda har bir categorydan qancha transaction bo'lgani
func TopBranchTransactionCategory() {
	productes, _ := readProducts("data/products.json")
	transactions, _ := readTransaction("data/branch_pr_transaction.json")
	categories, _ := readTCategory("data/categories.json")

	var transactionCategoryCount = make(map[int]int) //[productId]count
	var transactionBranchCount = make(map[int]int)   //[branchID]count

	for _, p := range productes {
		for _, t := range transactions {
			if t.ProductID == p.Id {
				transactionCategoryCount[p.CategoryId]++
				transactionBranchCount[t.BranchID]++
			}
		}

	}

	var sortedTopProducts []models.ProductTop
	for _, c := range categories {
		for k, t := range transactionCategoryCount {
			if c.Id == k {
				sortedTopProducts = append(sortedTopProducts, models.ProductTop{
					Id:    k,
					Name:  c.Name,
					Count: t,
				})
			}
		}
	}

	sort.Slice(sortedTopProducts, func(i, j int) bool {
		return sortedTopProducts[i].Count > sortedTopProducts[j].Count
	})

	for _, v := range sortedTopProducts {
		fmt.Printf("%s: %d times transacted\n", v.Name, v.Count)
	}
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

func readTCategory(data string) ([]models.ProductTop, error) {
	var categories []models.ProductTop

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
