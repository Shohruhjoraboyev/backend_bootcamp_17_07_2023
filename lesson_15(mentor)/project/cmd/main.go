package main

import (
	"backend_bootcamp_17_07_2023/lesson_8/project/config"
	"backend_bootcamp_17_07_2023/lesson_8/project/handler"
	"backend_bootcamp_17_07_2023/lesson_8/project/storage/memory"
	"fmt"
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

		if object == "branch" && method == "getAll" {
			h.GetAllBranch(1, 10, "")
		}
	}

	// fmt.Println("Enter name and adress: ")
	// name, adress := "", ""
	// fmt.Scan(&name, &adress)
	// h.CreateBranch(name, adress)

	// fmt.Print("To get Branch enter ID: ")
	// id := 0
	// fmt.Scan(&id)
	// h.GetBranch(id)

	// STAFF
	// fmt.Println("Enter BranchId, TariffId, TypeId, Name and Balance: ")
	// branchId, tariffId, typeId := 0, 0, 0
	// name := ""
	// balance := 0.0
	// fmt.Scan(&branchId, &tariffId, &typeId, &name, &balance)
	// h.CreateStaff(branchId, tariffId, typeId, name, balance)

	// fmt.Print("To get Staff enter ID: ")
	// id := 0
	// fmt.Scan(&id)
	// h.GetStaff(id)
}
