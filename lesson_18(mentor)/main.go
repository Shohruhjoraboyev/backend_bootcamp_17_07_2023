package main

import (
	"fmt"
	"task/second"
)

func main() {
	// fmt.Println("1.transactionlar soni bo'yicha top branches")
	// first.CalculateTranTopBranches()
	// fmt.Println()

	fmt.Println("2.transactionlar summasi bo'yicha top branches")
	second.CalculateSumOfPriceTopBranches()
}

// 2.transactionlar summasi bo'yicha top branches
// 3.transactionda bo'lgan top productlar
// 4.transactionda bo'lgan top categorylar
// 5.har bir branchda har bir categorydan qancha transaction bo'lgani
