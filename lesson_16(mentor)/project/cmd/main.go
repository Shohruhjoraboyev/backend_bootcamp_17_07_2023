package main

import (
	"fmt"
	"lesson_15/config"
	"lesson_15/handler"
	"lesson_15/models"
	"lesson_15/storage/memory"
)

func main() {

	cfg := config.Load()
	strgBranch := memory.NewStorage("data/branch.json")
	strgStaff := memory.NewStorage("data/staff.json")
	strgSale := memory.NewStorage("data/sale.json")
	handlerBranch := handler.NewHandler(strgBranch, *cfg)
	handlerStaff := handler.NewHandler(strgStaff, *cfg)
	handlerSale := handler.NewHandler(strgSale, *cfg)

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
				handlerBranch.CreateBranch(name, adress, year)
			case "get":
				fmt.Print("Enter ID: ")
				var id string
				fmt.Scan(&id)
				handlerBranch.GetBranch(id)
			case "getAll":
				fmt.Print("Enter search text: ")
				var search string
				fmt.Scan(&search)
				handlerBranch.GetAllBranch(1, 10, search)
			case "update":
				fmt.Println("Enter ID, name, adress and founded year: ")
				id, name, adress, year := "", "", "", ""
				fmt.Scan(&id, &name, &adress, &year)
				handlerBranch.UpdateBranch(id, name, adress, year)
			case "delete":
				fmt.Print("Enter ID that you want to delete: ")
				id := ""
				fmt.Scan(&id)
				handlerBranch.DeleteBranch(id)
			}
		// STAFF
		case "staff":
			switch method {
			case "create":
				fmt.Println("Enter branchId, TariffId, type, Name and balance: ")
				branchId, TariffId := 0, 0
				var typId models.StaffType = ""
				name := ""
				balance := 0.0
				fmt.Scan(&branchId, &TariffId, &typId, &name, &balance)
				handlerStaff.CreateStaff(branchId, TariffId, typId, name, balance)
			case "get":
				fmt.Print("Enter ID: ")
				var id string
				fmt.Scan(&id)
				handlerStaff.GetStaff(id)
			case "getAll":
				fmt.Println("Enter Type(cashier, shop_assistant), Name, balanceFrom and BalanceTo: ")
				var typId models.StaffType = ""
				name := ""
				balanceFrom, balanceTo := 0.0, 0.0
				fmt.Scan(&typId, &name, &balanceFrom, &balanceTo)
				handlerStaff.GetAllStaff(1, 10, typId, name, balanceFrom, balanceTo)
			case "update":
				fmt.Println("Enter ID, BranchId, TariffId, Type(cashier, shop_assistant), Name, Balance:")
				id := ""
				BranchId, TariffId := 0, 0
				var TypeId models.StaffType
				Name := ""
				Balance := 0.0
				fmt.Scan(&id, &BranchId, &TariffId, &TypeId, &Name, &Balance)
				// BranchId int, TariffId int, TypeId models.StaffType, Name string, Balance float64
				handlerStaff.UpdateStaff(id, BranchId, TariffId, TypeId, Name, Balance)
			case "delete":
				fmt.Print("Enter ID that you want to delete: ")
				id := ""
				fmt.Scan(&id)
				handlerStaff.DeleteStaff(id)
			}
		// SALE
		case "sale":
			switch method {
			case "create":
				fmt.Println("Client_name, Branch_id, Shop_asissent_id, Cashier_id, Price, Payment_Type(card, cash), Status(seccess, cancel): ")
				client_name := ""
				Branch_id, Shop_asissent_id, Cashier_id := 0, 0, 0
				price := 0.0
				var payment models.Payment
				var status models.Status
				fmt.Scan(&client_name, &Branch_id, &Shop_asissent_id, &Cashier_id, &price, &payment, &status)
				handlerSale.CreateSale(client_name, Branch_id, Shop_asissent_id, Cashier_id, price, payment, status)
			case "get":
				fmt.Print("Enter ID: ")
				var id string
				fmt.Scan(&id)
				handlerSale.GetSale(id)
			case "getAll":
				fmt.Println("Enter Client name: ")
				client_name := ""
				fmt.Scan(&client_name)
				handlerSale.GetAllSale(1, 10, client_name)
			case "update":
				fmt.Println("Enter ID, Client_name, Branch_id, Shop_asissent_id, Cashier_id, Price, Payment_Type(card, cash), Status(seccess, cancel): ")
				id, client_name := "", ""
				Branch_id, Shop_asissent_id, Cashier_id := 0, 0, 0
				price := 0.0
				var payment models.Payment
				var status models.Status
				fmt.Scan(&id, &client_name, &Branch_id, &Shop_asissent_id, &Cashier_id, &price, &payment, &status)
				handlerSale.UpdateSale(id, client_name, Branch_id, Shop_asissent_id, Cashier_id, price, payment, status)
			case "delete":
				fmt.Print("Enter ID that you want to delete: ")
				id := ""
				fmt.Scan(&id)
				handlerSale.DeleteSale(id)
			}
			// TRANSACTION
		case "transaction":
			switch method {
			case "create":
				fmt.Println("amount, staffId, type(withdraw,topup), sourceType(sales,bonus), Text, saleId: ")
				amount, staffId := 0, 0
				typ, sourceType, text, saleId := "", "", "", ""
				fmt.Scan(&amount, &staffId, &typ, &sourceType, &text, &saleId)
				handlerBranch.CreateTransaction(amount, staffId, typ, sourceType, text, saleId)
			case "get":
				fmt.Print("Enter ID: ")
				var id string
				fmt.Scan(&id)
				handlerBranch.GetTransaction(id)
			case "getAll":
				fmt.Print("Enter Text: ")
				text := ""
				fmt.Scan(&text)
				handlerBranch.GetAllTransaction(1, 10, text)
			case "update":
				fmt.Println("Enter ID, amount, staffId, type(withdraw,topup), sourceType(sales,bonus), Text, saleId: ")
				amount, staffId := 0, 0
				Id, typ, sourceType, text, saleId := "", "", "", "", ""
				fmt.Scan(&Id, &amount, &staffId, &typ, &sourceType, &text, &saleId)
				handlerBranch.UpdateTransaction(amount, staffId, Id, typ, sourceType, text, saleId)
			case "delete":
				fmt.Print("Enter ID that you want to delete: ")
				id := ""
				fmt.Scan(&id)
				handlerBranch.DeleteTransaction(id)
			}
		}
	}
}
