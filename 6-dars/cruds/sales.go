package cruds

import (
	"fmt"
	"time"
)

const (
	card = iota
	cash
)

const (
	cancel = iota
	success
)

type Sale struct {
	ID              int
	BranchID        int
	ShopAssistantID int
	CashierID       int
	Price           float64
	PaymentType     int //card, cash
	Status          int // success,cancel
	ClientName      string
	CreatedAt       string
}

type Sales struct {
	Data []Sale
}

func (s *Sales) Create(newSale Sale) string {
	newSale.ID = len(s.Data) + 1
	newSale.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	s.Data = append(s.Data, newSale)
	return "successfully created"
}

func (s *Sales) Update(changedSale Sale) error {
	for i, v := range s.Data {
		if v.ID == changedSale.ID {
			s.Data[i].BranchID = changedSale.BranchID
			s.Data[i].ShopAssistantID = changedSale.ShopAssistantID
			s.Data[i].CashierID = changedSale.CashierID
			s.Data[i].Price = changedSale.Price
			s.Data[i].PaymentType = changedSale.PaymentType
			s.Data[i].Status = changedSale.Status
			s.Data[i].ClientName = changedSale.ClientName
			return nil
		}
	}
	return fmt.Errorf("Sale with ID %d not found", changedSale.ID)
}

func (s *Sales) GetByID(id int) (Sale, error) {
	for i := range s.Data {
		if s.Data[i].ID == id {
			return s.Data[i], nil
		}
	}
	return Sale{}, fmt.Errorf("Sale with ID %d not found", id)
}

func (s *Sales) GetAll(page, limit, branchID, shopAssistantID, cashierID, paymentType, status int, clientName string, priceFrom, priceTo float64) ([]Sale, error) {
	searched := []Sale{}
	for i, v := range s.Data {
		if v.BranchID == branchID && v.ShopAssistantID == shopAssistantID && v.CashierID == cashierID && v.PaymentType == paymentType && v.Status == status && search(v.ClientName, clientName) && v.Price >= priceFrom && v.Price <= priceTo {
			searched = append(searched, s.Data[i])
		}
	}

	start := (page - 1) * limit
	end := start + limit

	if len(searched) < limit && page > 1 {
		return []Sale{}, fmt.Errorf("page not found")
	}
	if end > len(searched) {
		end = len(searched)
	}
	return searched[start:end], nil
}

func (s *Sales) Delete(id int) string {
	for i, v := range s.Data {
		if v.ID == id {
			s.Data = append(s.Data[:i], s.Data[i+1:]...)
			return "successfully deleted"
		}
	}
	return fmt.Sprintf("Sale with ID %d not found", id)
}

func CardCash(n int) string {
	if n == card {
		return "card"
	}
	if n == cash {
		return "cash"
	}
	return ""
}

func StatusName(n int) string {
	if n == success {
		return "success"
	}
	if n == cancel {
		return "cancel"
	}
	return ""
}
