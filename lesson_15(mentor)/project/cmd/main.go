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

		// if object == "branch" && method == "getAll" {
		// 	h.GetAllBranch(1, 10, "")
		// }
		switch object {
		// BRANCH
		case "branch":
			switch method {
			case "create":
				fmt.Println("Enter name and adress: ")
				name, adress := "", ""
				fmt.Scan(&name, &adress)
				h.CreateBranch(name, adress)
			case "getAll":
				h.GetAllBranch(1, 10, "")
			case "get":
				fmt.Print("Enter ID: ")
				var id string
				fmt.Scan(&id)
				h.GetBranch(id)
			}
		}
	}

}
