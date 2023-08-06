package main

import (
	"fmt"
	"lesson_15/config"
	"lesson_15/handler"
	"lesson_15/storage/memory"
)

func main() {

	cfg := config.Load()
	strg := memory.NewStorage()
	h := handler.NewHandler(strg, *cfg)

	fmt.Println("Welcome to my Golang Project!")
	fmt.Println("Available methods:")
	for _, m := range cfg.Methods {
		fmt.Println("- ", m)
	}
	fmt.Println("Available objects:")
	for _, o := range cfg.Objects {
		fmt.Println("- ", o)
	}

	for {
		fmt.Print("enter methods and object: ")
		method, object := "", ""
		fmt.Scan(&method, &object)

		switch object {
		// BRANCH
		case "branch":
			switch method {
			case "create":
				fmt.Println("Enter name, adress and founded year: ")
				name, adress, year := "", "", ""
				fmt.Scan(&name, &adress, &year)
				h.CreateBranch(name, adress, year)
			case "getAll":
				h.GetAllBranch(1, 10, "")
			case "get":
				fmt.Print("Enter ID: ")
				var id string
				fmt.Scan(&id)
				h.GetBranch(id)
			case "update":
				fmt.Println("Enter ID, name, adress and founded year: ")
				id, name, adress, year := "", "", "", ""
				fmt.Scan(&id, &name, &adress, &year)
				h.UpdateBranch(id, name, adress, year)
			}

		}
	}

}
