package main

import (
	"errors"
	"fmt"
	"sort"
	"time"
)

type Branch struct {
	ID        int
	Name      string
	Address   string
	CreatedAt string // 2-task
	FoundedAt string // 3-task
	Year      int    // 3-task
}

func (b *Branch) CalculateYear() error {
	if b.FoundedAt == "" {
		return errors.New("FoundedAt is not set")
	}

	foundedTime, err := time.Parse("2006-01-02", b.FoundedAt)
	if err != nil {
		return err
	}

	duration := time.Since(foundedTime)
	b.Year = int(duration.Hours() / 24 / 365)
	return nil
}

type Branches struct {
	Data []Branch
}

func (b *Branches) Create(newBranch Branch) {
	newBranch.ID = len(b.Data) + 1
	newBranch.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	b.Data = append(b.Data, newBranch)
	fmt.Println("Successfully added new branch")
}

func (b *Branches) Update(updatedBranch Branch) {
	for i, branch := range b.Data {
		if branch.ID == updatedBranch.ID {
			updatedBranch.CreatedAt = branch.CreatedAt
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
			err := branch.CalculateYear()
			if err != nil {
				fmt.Println(err)
			}
			return branch, nil
		}
	}
	return Branch{}, errors.New("Branch not found")
}

func (b *Branches) GetAll(name, address string) []Branch {
	result := make([]Branch, 0)
	for _, branch := range b.Data {
		if (name == "" || branch.Name == name) && (address == "" || branch.Address == address) {
			err := branch.CalculateYear()
			if err != nil {
				fmt.Println(err)
				continue
			}
			result = append(result, branch)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].CreatedAt > result[j].CreatedAt
	})

	return result
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
		newBranch := Branch{
			Name:      "Logo city",
			Address:   "Andijan",
			FoundedAt: "2010-01-01",
		}
		branches.Create(newBranch)
	}

	// Update
	updatedBranch := Branch{
		ID:      2,
		Name:    "Marketing agency",
		Address: "Tashkent",
	}
	branches.Update(updatedBranch)

	// GetById
	branch, err := branches.GetById(3)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("ID: %d\n", branch.ID)
		fmt.Printf("Name: %s\n", branch.Name)
		fmt.Printf("Address: %s\n", branch.Address)
		fmt.Printf("CreatedAt: %s\n", branch.CreatedAt)
		fmt.Printf("FoundedAt: %s\n", branch.FoundedAt)
		fmt.Printf("Since: %d\n", branch.Year)
	}

	// GetAllda
	allBranches := branches.GetAll("Logo city", "")
	for _, branch := range allBranches {
		fmt.Printf("ID: %d\n", branch.ID)
		fmt.Printf("Name: %s\n", branch.Name)
		fmt.Printf("Address: %s\n", branch.Address)
		fmt.Printf("CreatedAt: %s\n", branch.CreatedAt)
		fmt.Printf("FoundedAt: %s\n", branch.FoundedAt)
		fmt.Printf("Since: %d\n", branch.Year)
	}

	// Delete
	branches.Delete(2)
}
