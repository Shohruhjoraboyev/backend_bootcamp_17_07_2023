package main

//constructor  boshqa filedagi structlarni foydalanish va enkapsulatysa uchun ishlatiladi
import (
	"fmt"
	"project/config"
	"project/handler"
	"project/models"
	"project/storage/memory"
)

func main() {

	cfg := config.Load()
	strg := memory.NewStorage(models.FileNames{
		BranchFile: "data/branch.json",
		// StaffFile:       "data/staff.json",
		// TransactionFile: "data/transaction.json",
		// SaleFile:        "data/sale.json",
	})
	h := handler.NewHandler(strg, *cfg)

	fmt.Println("Welcome to food deliver Golang project")
	fmt.Println("Avialable methods: ")

	for _, m := range cfg.Methods {
		fmt.Println("--", m)

	}
	fmt.Println("Aviable objects: ")
	for _, o := range cfg.Objects {
		fmt.Println("--", o)
	}
	for {
		fmt.Print("enter methods and object: ")
		method, object := "", ""
		fmt.Scan(&method, &object)

		if object == "branch" {
			if method == "create" {
				fmt.Println("Enter name, address, and founded year: ")
				name, address, year := "", "", ""
				_, err := fmt.Scan(&name, &address, &year)
				if err != nil {
					fmt.Println(err.Error())
				}
				h.CreateBranch(name, address, year)
			} else if method == "get" {
				fmt.Print("Enter ID: ")
				var id string
				fmt.Scan(&id)
				h.GetBranch(id)
			} else if method == "getAll" {
				fmt.Print("Enter search text: ")
				var search string
				fmt.Scan(&search)
				h.GetAllBranch(1, 10, search)
			} else if method == "update" {
				fmt.Println("Enter ID, name, address, and founded year: ")
				id, name, address, year := "", "", "", ""
				fmt.Scan(&id, &name, &address, &year)
				h.UpdateBranch(id, name, address, year)
			} else if method == "delete" {
				fmt.Print("Enter ID that you want to delete: ")
				id := ""
				fmt.Scan(&id)
				h.DeleteBranch(id)
			}

			// staff
		} else if object == "staff" {
			if method == "create" {
				fmt.Println("Enter branchId, TariffId, type, Name and balance: ")
				branchId, TariffId := 0, 0
				var typId models.StaffType = ""
				name := ""
				balance := 0.0
				fmt.Scan(&branchId, &TariffId, &typId, &name, &balance)
				h.CreateStaff(branchId, TariffId, typId, name, balance)
			} else if method == "get" {
				fmt.Print("Enter ID: ")
				var id string
				fmt.Scan(&id)
				h.GetStaff(id)
			} else if method == "getAll" {
				fmt.Println("Enter Type(cashier, shop_assistant), Name, balanceFrom and BalanceTo: ")
				var typId models.StaffType = ""
				name := ""
				balanceFrom, balanceTo := 0.0, 0.0
				fmt.Scan(&typId, &name, &balanceFrom, &balanceTo)
				h.GetAllStaff(1, 10, typId, name, balanceFrom, balanceTo)
			} else if method == "update" {
				fmt.Println("Enter ID, BranchId, TariffId, Type(cashier, shop_assistant), Name, Balance:")
				id := ""
				BranchId, TariffId := 0, 0
				var TypeId models.StaffType
				Name := ""
				Balance := 0.0
				fmt.Scan(&id, &BranchId, &TariffId, &TypeId, &Name, &Balance)
				h.UpdateStaff(id, BranchId, TariffId, string(TypeId), Name, Balance)
			} else if method == "delete" {
				fmt.Print("Enter ID that you want to delete: ")
				id := ""
				fmt.Scan(&id)
				h.DeleteStaff(id)
			}
		}
		if object == "sale" {
			if method == "create" {
				fmt.Println("Client_name, Branch_id, Shop_asissent_id, Cashier_id, Price, Payment_Type(card, cash), Status(success, cancel): ")
				client_name := ""
				Branch_id, Shop_asissent_id, Cashier_id := 0, 0, 0
				price := 0.0
				var payment models.Payment
				var status models.Status
				fmt.Scan(&client_name, &Branch_id, &Shop_asissent_id, &Cashier_id, &price, &payment, &status)
				h.CreateSale(client_name, Branch_id, Shop_asissent_id, Cashier_id, price, payment, status)
			} else if method == "get" {
				fmt.Print("Enter ID: ")
				var id string
				fmt.Scan(&id)
				h.GetSale(id)
			} else if method == "getAll" {
				fmt.Println("Enter Client name: ")
				client_name := ""
				fmt.Scan(&client_name)
				h.GetAllSale(1, 10, client_name)
			} else if method == "update" {
				fmt.Println("Enter ID, Client_name, Branch_id, Shop_asissent_id, Cashier_id, Price, Payment_Type(card, cash), Status(success, cancel): ")
				id, client_name := "", ""
				Branch_id, Shop_asissent_id, Cashier_id := 0, 0, 0
				price := 0.0
				var payment models.Payment
				var status models.Status
				fmt.Scan(&id, &client_name, &Branch_id, &Shop_asissent_id, &Cashier_id, &price, &payment, &status)
				h.UpdateSale(id, client_name, Branch_id, Shop_asissent_id, Cashier_id, price, payment, status)
			} else if method == "delete" {
				fmt.Print("Enter ID that you want to delete: ")
				id := ""
				fmt.Scan(&id)
				h.DeleteSale(id)
			}
		} else if object == "transaction" {
			if method == "create" {
				fmt.Println("type(withdraw,topup), amount, sourceType(sales,bonus), Text, saleId, staffId:")
				amount, staffId := 0.0, 0
				typ, sourceType, text, saleId := "", "", "", ""
				fmt.Scan(&typ, &amount, &sourceType, &text, &saleId, &staffId)
				h.CreateTransaction(typ, amount, sourceType, text, saleId, staffId)
			} else if method == "get" {
				fmt.Print("Enter ID: ")
				var id string
				fmt.Scan(&id)
				h.GetTransaction(id)
			} else if method == "getAll" {
				fmt.Print("Enter Text: ")
				text := ""
				fmt.Scan(&text)
				h.GetAllTransaction(1, 10, text)
			} else if method == "update" {
				fmt.Println("Enter ID, type(withdraw,topup), amount, sourceType(sales,bonus), Text, saleId, staffId:")
				amount, staffId := 0.0, 0
				Id, typ, sourceType, text, saleId := "", "", "", "", ""
				fmt.Scan(&Id, &typ, &amount, &sourceType, &text, &saleId, &staffId)
				h.UpdateTransaction(Id, typ, amount, sourceType, text, saleId, staffId)
			} else if method == "delete" {
				fmt.Print("Enter ID that you want to delete: ")
				id := ""
				fmt.Scan(&id)
				h.DeleteTransaction(id)
			}
		}

	}

}
