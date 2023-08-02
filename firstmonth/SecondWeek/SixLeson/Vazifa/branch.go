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
	CreatedAt string //2-task
	FoundedAt string //3-task
	Year      int    //3-task
}

type Branches struct {
	Data []Branch
}

// 3-task
// method of to calculate years
func (b *Branch) CalculateYear() error {

	foundedTime, err := time.Parse("2006-01-02", b.FoundedAt)
	if err != nil {
		return err
	}
	duration := time.Since(foundedTime)
	b.Year = int(duration.Hours() / 24 / 365)
	return nil
}
func (b *Branches) Create(newBranch Branch) {
	newBranch.ID = len(b.Data) + 1
	newBranch.CreatedAt = time.Now().Format("2006-01-02 15:04:05") //1-task
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
			err := branch.CalculateYear() // calculate year
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
			err := branch.CalculateYear() // calculate year
			if err != nil {
				fmt.Println(err)
				continue
			}
			result = append(result, branch)
		}
	}
	//2-task
	// sort by created_at in descending order

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
func (b Branch) String() string {
	return fmt.Sprintf("{%d Name: %s Address: %s The time of now: %s Founded: %s Since: %d}", b.ID, b.Name, b.Address, b.CreatedAt, b.FoundedAt, b.Year)
}
func main() {
	branches := Branches{Data: make([]Branch, 0)}

	// Create
	for i := 0; i < 3; i++ {
		branches.Create(Branch{
			Name:      "Logo city ",
			Address:   "Andijan ",
			FoundedAt: "2002-01-01",
		})
	}

	// Update
	branches.Update(Branch{
		ID:        2,
		Name:      "Marketing agency",
		Address:   "Tashkent",
		FoundedAt: "2002-02-01",
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
	fmt.Println("All branches with name 'Logo city':", allBranches)

	// Delete
	branches.Delete(2)

	// GetAll after delete
	allBranches = branches.GetAll("", "")
	fmt.Println("All branches after delete:", allBranches)
}
