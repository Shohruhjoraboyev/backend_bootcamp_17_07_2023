package ontort

import (
	"Transaction/models"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func DailyInputOutput() {
	transactions, _ := readTransaction("data/branch_pr_transaction.json")
	products, _ := readProducts("data/products.json")
	users, _ := readUsers("data/users.json")
	productMap := make(map[int]int)
	userMap := make(map[int]string)

	for _, p := range products {
		productMap[p.ID] += p.Price
	}
	for _, s := range users {
		userMap[s.ID] = s.Name
	}

	type DailyS struct {
		countPlus  int
		countMinus int
	}

	DailyMap := make(map[int]map[string]DailyS)
	for _, v := range transactions {
		if v.Type == "plus" {
			if DailyMap[v.UserID] == nil {
				DailyMap[v.UserID] = make(map[string]DailyS)
			}
			dailyS := DailyMap[v.UserID][v.CreatedAt]
			dailyS.countPlus += v.Quantity
			DailyMap[v.UserID][v.CreatedAt] = dailyS
		} else {
			if DailyMap[v.UserID] == nil {
				DailyMap[v.UserID] = make(map[string]DailyS)
			}
			dailyS := DailyMap[v.UserID][v.CreatedAt]
			dailyS.countMinus += v.Quantity
			DailyMap[v.UserID][v.CreatedAt] = dailyS
		}

	}

	for urId, brCount := range DailyMap {
		Counter := 0
		PlTrCount := 0
		mTrCount := 0
		for _, count := range brCount {
			Counter++
			PlTrCount += count.countPlus
			mTrCount += count.countMinus
		}
		fmt.Printf("%s %d %d\n", userMap[urId], productMap[urId]*(PlTrCount/Counter), productMap[urId]*(mTrCount/Counter))
	}
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
