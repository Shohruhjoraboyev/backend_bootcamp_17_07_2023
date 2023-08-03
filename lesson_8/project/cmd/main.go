package main

import (
	"fmt"
	"playground/project/handler"
	"playground/project/storage/memory"
)

func main() {
	strg := memory.NewStorage()
	h := handler.NewHandler(strg)

	fmt.Println("Enter name and adress: ")
	name, adress := "", ""
	fmt.Scan(&name, &adress)
	h.CreateBranch(name, adress)

	fmt.Print("To get Branch enter ID: ")
	id := 0
	fmt.Scan(&id)
	h.GetBranch(id)

}
