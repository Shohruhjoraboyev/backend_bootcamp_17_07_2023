package main

import (
	"fmt"
	task13 "task/task_13"
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

	// fmt.Println("6. har bir branch nechta plus/minus transactionlar soni, plus/minus transactionlar summasini quyidagicha chiqarish")
	// task6.PlusMinus()
	// fmt.Println()

	// fmt.Println("7. har bir kunda kirgan productlar sonini kamayish tartibida chiqarish")
	// task7.CalculateProductIncome()
	// fmt.Println()

	// fmt.Println("8. Product qancha kiritilgan va chiqarilganligi jadvali")
	// task8.Task8()
	// fmt.Println()

	// fmt.Println("9. Filialda qancha summalik product borligi jadvali")
	// task9.CalculateProductSum()
	// fmt.Println()

	// fmt.Println("10. har bir user transaction qilgan summasi jadvali")
	// task10.Task10()
	// fmt.Println()

	// fmt.Println("11. har bir user kun bo'yicha nechta va necha sumlik transaction qilgani jadvali")
	// task11.Task11()
	// fmt.Println()

	// fmt.Println("12. har bir user qancha product kiritgani va chiqargani jadvali")
	// task12.Task12()
	// fmt.Println()

	fmt.Println("13. Har bir kunda o'rtacha qancha product kiritilgani va chiqarilgani bo'yicha jadvals")
	task13.Task13()
	fmt.Println()
}

/*

...    ...     ...    ...       ...

14. Har kuni o'rtacha qancha product kiritilgani va chiqarilgani bo'yicha jadval:
    branch      o'rtacha+   o'rtacha-
1. Chilonzor      73         34
2. MGorkiy        60         75
...    ...     ...    ...       ...

15. Har kuni o'rtacha user qancha summa product kiritgani va chiqargani bo'yicha jadval:
    branch      o'rtacha+        o'rtacha-
1. Anvar          370 000         435 000
2. Shuhrat        60 000          875 000
...    ...     ...    ...       ...
*/
