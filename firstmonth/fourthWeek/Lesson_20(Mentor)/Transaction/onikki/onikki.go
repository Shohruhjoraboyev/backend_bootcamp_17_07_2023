package onikki

import (
	"Transaction/models"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func GetProductInputOutput() {
	users, _ := readUsers("data/users.json")
	transactions, _ := readTransaction("data/branch_pr_transaction.json")
	plusMap := make(map[int]int)
	minusMap := make(map[int]int)

	for _, v := range transactions {
		if v.Type == "plus" {
			plusMap[v.UserID] += v.Quantity
		} else {
			minusMap[v.UserID] += v.Quantity

		}
	}

	fmt.Println("        Kiritilgan\tChiqarilgan")
	for _, user := range users {
		fmt.Printf("%s\t\t%d\t\t%d\n", user.Name, plusMap[user.ID], minusMap[user.ID])
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
