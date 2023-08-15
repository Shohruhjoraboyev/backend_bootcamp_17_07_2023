package onuch

import (
	"Transaction/models"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func DailyInputOutput() {
	transactions, _ := readTransaction("data/branch_pr_transaction.json")
	branches, _ := readBranches("data/branches.json")
	branchMap := make(map[int]string)

	for _, b := range branches {
		branchMap[b.ID] = b.Name
	}

	type DailyS struct {
		countPlus  int
		countMinus int
	}

	DailyMap := make(map[int]map[string]DailyS)
	for _, v := range transactions {
		if v.Type == "plus" {
			if DailyMap[v.BranchID] == nil {
				DailyMap[v.BranchID] = make(map[string]DailyS)
			}
			dailyS := DailyMap[v.BranchID][v.CreatedAt]
			dailyS.countPlus += v.Quantity
			DailyMap[v.BranchID][v.CreatedAt] = dailyS
		} else {
			if DailyMap[v.BranchID] == nil {
				DailyMap[v.BranchID] = make(map[string]DailyS)
			}
			dailyS := DailyMap[v.BranchID][v.CreatedAt]
			dailyS.countMinus += v.Quantity
			DailyMap[v.BranchID][v.CreatedAt] = dailyS
		}

	}

	for brId, brCount := range DailyMap {
		Counter := 0
		PlTrCount := 0
		mTrCount := 0
		for _, count := range brCount {
			Counter++
			PlTrCount += count.countPlus
			mTrCount += count.countMinus
		}
		fmt.Printf("%s %d %d\n", branchMap[brId], PlTrCount/Counter, mTrCount/Counter)
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
