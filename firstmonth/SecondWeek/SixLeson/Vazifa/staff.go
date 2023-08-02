package main

import (
	"errors"
	"fmt"
	"sort"
	"time"
)

type Staff struct {
	ID           int
	BranchID     int
	TariffID     int
	Type         string
	Name         string
	Balance      float64
	CreatedAt    string //2-task
	BirthdayDate string // 4-task
	Age          int    //4-task
}

type Staffs struct {
	Data []Staff
}

// method of to calculate years
func (b *Staff) CalculateAge() error {

	birthdayTime, err := time.Parse("2006-01-02", b.BirthdayDate)
	if err != nil {
		return err
	}
	duration := time.Since(birthdayTime)
	b.Age = int(duration.Hours() / 24 / 365)
	return nil
}
func (s *Staffs) Create(newStaff Staff) {
	newStaff.ID = len(s.Data) + 1
	newStaff.CreatedAt = time.Now().Format("2006-01-02 15:04:05") //1-task
	s.Data = append(s.Data, newStaff)
	fmt.Println("Successfully added new staff")
}

func (s *Staffs) Update(updatedStaff Staff) {
	for i, staff := range s.Data {
		if staff.ID == updatedStaff.ID {
			updatedStaff.CreatedAt = staff.CreatedAt
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
			err := staff.CalculateAge() // calculate year
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
			err := staff.CalculateAge() // calculate year
			if err != nil {
				fmt.Println(err)
				continue
			}
			result = append(result, staff)
		}
	}

	//2-task
	//Sorting by descending order
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
func (staff Staff) String() string {
	return fmt.Sprintf("ID: %d,\nType: %q,\nName: %q,\nBalance: %.1f,\nBirthdayDate: %q\n", staff.ID, staff.Type, staff.Name, staff.Balance, staff.BirthdayDate)
}

func main() {
	staffs := Staffs{Data: make([]Staff, 0)}

	// Create
	for i := 0; i < 5; i++ {
		staffs.Create(Staff{
			BranchID:     1,
			TariffID:     1,
			Type:         "cashier",
			Name:         "Omadbek",
			Balance:      100.0,
			BirthdayDate: "2002-01-04",
		})
	}

	// Update
	staffs.Update(Staff{
		ID:           2,
		Type:         "shop_assistant",
		Name:         "Sarvarbek",
		Balance:      200.0,
		BirthdayDate: "1999-11-04",
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
