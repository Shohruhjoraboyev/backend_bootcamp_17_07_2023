package main

import (
	"errors"
	"fmt"
)

type Branch struct {
	ID      int
	Name    string
	Address string
}

type Branches struct {
	Data []Branch
}

func (b *Branches) Create(newBranch Branch) {
	newBranch.ID = len(b.Data) + 1
	b.Data = append(b.Data, newBranch)
	fmt.Println("Successfully added new branch")
}

func (b *Branches) Update(updatedBranch Branch) {
	for i, branch := range b.Data {
		if branch.ID == updatedBranch.ID {
			b.Data[i] = updatedBranch
			fmt.Println("Successfully updated branch")
			return
		}
	}
	fmt.Println("Branch not found")
}

func (b *Branches) GetById(id int) (Branch, error) {
	for _, branch := range b.Data {
		if branch.ID == id {
			return branch, nil
		}
	}
	return Branch{}, errors.New("Branch not found")
}

func (b *Branches) GetAll(name, address string) []Branch {
	result := make([]Branch, 0)
	for _, branch := range b.Data {
		if (name == "" || branch.Name == name) && (address == "" || branch.Address == address) {
			result = append(result, branch)
		}
	}
	return []Branch{}

}

func (b *Branches) Delete(id int) {
	for i, branch := range b.Data {
		if branch.ID == id {
			b.Data = append(b.Data[:i], b.Data[i+1:]...)
			fmt.Println("Successfully deleted branch")
			return
		}
	}
	fmt.Println("Branch not found")
}

func main() {
	branches := Branches{Data: make([]Branch, 0)}

	// Create
	for i := 0; i < 3; i++ {
		branches.Create(Branch{
			Name:    "Logo city ",
			Address: "Andijan ",
		})
	}

	// Update
	branches.Update(Branch{
		ID:      2,
		Name:    "Marketing agency",
		Address: "Tashkent",
	})

	// GetById
	branch, err := branches.GetById(3)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Branch found:", branch)
	}

	// GetAll
	allBranches := branches.GetAll("Logo city", "")
	fmt.Println("All branches with name 'Branch 1':", allBranches)

	// Delete
	branches.Delete(2)

}
