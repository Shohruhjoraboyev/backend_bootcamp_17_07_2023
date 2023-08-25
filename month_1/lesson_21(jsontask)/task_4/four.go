package task4

import (
	"encoding/json"
	"log"
	"os"
	"task/models"
)

// 4.transactionda bo'lgan top categorylar
func TopTransactionCategory() {
	transactions, _ := readTransaction("data/branch_pr_transaction.json")
	categories, _ := readTCategory("data/categories.json")
	products, _ := readProduct("data/categories.json")

	categoryMap := make(map[int]string)
	productCategoryId := make(map[int]int)

	for _, v := range categories {
		categoryMap[v.Id] = v.Name
	}

	// for catId, count := range productMap {
	// 	fmt.Printf("%s  %d", categoryMap[catId], count)
	// }

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

func readProduct(data string) ([]models.Products, error) {
	var productes []models.Products

	d, err := os.ReadFile(data)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}
	err = json.Unmarshal(d, &productes)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}
	return productes, nil
}
