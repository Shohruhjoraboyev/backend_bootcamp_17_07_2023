package main

import (
	"fmt"
	"task/seventh"
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

	// fmt.Println("5.har bir branchda har bir categorydan qancha transaction bo'lgani")
	// fifth.TopBranchTransactionCategory()
	// fmt.Println()

	// fmt.Println("6. har bir branch nechta plus/minus transactionlar soni, plus/minus transactionlar summasini quyidagicha chiqarish")
	// sixth.PlusMinus()
	// fmt.Println()

	fmt.Println("7. har bir kunda kirgan productlar sonini kamayish tartibida chiqarish")
	seventh.CalculateProductIncome()
	fmt.Println()
}

// 6. har bir branch nechta plus/minus transactionlar soni, plus/minus transactionlar summasini quyidagicha chiqarish:
//                     Transactions            Summ
//                     plus   minus        plus     minus
//     1. Branch1      53      20          853 000  278 000
//     2. Branch2      38      185         492 000  1 982 000

// 7. har bir kunda kirgan productlar sonini kamayish tartibida chiqarish:
//         kun         soni
//     1. 2023-08-04   789
//     2. 2023-08-12   634
//     ...  ...  ...  ...
// 8. Product qancha kiritilgan va chiqarilganligi jadvali:
//     Name    Kiritilgan  Chiqarilgan
//     Olma     345            847
//     Cola     374            219
//     ....     ...       ...   ....

// 9. Filialda qancha summalik product borligi jadvali:

// 1. Branch1        853 000
// 2. Branch2      1 982 000
