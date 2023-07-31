package cruds

import (
	"fmt"
	"time"
)

const (
	Fixed = iota + 1
	Percent
)

type StaffTariff struct {
	ID            int
	Name          string
	Type          int // fixed, percent
	AmountForCash float64
	AmountForCard float64
	CreatedAt     string
	FoundedAt     string
}

type StaffTariffs struct {
	Data []StaffTariff
}

func (s *StaffTariffs) Create(newTariff StaffTariff) string {
	newTariff.ID = len(s.Data) + 1
	newTariff.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	s.Data = append(s.Data, newTariff)
	return "created successfully"
}

func (s *StaffTariffs) Update(changedTariff StaffTariff) error {
	for i, v := range s.Data {
		if v.ID == changedTariff.ID {
			s.Data[i].Name = changedTariff.Name
			s.Data[i].Type = changedTariff.Type
			s.Data[i].AmountForCash = changedTariff.AmountForCash
			s.Data[i].AmountForCard = changedTariff.AmountForCard
			s.Data[i].FoundedAt = changedTariff.FoundedAt
			return nil
		}
	}
	return fmt.Errorf("tariff with ID %d not found", changedTariff.ID)
}

func (s *StaffTariffs) GetByID(id int) (StaffTariff, error) {
	for i := range s.Data {
		if s.Data[i].ID == id {
			return s.Data[i], nil
		}
	}
	return StaffTariff{}, fmt.Errorf("tariff with ID %d not found", id)
}

func (s *StaffTariffs) GetAll(page, limit, typ int, name string) ([]StaffTariff, error) {
	searched := []StaffTariff{}
	for i, v := range s.Data {
		if search(v.Name, name) && v.Type == typ {
			searched = append(searched, s.Data[i])
		}
	}
	start := (page - 1) * limit
	end := start + limit

	if len(searched) < limit && page > 1 {
		return []StaffTariff{}, fmt.Errorf("page not found")
	}
	if end > len(searched) {
		end = len(searched)
	}
	return searched[start:end], nil
}

func (s *StaffTariffs) Delete(id int) string {
	for i, v := range s.Data {
		if v.ID == id {
			s.Data = append(s.Data[:i], s.Data[i+1:]...)
			return "successfully deleted"
		}
	}
	return fmt.Sprintf("tariff with ID %d not found", id)
}
