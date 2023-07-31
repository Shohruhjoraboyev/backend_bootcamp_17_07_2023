package main

import (
	"dars_6/cruds"
	"fmt"
	"sort"
)

func main() {
	branches := cruds.Branches{Data: make([]cruds.Branch, 0)}

	branches.Create(cruds.Branch{Name: "first", Address: "Yunusobod", FoundedAt: "2003-09-12"})
	branches.Create(cruds.Branch{Name: "second", Address: "Olmazor", FoundedAt: "2001-08-21"})
	branches.Create(cruds.Branch{Name: "third", Address: "Chilonzor", FoundedAt: "2007-01-26"})

	// res, err := branches.GetAll(1, 10, "i", "o")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	for _, v := range res {
	// 		fmt.Println(v)
	// 	}
	// }

	// res, err := branches.GetByID(2)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(res)
	// }

	staffTariffs := cruds.StaffTariffs{Data: make([]cruds.StaffTariff, 0)}
	staffTariffs.Create(cruds.StaffTariff{ID: 1, Name: "percent-20", Type: 0, FoundedAt: "2002-10-12"})

	fmt.Println(staffTariffs.GetByID(1))

	products := []cruds.Product{}
	products = append(products, cruds.Product{ID: 1, Name: "coca-cola", Sizes: []cruds.Size{{ID: 1, Name: "0.5L", Price: 6}, {ID: 2, Name: "1L", Price: 9}, {ID: 3, Name: "1.5L", Price: 12}}})
	products = append(products, cruds.Product{ID: 2, Name: "dena", Sizes: []cruds.Size{{ID: 1, Name: "1L", Price: 13}}})

	for _, v := range products {
		fmt.Println(v)
	}

	clients := []cruds.Client{}
	clients = append(clients, cruds.Client{ID: 1, Name: "Diyorbek", Card: cruds.Card{Products: []cruds.CardProducts{{ProductID: 1, SizeID: 3, Quantity: 2}, {ProductID: 2, SizeID: 1, Quantity: 2}}}})
	clients = append(clients, cruds.Client{ID: 2, Name: "Sardor", Card: cruds.Card{Products: []cruds.CardProducts{{ProductID: 1, SizeID: 2, Quantity: 1}, {ProductID: 2, SizeID: 1, Quantity: 2}}}})

	fmt.Printf("\n--------------------------------------------\n")
	card := make(map[string]int)
	var sum float64
	for _, client := range clients {
		for _, product := range client.Card.Products {
			for _, pr := range products {
				if product.ProductID == pr.ID {
					card[pr.Name] += product.Quantity
					for _, size := range pr.Sizes {
						if product.SizeID == size.ID {
							sum += size.Price * float64(product.Quantity)
						}
					}
				}
			}
		}
		fmt.Printf("%s - %.2f\n", client.Name, sum)
		sum = 0
	}

	fmt.Printf("\n--------------------------------------------\n")

	result := make([]string, 0, len(card))
	for key := range card {
		result = append(result, key)
	}

	sort.SliceStable(result, func(i, j int) bool {
		return card[result[i]] > card[result[j]]
	})

	for _, v := range result {
		fmt.Printf("%s - %d\n", v, card[v])
	}
}
