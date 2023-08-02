package main

import (
	"errors"
	"fmt"
	"sort"
	"time"
)

type StaffTariff struct {
	ID            int
	Name          string
	Type          string
	AmountForCash float64
	AmountForCard float64
	CreatedAt     string //2-task
}

type StaffTariffs struct {
	Data []StaffTariff
}

func (st *StaffTariffs) Create(newStaffTariff StaffTariff) {
	newStaffTariff.ID = len(st.Data) + 1
	newStaffTariff.CreatedAt = time.Now().Format("2006-01-02 15:04:05") //2-task

	st.Data = append(st.Data, newStaffTariff)
	fmt.Println("Successfully added new staff tariff")
}

func (st *StaffTariffs) Update(updatedStaffTariff StaffTariff) {
	for i, staffTariff := range st.Data {
		if staffTariff.ID == updatedStaffTariff.ID {
			updatedStaffTariff.CreatedAt = staffTariff.CreatedAt
			st.Data[i] = updatedStaffTariff
			fmt.Println("Successfully updated staff tariff")
			return
		}
	}
	fmt.Println("Staff tariff not found")
}

func (st *StaffTariffs) GetById(id int) (StaffTariff, error) {
	for _, staffTariff := range st.Data {
		if staffTariff.ID == id {
			return staffTariff, nil
		}
	}
	return StaffTariff{}, errors.New("Staff tariff not found")
}

func (st *StaffTariffs) GetAll(name, tariffType string) []StaffTariff {
	result := make([]StaffTariff, 0)
	for _, staffTariff := range st.Data {
		if (name == "" || staffTariff.Name == name) &&
			(tariffType == "" || staffTariff.Type == tariffType) {
			result = append(result, staffTariff)
		}
	}
	//2-task
	// sort by created_at in descending order
	sort.Slice(result, func(i, j int) bool {
		return result[i].CreatedAt > result[j].CreatedAt
	})

	return result
}

func (st *StaffTariffs) Delete(id int) {
	for i, staffTariff := range st.Data {
		if staffTariff.ID == id {
			st.Data = append(st.Data[:i], st.Data[i+1:]...)
			fmt.Println("Successfully deleted staff tariff")
			return
		}
	}
	fmt.Println("Staff tariff not found")
}

func main() {
	staffTariffs := StaffTariffs{Data: make([]StaffTariff, 0)}

	// Create
	for i := 0; i < 3; i++ {
		staffTariffs.Create(StaffTariff{
			Name:          "Basic",
			Type:          "fixed",
			AmountForCash: 100.0,
			AmountForCard: 110.0,
		})
	}

	// Update
	staffTariffs.Update(StaffTariff{
		ID:            2,
		Name:          "Premium",
		Type:          "percent",
		AmountForCash: 10.0,
		AmountForCard: 5.0,
	})

	// GetById
	staffTariff, err := staffTariffs.GetById(3)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Staff tariff found:", staffTariff)
	}

	// GetAll
	allStaffTariffs := staffTariffs.GetAll("Basic", "fixed")
	fmt.Println("All staff tariffs with name 'Basic':", allStaffTariffs)

	// Delete
	staffTariffs.Delete(2)
}
