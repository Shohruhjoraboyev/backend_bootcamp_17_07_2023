package main

import (
	"fmt"
	"strings"
)

func main() {

	staf1 := Staff{
		Id:       1,
		BranchId: 1,
		TarifId:  1,
		TypeId:   1,
		Name:     "Zo'r zo'r",
		Balance:  123.3,
	}
	staf2 := Staff{
		Id:       2,
		BranchId: 2,
		TarifId:  2,
		TypeId:   2,
		Name:     "Zo'r UPDATED",
		Balance:  223.3,
	}

	// CREATE STAFF
	res, err := staffes.CreateStaff(staf1)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(res)
	}
	fmt.Println(staffes)

	// GET STAFF BY ID
	staff, err := staffes.GetStaffById(1)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(staff)
	}

	// GET ALL DATA
	// (page, limit, branchId, tariffId, TypeId int, name string, balanceFrom, balanceTo float64)
	data, err := staffes.getAllStaffes(1, 1, 1, 1, 1, "zo'r", 12.2, 1221.2)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(data)
	}

	// update staff
	result, err := staffes.UpdateStaff(1, staf2)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}
	fmt.Println(staffes)

	// DELETE STAFF
	result, err = staffes.Deletestaff(1)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}
	fmt.Println(staffes)
}

const (
	Cashier = iota + 1
	Shop_assistant
)

var staffes = Staffes{Data: make([]Staff, 0)}

type Staff struct {
	Id       int
	BranchId int
	TarifId  int
	TypeId   int
	Name     string
	Balance  float64
}

type Staffes struct {
	Data []Staff
}

func (s *Staffes) CreateStaff(newStaff Staff) (string, error) {
	newStaff.Id = len(s.Data) + 1
	for _, staff := range s.Data {
		if staff.Id == newStaff.Id {
			return "", fmt.Errorf("branch with ID %d already exits", newStaff.Id)
		}
	}

	s.Data = append(s.Data, newStaff)
	return "created", nil

}

func (s *Staffes) GetStaffById(id int) (Staff, error) {
	for _, s := range staffes.Data {
		if s.Id == id {
			return s, nil
		}
	}
	return Staff{}, fmt.Errorf("no branch found with ID %d", id)
}

func (s *Staffes) getAllStaffes(page, limit, branchId, tariffId, TypeId int, name string, balanceFrom, balanceTo float64) ([]Staff, error) {
	filtered := []Staff{}

	for i, v := range s.Data {
		if v.BranchId == branchId && v.TarifId == TypeId && v.TypeId == TypeId && strings.Contains(v.Name, name) && v.Balance > balanceFrom && v.Balance < balanceTo {
			filtered = append(filtered, s.Data[i])
		}
	}

	start := (page - 1) * limit
	end := start + limit

	if len(filtered) < limit && page > 1 {
		return []Staff{}, fmt.Errorf("page not found")
	}
	if end > len(filtered) {
		end = len(filtered)
	}

	return filtered[start:end], nil
}

func (s *Staffes) UpdateStaff(id int, updatedStaff Staff) (string, error) {
	index := -1
	for i, staff := range s.Data {
		if staff.Id == id {
			index = i
			break
		}
	}

	if index == -1 {
		return "", fmt.Errorf("no staff odun with ID %d", id)
	}

	// Check if the updated branch ID already exists
	for i, staff := range s.Data {
		if staff.Id == updatedStaff.Id && i != id {
			return "", fmt.Errorf("branch with ID %d already exits", updatedStaff.Id)
		}
	}

	updatedStaff.Id = id
	s.Data[index] = updatedStaff
	return "updated", nil
}

func (s *Staffes) Deletestaff(id int) (string, error) {
	index := -1
	for i, staff := range s.Data {
		if staff.Id == id {
			index = i
			break
		}
	}
	if index == -1 {
		return "", fmt.Errorf("no staff found with ID %d", id)
	}

	s.Data = append(s.Data[:index], s.Data[index+1:]...)

	return fmt.Sprintf("staff with ID %d deleted", id), nil
}
