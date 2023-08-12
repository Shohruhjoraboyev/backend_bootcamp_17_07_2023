package memory

import (
	"encoding/json"
	"errors"
	"fmt"
	"lesson_20/models"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

type transactionRepo struct {
	fileName string
}

func NewTransactionRepo(fileName string) *transactionRepo {
	return &transactionRepo{fileName: fileName}
}

func (t *transactionRepo) CreateTransaction(req models.CreateTransaction) (string, error) {
	id := uuid.NewString()
	transactions, err := t.read()
	if err != nil {
		return "", err
	}
	transactions = append(transactions, models.Transaction{
		Id:          id,
		Type:        req.Type,
		Amount:      req.Amount,
		Source_type: req.Source_type,
		Text:        req.Text,
		Sale_id:     req.Sale_id,
		Staff_id:    req.Staff_id,
		Created_at:  time.Now().Format("2006-01-02 15:04:05"),
	})
	err = t.write(transactions)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (t *transactionRepo) UpdateTransaction(req models.Transaction) (string, error) {
	transactions, err := t.read()
	if err != nil {
		return "", err
	}
	for i, v := range transactions {
		if v.Id == req.Id {
			transactions[i] = req
			transactions[i].Created_at = time.Now().Format("2006-01-02 15:04:05")
			err = t.write(transactions)
			if err != nil {
				return "", err
			}
			return "updated", nil
		}
	}
	return "", errors.New("not transaction found with ID")
}

func (t *transactionRepo) GetTransaction(req models.IdRequest) (models.Transaction, error) {
	transactions, err := t.read()
	if err != nil {
		return models.Transaction{}, err
	}
	for _, v := range transactions {
		if v.Id == req.Id {
			return v, nil
		}
	}
	return models.Transaction{}, errors.New("not transaction found with ID")
}

func (t *transactionRepo) GetAllTransaction(req models.GetAllTransactionRequest) (resp models.GetAllTransactionResponse, err error) {
	transactions, err := t.read()
	if err != nil {
		return models.GetAllTransactionResponse{}, err
	}
	start := req.Limit * (req.Page - 1)
	end := start + req.Limit

	var filtered []models.Transaction

	for _, v := range transactions {
		if strings.Contains(v.Text, req.Text) {
			filtered = append(filtered, v)
		}
	}

	if start > len(filtered) {
		resp.Transactions = []models.Transaction{}
		resp.Count = len(filtered)
		return
	} else if end > len(filtered) {
		return models.GetAllTransactionResponse{
			Transactions: filtered[start:],
			Count:        len(filtered),
		}, nil
	} else {
		return models.GetAllTransactionResponse{
			Transactions: filtered[start:end],
			Count:        len(filtered),
		}, nil
	}

}

func (t *transactionRepo) DeleteTransaction(req models.IdRequest) (resp string, err error) {
	transactions, err := t.read()
	if err != nil {
		return "", err
	}
	for i, v := range transactions {
		if v.Id == req.Id {
			transactions = append(transactions[:i], transactions[i+1:]...)
			err = t.write(transactions)
			if err != nil {
				return "", err
			}
			return "deleted", nil
		}
	}
	return "", errors.New("not found")
}

func (u *transactionRepo) read() ([]models.Transaction, error) {
	var (
		transactions []models.Transaction
	)

	data, err := os.ReadFile(u.fileName)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}

	err = json.Unmarshal(data, &transactions)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}

	return transactions, nil
}

func (u *transactionRepo) write(transasctions []models.Transaction) error {
	body, err := json.Marshal(transasctions)
	if err != nil {
		return err
	}
	err = os.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (t *transactionRepo) GetTopStaffs(req models.TopWorkerRequest) (str string, err error) {
	var branches []models.Branch
	var staffes []models.Staff
	data, err := os.ReadFile("data/branch.json")
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return "", err
	}
	err = json.Unmarshal(data, &branches)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return "", err
	}
	datastaff, err := os.ReadFile("data/staff.json")
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return "", err
	}
	err = json.Unmarshal(datastaff, &staffes)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return "", err
	}
	transactions, err := t.read()
	if err != nil {
		return "", err
	}
	// =======================time parsing=======
	startDate, err := time.Parse("2006-01-02", req.FromDate)
	if err != nil {
		fmt.Println("Error parsing start date:", err)
		return
	}
	endDate, err := time.Parse("2006-01-02", req.ToDate)
	if err != nil {
		fmt.Println("Error parsing end date:", err)
		return
	}
	staffIdAmountMap := make(map[string]int)
	for _, tr := range transactions {
		createdAt, err := time.Parse("2006-01-02 15:04:05", tr.Created_at)
		if err != nil {
			fmt.Println("Error parsing createdAt:", err)
			continue
		}
		if createdAt.After(startDate) && createdAt.Before(endDate) {
			staffIdAmountMap[tr.Staff_id] += int(tr.Amount)

		}

	}

	type fullStruct struct {
		BranchName string
		StaffName  string
		Amount     int
	}

	fullCountMap := make(map[string]fullStruct)
	for _, br := range branches {
		for _, s := range staffes {
			for staffId, amount := range staffIdAmountMap {
				if staffId == s.Id && br.Id == s.BranchId && req.Type == string(s.TypeId) {
					fullCountMap[string(s.TypeId)] = fullStruct{
						BranchName: br.Name,
						StaffName:  s.Name,
						Amount:     amount,
					}
					fmt.Printf("Branch: %s Staff: %s Earning: %d\n", br.Name, s.Name, amount)
				}
			}
		}
	}
	return
}
