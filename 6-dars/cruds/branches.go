package cruds

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Branch struct {
	ID        int
	Name      string
	Address   string
	CreatedAt string
	FoundedAt string
	Year      int
}

type Branches struct {
	Data []Branch
}

var search = strings.Contains

func (b *Branches) Create(newBranch Branch) string {
	newBranch.ID = len(b.Data) + 1
	newBranch.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	b.Data = append(b.Data, newBranch)
	return "successfully created"
}

func (b *Branches) Update(changedBranch Branch) error {
	for i, v := range b.Data {
		if v.ID == changedBranch.ID {
			b.Data[i].Name = changedBranch.Name
			b.Data[i].Address = changedBranch.Address
			b.Data[i].FoundedAt = changedBranch.FoundedAt
			return nil
		}
	}
	return fmt.Errorf("branch with ID %d not found", changedBranch.ID)
}

func (b *Branches) GetByID(id int) (Branch, error) {
	for i := range b.Data {
		if b.Data[i].ID == id {
			foundedAt, _ := strconv.Atoi(b.Data[i].FoundedAt[:4])
			b.Data[i].Year = time.Now().Year() - foundedAt
			return b.Data[i], nil
		}
	}
	return Branch{}, fmt.Errorf("Branch with ID %d not found", id)
}

func (b *Branches) GetAll(page, limit int, name, address string) ([]Branch, error) {
	searched := []Branch{}
	for i, v := range b.Data {
		if search(v.Name, name) || search(v.Address, address) {
			foundedAt, _ := strconv.Atoi(b.Data[i].FoundedAt[:4])
			b.Data[i].Year = time.Now().Year() - foundedAt
			searched = append(searched, b.Data[i])
		}
	}

	start := (page - 1) * limit
	end := start + limit

	if len(searched) < limit && page > 1 {
		return []Branch{}, fmt.Errorf("page not found")
	}
	if end > len(searched) {
		end = len(searched)
	}

	return b.Reverse(searched[start:end]), nil
}

func (b *Branches) Delete(id int) string {
	for i, v := range b.Data {
		if v.ID == id {
			b.Data = append(b.Data[:i], b.Data[i+1:]...)
			return "successfully deleted"
		}
	}
	return fmt.Sprintf("Branch with ID %d not found", id)
}

func (b *Branches) Reverse([]Branch) []Branch {
	for i, j := 0, len(b.Data)-1; i < j; i, j = i+1, j-1 {
		b.Data[i], b.Data[j] = b.Data[j], b.Data[i]
	}
	return b.Data
}
