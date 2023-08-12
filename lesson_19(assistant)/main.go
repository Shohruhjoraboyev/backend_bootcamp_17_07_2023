package main

import (
	"fmt"
	task6 "task/task_6"
)

func main() {
	// fmt.Println("1.transactionlar soni bo'yicha top branches")
	// task1.CalculateTranTopBranches()
	// fmt.Println()

	// fmt.Println("2.transactionlar summasi bo'yicha top branches")
	// task2.CalculateSumOfPriceTopBranches()
	// fmt.Println()

	// fmt.Println("3.transactionda bo'lgan top productlar")
	// task3.TopTransactionProducts()
	// fmt.Println()

	// fmt.Println("4.transactionda bo'lgan top categorylar")
	// task4.TopTransactionCategory()
	// fmt.Println()

	// fmt.Println("5.har bir branchda har bir categorydan qancha transaction bo'lgani")
	// task5.TopBranchTransactionCategory()
	// fmt.Println()

	fmt.Println("6. har bir branch nechta plus/minus transactionlar soni, plus/minus transactionlar summasini quyidagicha chiqarish")
	task6.PlusMinus()
	fmt.Println()

	// fmt.Println("7. har bir kunda kirgan productlar sonini kamayish tartibida chiqarish")
	// task7.CalculateProductIncome()
	// fmt.Println()

	// fmt.Println("8. Product qancha kiritilgan va chiqarilganligi jadvali")
	// task8.Task8()
	// fmt.Println()

	// fmt.Println("9. Filialda qancha summalik product borligi jadvali")
	// task9.CalculateProductSum()
	// fmt.Println()
}

// 6. har bir branch nechta plus/minus transactionlar soni, plus/minus transactionlar summasini quyidagicha chiqarish:
//                     Transactions            Summ
//                     plus   minus        plus     minus
//     1. Branch1      53      20          853 000  278 000
//     2. Branch2      38      185         492 000  1 982 000

// 8. Product qancha kiritilgan va chiqarilganligi jadvali:
//     Name    Kiritilgan  Chiqarilgan
//     Olma     345            847
//     Cola     374            219
//     ....     ...       ...   ....
