package main

import (
	"errors"
	"fmt"
	"sort"
	"time"
)

type Staff struct {
	ID        int
	BranchID  int
	TariffID  int
	Type      string
	Name      string
	Balance   float64
	CreatedAt string // 2-task
	BirthDay  string // 4-task
	Age       int    // 4-task
}

func (s *Staff) CalculateAge() error {
	if s.BirthDay == "" {
		return errors.New("Birthday is not set")
	}

	birthdayTime, err := time.Parse("2006-01-02", s.BirthDay)
	if err != nil {
		return err
	}

	duration := time.Since(birthdayTime)
	s.Age = int(duration.Hours() / 24 / 365)
	return nil
}

type Staffs struct {
	Data []Staff
}

func (s *Staffs) Create(newStaff Staff) {
	newStaff.ID = len(s.Data) + 1
	newStaff.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

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
			err := staff.CalculateAge()
			if err != nil {
				fmt.Println(err)
			}
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
			err := staff.CalculateAge()
			if err != nil {
				fmt.Println(err)
				continue
			}
			result = append(result, staff)
		}
	}
	// sort by created_at in descending order
	sort.Slice(result, func(i, j int) bool {
		return result[i].CreatedAt > result[j].CreatedAt
	})
	return result

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
	for i := 0; i < 5; i++ {
		staffs.Create(Staff{
			BranchID: 1,
			TariffID: 1,
			Type:     "cashier",
			Name:     "Omadbek",
			Balance:  100.0,
			BirthDay: "2002-01-01",
		})
	}

	// Update
	staffs.Update(Staff{
		ID:       2,
		Type:     "shop_assistant",
		Name:     "Sarvarbek",
		Balance:  200.0,
		BirthDay: "1999-04-11",
	})

	// GetById
	staff, err := staffs.GetById(3)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Staff found by ID:", staff)

		fmt.Println("ID:", staff.ID)
		fmt.Println("Name:", staff.Name)
		fmt.Println("Type:", staff.Type)
		fmt.Println("Balance:", staff.Balance)
		fmt.Println("Birthday:", staff.BirthDay)
		fmt.Println("  --- ---   ----   ----  ---")
	}

	// GetAll
	allStaffs := staffs.GetAll(1, 0, "", "Omadbek", 0, 0)
	fmt.Println("All staffs with name 'Omadbek':", " ")
	for _, staff := range allStaffs {
		fmt.Println("ID:", staff.ID)
		fmt.Println("Name:", staff.Name)
		fmt.Println("Type:", staff.Type)
		fmt.Println("Balance:", staff.Balance)
		fmt.Println("Birthday:", staff.BirthDay)
		fmt.Println("--//--//--//---//---///---///----///")
	}

	// Delete
	staffs.Delete(2)

}
