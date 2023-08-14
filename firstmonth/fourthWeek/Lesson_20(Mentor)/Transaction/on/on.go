package on

import (
	"Transaction/models"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func GetUserTrans() {
	products, _ := readProducts("data/products.json")
	users, _ := readUsers("data/users.json")
	transactions, _ := readTransaction("data/branch_pr_transaction.json")
	UserNameSumMap := make(map[string]int)
	UserIdNameMap := make(map[int]string)
	ProductMap := make(map[int]int)

	for _, v := range users {
		UserIdNameMap[v.ID] = v.Name
	}
	for _, v := range products {
		ProductMap[v.ID] = v.Price
	}

	for _, v := range transactions {
		UserNameSumMap[UserIdNameMap[v.UserID]] += v.Quantity * ProductMap[v.ProductID]
	}
	for k, v := range UserNameSumMap {
		fmt.Printf("%s --%d\n", k, v)
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

func readUsers(data string) ([]models.Users, error) {
	var users []models.Users
	branch, err := os.ReadFile(data)
	if err != nil {
		log.Printf("Error while Read users: %+v", err)
		return nil, err
	}
	err = json.Unmarshal(branch, &users)
	if err != nil {
		log.Printf("Error while Unmarshal users: %+v", err)
		return nil, err
	}
	return users, nil
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
