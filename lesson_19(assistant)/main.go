package main

import (
	"fmt"
	"task/fifth"
)

func main() {
	// fmt.Println("1.transactionlar soni bo'yicha top branches")
	// first.CalculateTranTopBranches()
	// fmt.Println()

	// fmt.Println("2.transactionlar summasi bo'yicha top branches")
	// second.CalculateSumOfPriceTopBranches()
	// fmt.Println()

	// fmt.Println("3.transactionda bo'lgan top productlar")
	// third.TopTransactionProducts()
	// fmt.Println()

	// fmt.Println("4.transactionda bo'lgan top categorylar")
	// fourth.TopTransactionCategory()
	// fmt.Println()

	fmt.Println("5.har bir branchda har bir categorydan qancha transaction bo'lgani")
	fifth.TopBranchTransactionCategory()
	fmt.Println()
}
