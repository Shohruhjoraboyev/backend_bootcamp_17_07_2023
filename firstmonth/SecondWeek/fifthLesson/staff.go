package main

import (
	"errors"
	"fmt"
)

type Staff struct {
	ID       int
	BranchID int
	TariffID int
	Type     string
	Name     string
	Balance  float64
}

type Staffs struct {
	Data []Staff
}

func (s *Staffs) Create(newStaff Staff) {
	newStaff.ID = len(s.Data) + 1
	s.Data = append(s.Data, newStaff)
	fmt.Println("Successfully added new staff")
}

func (s *Staffs) Update(updatedStaff Staff) {
	for i, staff := range s.Data {
		if staff.ID == updatedStaff.ID {
			s.Data[i] = updatedStaff
			fmt.Println("Successfully updated staff")
			return
		}
	}
	fmt.Println("Staff not found")
}

func (s *Staffs) GetById(id int) (Staff, error) {
	for _, staff := range s.Data {
		if staff.ID == id {
			return staff, nil
		}
	}
	return Staff{}, errors.New("Staff not found")
}

func (s *Staffs) GetAll(branchID, tariffID int, staffType, name string, balanceFrom, balanceTo float64) []Staff {
	result := make([]Staff, 0)
	for _, staff := range s.Data {
		if (branchID == 0 || staff.BranchID == branchID) &&
			(tariffID == 0 || staff.TariffID == tariffID) &&
			(staffType == "" || staff.Type == staffType) &&
			(name == "" || staff.Name == name) &&
			(balanceFrom == 0 || staff.Balance >= balanceFrom) &&
			(balanceTo == 0 || staff.Balance <= balanceTo) {
			result = append(result, staff)
		}
	}
	return []Staff{}

}

func (s *Staffs) Delete(id int) {
	for i, staff := range s.Data {
		if staff.ID == id {
			s.Data = append(s.Data[:i], s.Data[i+1:]...)
			fmt.Println("Successfully deleted staff")
			return
		}
	}
	fmt.Println("Staff not found")
}

func main() {
	staffs := Staffs{Data: make([]Staff, 0)}

	// Create
	for i := 0; i < 3; i++ {
		staffs.Create(Staff{
			BranchID: 1,
			TariffID: 1,
			Type:     "cashier",
			Name:     "Omadbek",
			Balance:  100.0,
		})
	}

	// Update
	staffs.Update(Staff{
		ID:      2,
		Type:    "shop_assistant",
		Name:    "Sarvarbek",
		Balance: 200.0,
	})

	// GetById
	staff, err := staffs.GetById(3)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Staff found:", staff)
	}

	// GetAll
	allStaffs := staffs.GetAll(1, 0, "", "Jonibek", 0, 0)
	fmt.Println("All staffs with name 'Jonibek':", allStaffs)

	// Delete
	staffs.Delete(2)

}
