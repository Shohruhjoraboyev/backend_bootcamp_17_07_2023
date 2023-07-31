package cruds

import (
	"fmt"
	"strconv"
	"time"
)

type Staff struct {
	ID        int
	BranchID  int
	TarifID   int
	Type      string // cashier, shop_assistant
	Name      string
	Balance   float64
	CreatedAt string
	BirthDay  string
	Age       int
}

type Staffs struct {
	Data []Staff
}

func (s *Staffs) Create(newStaff Staff) string {
	newStaff.ID = len(s.Data) + 1
	newStaff.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	s.Data = append(s.Data, newStaff)
	return "successfully created"
}

func (s *Staffs) Update(changedStaff Staff) error {
	for i, v := range s.Data {
		if v.ID == changedStaff.ID {
			s.Data[i].BranchID = changedStaff.BranchID
			s.Data[i].TarifID = changedStaff.TarifID
			s.Data[i].Type = changedStaff.Type
			s.Data[i].Name = changedStaff.Name
			s.Data[i].Balance = changedStaff.Balance
			s.Data[i].BirthDay = changedStaff.BirthDay
			return nil
		}
	}
	return fmt.Errorf("Staff with ID %d not found", changedStaff.ID)
}

func (s *Staffs) GetByID(id int) (Staff, error) {
	for i := range s.Data {
		if s.Data[i].ID == id {
			brYear, _ := strconv.Atoi(s.Data[i].BirthDay[:4])
			s.Data[i].Age = time.Now().Year() - brYear
			return s.Data[i], nil
		}
	}
	return Staff{}, fmt.Errorf("Staff with ID %d not found", id)
}

func (s *Staffs) GetAll(page, limit, branchId, tarifId int, typ, name string, balanceFrom, balanceTo float64) ([]Staff, error) {
	searched := []Staff{}
	for i, v := range s.Data {
		if v.BranchID == branchId && v.TarifID == tarifId && v.Type == typ && search(v.Name, name) && v.Balance > balanceFrom && v.Balance < balanceTo {
			brYear, _ := strconv.Atoi(s.Data[i].BirthDay[:4])
			s.Data[i].Age = time.Now().Year() - brYear
			searched = append(searched, s.Data[i])
		}
	}

	start := (page - 1) * limit
	end := start + limit

	if len(searched) < limit && page > 1 {
		return []Staff{}, fmt.Errorf("page not found")
	}
	if end > len(searched) {
		end = len(searched)
	}

	return searched[start:end], nil
}

func (s *Staffs) Delete(id int) string {
	for i, v := range s.Data {
		if v.ID == id {
			s.Data = append(s.Data[:i], s.Data[i+1:]...)
			return "successfully deleted"
		}
	}
	return fmt.Sprintf("Staff with ID %d not found", id)
}
