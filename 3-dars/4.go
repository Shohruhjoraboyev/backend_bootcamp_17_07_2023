package main

import "fmt"

type User struct {
	ID                int
	Name              string
	FavouriteProducts []int
}

type Product struct {
	ID    int
	Name  string
	Price int
}

type Category struct {
	ID       int
	Name     string
	Products []string
}

func main() {
	users := []User{
		{ID: 1, Name: "Sardor", FavouriteProducts: []int{2, 1, 4, 3, 10}},
		{ID: 2, Name: "Omadbek", FavouriteProducts: []int{1, 2, 6, 8, 11}},
		{ID: 3, Name: "Nodir", FavouriteProducts: []int{2, 3, 4, 9, 12}},
		{ID: 4, Name: "Samandar", FavouriteProducts: []int{10, 8, 12, 1, 5}},
		{ID: 5, Name: "Sherzod", FavouriteProducts: []int{4, 5, 6, 9, 11, 3}},
	}

	products := []Product{
		{ID: 1, Name: "olma", Price: 10},
		{ID: 2, Name: "shaftoli", Price: 19},
		{ID: 3, Name: "uzum", Price: 8},
		{ID: 4, Name: "shashlik", Price: 45},
		{ID: 5, Name: "manti", Price: 25},
		{ID: 6, Name: "chuchvara", Price: 23},
		{ID: 7, Name: "lavash", Price: 22},
		{ID: 8, Name: "burger", Price: 16},
		{ID: 9, Name: "sandvich", Price: 18},
		{ID: 10, Name: "coca-cola", Price: 12},
		{ID: 11, Name: "fanta", Price: 11},
		{ID: 12, Name: "pepsi", Price: 11},
	}

	categories := []Category{
		{ID: 1, Name: "meva", Products: []string{"olma", "shaftoli", "uzum"}},
		{ID: 2, Name: "ovqat", Products: []string{"shashlik", "manti", "chuchvara"}},
		{ID: 3, Name: "fastfood", Products: []string{"lavash", "burger", "sandvich"}},
		{ID: 4, Name: "ichimlik", Products: []string{"coca-cola", "fanta", "pepsi"}},
	}

	catUser := make(map[string]map[string]int)
	for _, user := range users {
		cats := make(map[string]int)
		for _, cat := range categories {
			cats[cat.Name] = 0
		}
		catUser[user.Name] = cats
	}

	for _, user := range users {

		for _, fp := range user.FavouriteProducts {

			t_f := findCategory(fp, categories, products)
			catUser[user.Name][t_f] += 1
		}
	}

	for key := range catUser {
		fmt.Printf("%s: ", key)
		count := 0
		for cat, catNum := range catUser[key] {
			fmt.Printf("%s:%d", cat, catNum)
			count++
			if count < len(catUser[key]) {
				fmt.Printf(", ")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}

func findCategory(n int, categories []Category, products []Product) string {
	for _, product := range products {
		if product.ID == n {
			for _, cat := range categories {
				if contains(cat.Products, product.Name) {
					return cat.Name
				}
			}
		}
	}
	return "False"
}

func contains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
