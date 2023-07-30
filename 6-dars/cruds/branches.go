package cruds

import (
	"fmt"
	"strings"
	"time"
)

type Branch struct {
	ID        int
	Name      string
	Address   string
	CreatetAt string
}

type Branches struct {
	Data []Branch
}

var search = strings.Contains

// func main() {
// 	branches := Branches{Data: make([]Branch, 0)}
// 	branches.Create(Branch{Name: "first", Address: "Yunusobod"})
// 	branches.Create(Branch{Name: "second", Address: "Olmazor"})
// 	branches.Create(Branch{Name: "third", Address: "Chilonzor"})

// 	fmt.Println(branches.Update(Branch{ID: 1, Name: "thirddd", Address: "olmazor"}))

// 	fmt.Println(branches)

// 	br, err := branches.GetByID(2)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(br)
// 	}

// 	res, err := branches.GetAll(1, 3, "i", "zor")
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(res)
// 	}

// 	fmt.Println(branches.Delete(3))
// 	fmt.Println(branches)
// }

func (b *Branches) Create(newBranch Branch) string {
	newBranch.ID = len(b.Data) + 1
	newBranch.CreatetAt = time.Now().Format("2006-01-02 15:04:05")
	b.Data = append(b.Data, newBranch)
	return "successfully created"
}

func (b *Branches) Update(changedBranch Branch) error {
	for i, v := range b.Data {
		if v.ID == changedBranch.ID {
			b.Data[i].Name = changedBranch.Name
			b.Data[i].Address = changedBranch.Address
			return nil
		}
	}
	return fmt.Errorf("Branch with ID %d not found", changedBranch.ID)
}

func (b *Branches) GetByID(id int) (Branch, error) {
	for i := range b.Data {
		if b.Data[i].ID == id {
			return b.Data[i], nil
		}
	}
	return Branch{}, fmt.Errorf("Branch with ID %d not found", id)
}

func (b *Branches) GetAll(page, limit int, name, address string) ([]Branch, error) {
	searched := []Branch{}
	for i, v := range b.Data {
		if search(v.Name, name) || search(v.Address, address) {
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
	return searched[start:end], nil
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
