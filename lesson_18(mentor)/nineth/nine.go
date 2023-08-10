package nineth

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"task/models"
)

// 9. Filialda qancha summalik product borligi jadvali:
// 1. Branch1        853 000
// 2. Branch2      1 982 000

// branch skladni range qilib branchId, productId va quantity ovolishim kerak
// keyin quantityni productid orqali pricega kopaytirish kerak
type BranchProduct struct {
	BranchID  int
	ProductID int
	Quantity  int
}

type Product struct {
	ID         int
	Name       string
	Price      int
	CategoryID int
}

type Branch struct {
	ID                int
	Name              string
	ProductQuantities map[int]int
}

func CalculateProductSum() {
	branch_products, _ := readBranches("data/branch_products.json")
	productes, _ := readProducts("data/products.json")
	branches, _ := readBranch("data/branches.json")

	branchMap := make(map[int]Branch)

	for _, branchProduct := range branch_products {
		branchID := branchProduct.BranchId
		productID := branchProduct.ProductId
		quantity := branchProduct.Quantity

		branch, exists := branchMap[branchID]
		if !exists {
			branch = Branch{
				ID:                branchID,
				ProductQuantities: make(map[int]int),
			}
		}

		branch.ProductQuantities[productID] += quantity
		branchMap[branchID] = branch
	}
	container := make(map[int]int)
	for branchID, v := range branchMap {
		for _, p := range productes {
			for product_id, quantity := range v.ProductQuantities {
				if p.Id == product_id {
					container[branchID] += quantity * p.Price
				}
			}
		}
	}

	sortedContainer := []models.ModelFor9{}
	for i, v := range container {
		for _, b := range branches {
			if i == b.ID {
				sortedContainer = append(sortedContainer, models.ModelFor9{
					Name: b.Name,
					Sum:  v,
				})
			}
		}
	}

	sort.Slice(sortedContainer, func(i, j int) bool {
		return sortedContainer[i].Sum > sortedContainer[j].Sum
	})

	for _, v := range sortedContainer {
		fmt.Printf("%s: sum: %d\n", v.Name, v.Sum)
	}
}

// ================================READERS======================================

func readProducts(data string) ([]models.Products, error) {
	var products []models.Products

	p, err := os.ReadFile(data)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}
	err = json.Unmarshal(p, &products)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}
	return products, nil
}

func readBranches(data string) ([]models.Product7task, error) {
	var branches []models.Product7task

	branch, err := os.ReadFile(data)
	if err != nil {
		log.Printf("Error while Read branch: %+v", err)
		return nil, err
	}
	err = json.Unmarshal(branch, &branches)
	if err != nil {
		log.Printf("Error while Unmarshal branch: %+v", err)
		return nil, err
	}
	return branches, nil
}
func readBranch(data string) ([]models.Branch, error) {
	var branches []models.Branch

	branch, err := os.ReadFile(data)
	if err != nil {
		log.Printf("Error while Read branch: %+v", err)
		return nil, err
	}
	err = json.Unmarshal(branch, &branches)
	if err != nil {
		log.Printf("Error while Unmarshal branch: %+v", err)
		return nil, err
	}
	return branches, nil
}
