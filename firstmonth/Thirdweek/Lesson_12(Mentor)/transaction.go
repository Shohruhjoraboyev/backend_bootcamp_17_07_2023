package main

import (
	"errors"
	"fmt"
	"time"
)

type StaffTransaction struct {
	Id          int
	Sale_Id     int
	Staff_Id    int
	Type        string
	Source_Type string
	Amount      float64
	Text        string
	CreatedAt   string
}

type StaffTransactions struct {
	Data []StaffTransaction
}

func (s *StaffTransactions) Create(newTransaction StaffTransaction) {
	newTransaction.Id = len(s.Data) + 1
	newTransaction.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	s.Data = append(s.Data, newTransaction)
	fmt.Println("Successfully added new transaction")
}

func (s *StaffTransactions) Update(updatedTransaction StaffTransaction) {
	for i, transaction := range s.Data {
		if transaction.Id == updatedTransaction.Id {
			updatedTransaction.CreatedAt = transaction.CreatedAt
			s.Data[i] = updatedTransaction
			fmt.Println("Successfully updated transaction")
			return
		}
	}
	fmt.Println("Transaction not found")
}

func (s *StaffTransactions) GetById(id int) (StaffTransaction, error) {
	for _, transaction := range s.Data {
		if transaction.Id == id {
			return transaction, nil
		}
	}
	return StaffTransaction{}, errors.New("Transaction not found")
}

func (s *StaffTransactions) GetAll(sale_Id, staff_Id int, transaction_Type, source_Type string, created_From, created_To *time.Time) []StaffTransaction {
	result := make([]StaffTransaction, 0)

	for _, transaction := range s.Data {
		if (sale_Id == 0 || transaction.Sale_Id == sale_Id) &&
			(staff_Id == 0 || transaction.Staff_Id == staff_Id) &&
			(transaction_Type == "" || transaction.Type == transaction_Type) &&
			(source_Type == "" || transaction.Source_Type == source_Type) &&
			(created_From == nil || transaction.CreatedAt >= created_From.Format("2006-01-02 15:04:05")) &&
			(created_To == nil || transaction.CreatedAt <= created_To.Format("2006-01-02 15:04:05")) {
			result = append(result, transaction)
		}
	}
	return result
}

func (s *StaffTransactions) Delete(id int) {
	for i, transaction := range s.Data {
		if transaction.Id == id {
			s.Data = append(s.Data[:i], s.Data[i+1:]...)
			fmt.Println("Successfully Deleted transaction")
			return
		}
	}
	fmt.Println("Transaction not found")
}

func main() {
	transactions := StaffTransactions{Data: make([]StaffTransaction, 0)}

	//Create
	for i := 0; i < 5; i++ {
		transactions.Create(StaffTransaction{
			Id:          1,
			Sale_Id:     1,
			Staff_Id:    2,
			Type:        "withdraw",
			Source_Type: "sales",
			Amount:      100.0,
			Text:        "Withdraw from sales",
		})
	}

	//Update
	transactions.Update(StaffTransaction{
		Id:          1,
		Sale_Id:     2,
		Staff_Id:    1,
		Type:        "topup",
		Source_Type: "bonus",
		Amount:      50.0,
		Text:        "Top up from bonus",
	})

	//GetById
	transaction, err := transactions.GetById(3)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Transaction found: ", transaction)
	}

	//GetAll
	from, _ := time.Parse("2006-01-02", "2022-01-01")
	to, _ := time.Parse("2006-01-02", "2022-12-31")
	allTransactions := transactions.GetAll(1, 0, "", "", &from, &to)
	fmt.Println("All transactions in 2022: ", allTransactions)

	//Delete
	transactions.Delete(3)
}
