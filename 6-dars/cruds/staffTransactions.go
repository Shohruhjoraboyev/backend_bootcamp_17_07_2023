package cruds

import (
	"fmt"
	"time"
)

type StaffTransaction struct {
	ID         int
	SaleID     int
	StaffID    int
	Type       string // withdraw, topup
	SourceType string // sales, bonus
	Amount     int
	Text       string
	CreatedAt  string
}

type StaffTransactions struct {
	Data []StaffTransaction
}

func (s *StaffTransactions) Create(newST StaffTransaction) string {
	newST.ID = len(s.Data) + 1
	newST.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	s.Data = append(s.Data, newST)
	return "successfully created"
}

func (s *StaffTransactions) Update(changedST StaffTransaction) error {
	for i, v := range s.Data {
		if v.ID == changedST.ID {
			s.Data[i].SaleID = changedST.SaleID
			s.Data[i].StaffID = changedST.StaffID
			s.Data[i].Type = changedST.Type
			s.Data[i].SourceType = changedST.SourceType
			s.Data[i].Amount = changedST.Amount
			s.Data[i].Text = changedST.Text
			return nil
		}
	}
	return fmt.Errorf("transaction with ID %d not found", changedST.ID)
}

func (s *StaffTransactions) GetByID(id int) (StaffTransaction, error) {
	for i := range s.Data {
		if s.Data[i].ID == id {
			return s.Data[i], nil
		}
	}
	return StaffTransaction{}, fmt.Errorf("transaction with ID %d not found", id)
}

// func (s *StaffTransactions) GetAll(page, limit, saleID, staffID int, typ, sourceType string)
