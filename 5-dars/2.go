package main

import (
	"fmt"
	"strings"
)

type Staff struct {
	ID       int
	BranchID int
	TarifID  int
	Type     string // cashier, shop_assistant
	Name     string
	Balance  float64
}

type Staffs struct {
	Data []Staff
}

func main() {
	staffs := Staffs{Data: make([]Staff, 0)}

	staffs.Create(Staff{BranchID: 12, TarifID: 2, Type: "cashier", Name: "Sardor", Balance: 12.4})
	staffs.Create(Staff{BranchID: 12, TarifID: 1, Type: "shop_assistant", Name: "Bexruz", Balance: 9.62})
	staffs.Create(Staff{BranchID: 13, TarifID: 1, Type: "cashier", Name: "Asadbek", Balance: 10.4})
	staffs.Create(Staff{BranchID: 12, TarifID: 2, Type: "shop_assistant", Name: "Nodir", Balance: 16.2})

	for _, v := range staffs.Data {
		fmt.Println(v)
	}
	fmt.Println()

	err := staffs.Update(Staff{ID: 1, BranchID: 12, TarifID: 1, Type: "cashier", Name: "Sardor", Balance: 14.4})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("data changed successfully")
		for _, v := range staffs.Data {
			fmt.Println(v)
		}
	}
	fmt.Print("\nGetByID id=2; result:\n")

	staff, err := staffs.GetByID(2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(staff)
	}

	fmt.Print("\nGetAll page=1,limit=5,brachID=12,tarifID=2,type=cashier,searchName ='a',balance from 0 to 20\n")
	stafs, err := staffs.GetAll(1, 5, 12, 1, "cashier", "a", 0, 20)
	if err != nil {
		fmt.Println(err)
	} else {
		for i := range stafs {
			fmt.Println(stafs[i])
		}
	}

	fmt.Print("\nDelete ID=2\n")
	fmt.Println(staffs.Delete(2))
	for _, v := range staffs.Data {
		fmt.Println(v)
	}
}

func (s *Staffs) Create(newStaff Staff) string {
	newStaff.ID = len(s.Data) + 1
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
			return nil
		}
	}
	return fmt.Errorf("Staff with ID %d not found", changedStaff.ID)
}

func (s *Staffs) GetByID(id int) (Staff, error) {
	for i := range s.Data {
		if s.Data[i].ID == id {
			return s.Data[i], nil
		}
	}
	return Staff{}, fmt.Errorf("Staff with ID %d not found", id)
}

var search = strings.Contains

func (s *Staffs) GetAll(page, limit, branchId, tarifId int, typ, name string, balanceFrom, balanceTo float64) ([]Staff, error) {
	searched := []Staff{}
	for i, v := range s.Data {
		if v.BranchID == branchId && v.TarifID == tarifId && v.Type == typ && search(v.Name, name) && v.Balance > balanceFrom && v.Balance < balanceTo {
			searched = append(searched, s.Data[i])
		}
	}

	start := (page - 1) * limit
	end := start + limit

	if len(searched) < limit && page > 1 {
		return []Staff{}, fmt.Errorf("Page not found")
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
