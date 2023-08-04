package main

import (
	"backend_bootcamp_17_07_2023/lesson_14/handler"
	"backend_bootcamp_17_07_2023/lesson_14/storage/memory"
)

func main() {
	strg := memory.NewStorage()
	h := handler.NewHandler(strg)
	//================== BRANCH ==================
	// fmt.Println("Enter name and adress: ")
	// name, adress := "", ""
	// fmt.Scan(&name, &adress)
	// h.CreateBranch(name, adress)

	// fmt.Print("To get Branch enter ID: ")
	// id := 0
	// fmt.Scan(&id)
	// h.GetBranch(id)

	//================== STAFF ==================
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

	//================== PRODUCT ==================
	// fmt.Println("Enter name, card_id, size_id, created_at: ")
	// card_id, size_id := 0, 0
	// name, created_at := "", ""
	// fmt.Scan(&name, &card_id, &size_id, &created_at)
	// h.CreateProduct(name, card_id, size_id, created_at)
	// fmt.Print("To get Product enter ID: ")
	// id := 0
	// fmt.Scan(&id)
	// h.GetProduct(id)

	//================== CLIENT ==================
	// fmt.Println("Enter Name, Card_id and Created_at: ")
	// card_id := 0
	// name, created_at := "", ""
	// fmt.Scan(&name, &card_id, &created_at)
	// h.CreateClient(name, card_id, created_at)
	// fmt.Print("To get Staff enter ID: ")
	// id := 0
	// fmt.Scan(&id)
	// h.GetClient(id)

	// ================== CARD ==================
	// fmt.Println("Enter Name, Quantity, Product_Id and Created_at: ")
	// quantity, product_id := 122, 2
	// name, created_at := "apple", "04-07-2023"
	// // fmt.Scan(&name, &quantity, &product_id, created_at)
	// h.CreateCard(name, quantity, product_id, created_at)
	// fmt.Print("To get Card enter ID: ")
	// id := 0
	// fmt.Scan(&id)
	// h.GetCard(id)

	// ================== SIZE ==================
	// fmt.Println("Enter Name, Price and Created_at: ")
	// price := 1002.
	// name, created_at := "1litr", "04-07-2023"
	// h.CreateSize(name, price, created_at)
	// fmt.Print("To get Size enter ID: ")
	// id := 0
	// fmt.Scan(&id)
	// h.GetSize(id)
}
