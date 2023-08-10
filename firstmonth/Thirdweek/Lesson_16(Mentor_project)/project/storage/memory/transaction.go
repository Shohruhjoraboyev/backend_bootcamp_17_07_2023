package memory

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"project/models"
	"time"

	"github.com/google/uuid"
)

type transactionRepo struct {
	fileName string
}

func NewTransactionRepo(fn string) *transactionRepo {
	return &transactionRepo{fileName: fn}
}

func (t *transactionRepo) CreateTransaction(req models.CreateTransaction) (string, error) {
	id := uuid.New()
	transactions, err := t.read()
	if err != nil {
		return "", err
	}
	transactions = append(transactions, models.Transaction{
		Id:          id.String(),
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

	return id.String(), nil
}

func (t *transactionRepo) UpdateTransaction(req models.Transaction) (string, error) {
	transactions, err := t.read()
	if err != nil {
		return "", err
	}
	for i, v := range transactions {
		if v.Id == req.Id {
			transactions[i] = req
			return " Transaction Updated succesfully", nil

		}
	}
	return "", errors.New("Not found ")
}

func (t *transactionRepo) GetTransaction(req models.IdRequest) (resp models.Transaction, err error) {
	transactions, err := t.read()
	if err != nil {
		return models.Transaction{}, err
	}
	for _, v := range transactions {
		if v.Id == resp.Id {
			return v, nil
		}
	}
	return models.Transaction{}, errors.New("not found")
}

func (t *transactionRepo) GetAllTransaction(req models.GetAllTransactionRequest) (resp models.GetAllTransactionResponse, err error) {
	start := req.Limit * (req.Page - 1)
	end := start + req.Limit
	transactions, err := t.read()
	if err != nil {
		return models.GetAllTransactionResponse{}, err
	}
	if start > len(transactions) {
		resp.Transactions = []models.Transaction{}
		resp.Count = len(transactions)
		return
	} else if end > len(transactions) {
		return models.GetAllTransactionResponse{
			Transactions: transactions[start:],
			Count:        len(transactions),
		}, nil
	}

	return models.GetAllTransactionResponse{
		Transactions: transactions[start:end],
		Count:        len(transactions),
	}, nil
}

func (t *transactionRepo) DeleteTransaction(req models.IdRequest) (string, error) {
	transactions, err := t.read()
	if err != nil {
		return "", err
	}
	for i, v := range transactions {
		if v.Id == req.Id {
			transactions = append(transactions[:i], transactions[i+1:]...)
			return "deleted", nil
		}
	}
	return "", errors.New("Not found ")
}

func (t *transactionRepo) read() ([]models.Transaction, error) {
	var (
		transactions []models.Transaction
	)

	data, err := os.ReadFile(t.fileName)
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

func (s *transactionRepo) write(transactions []models.Transaction) error {

	body, err := json.Marshal(transactions)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(s.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
