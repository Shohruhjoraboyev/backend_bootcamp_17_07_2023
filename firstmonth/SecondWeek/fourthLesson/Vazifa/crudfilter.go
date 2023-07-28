package main

import "fmt"

type Category struct {
	ID       int
	Name     string
	Products []Product
}

type Product struct {
	ID    int
	Name  string
	Price float64
}

type User struct {
	ID               int
	Name             string
	FavouriteProdIDs []int
}

func main() {
	categories := []Category{
		{
			ID:   1,
			Name: "Mebellar",
			Products: []Product{
				{ID: 1, Name: "Oshxona stoli", Price: 100.0},
				{ID: 2, Name: "Kreslo", Price: 50.0},
				{ID: 3, Name: "Sharlar", Price: 20.0},
			},
		},
		{
			ID:   2,
			Name: "Elektronika",
			Products: []Product{
				{ID: 4, Name: "Telefon", Price: 200.0},
				{ID: 5, Name: "Televizor", Price: 500.0},
				{ID: 6, Name: "Komp'yuter", Price: 1000.0},
			},
		},
		{
			ID:   3,
			Name: "Mobil qurilmalar",
			Products: []Product{
				{ID: 7, Name: "Smart watch", Price: 150.0},
				{ID: 8, Name: "Headphones", Price: 80.0},
			},
		},
	}

	users := []User{
		{
			ID:               1,
			Name:             "Bobur",
			FavouriteProdIDs: []int{1, 3, 4},
		},
		{
			ID:               2,
			Name:             "Ali",
			FavouriteProdIDs: []int{2, 5, 8},
		},
		{
			ID:               3,
			Name:             "Vali",
			FavouriteProdIDs: []int{1, 7, 8},
		},
	}

	for _, user := range users {
		fmt.Printf("%s ning yoqtirgan mahsulotlari:\n", user.Name)
		for _, category := range categories {
			count := 0
			for _, product := range category.Products {
				for _, prodID := range user.FavouriteProdIDs {
					if prodID == product.ID {
						count++
					}
				}
			}
			fmt.Printf("%s: %d\n", category.Name, count)
		}
		fmt.Println()
	}
}
