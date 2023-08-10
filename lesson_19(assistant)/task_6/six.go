package task6

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"os"
// 	"task/models"
// )

// // 6. har bir branch nechta plus/minus transactionlar soni, plus/minus transactionlar summasini quyidagicha chiqarish:
// //                     Transactions            Summ
// //                     plus   minus        plus     minus
// //     1. Branch1      53      20          853 000  278 000
// //     2. Branch2      38      185         492 000  1 982 000
// // {
// //     "id": 1,
// //     "branch_id": 4,
// //     "product_id": 3,
// //     "type": "plus",
// //     "quantity": 74,
// //     "created_at": "2023-08-09 20:05:37"
// //   },

// func PlusMinus() {
// 	productes, _ := readProduct("data/products.json")
// 	transactions, _ := readTransaction("data/branch_pr_transaction.json")
// 	branches, _ := readBranches("data/branches.json")

// 	var branchNameValues = make(map[string]models.PlusMinus)

// 	// for k, v := range branchNameValues {
// 	// 	fmt.Printf("%s: quantity: %d TranPlus: %d TranMinus %d SumPlus: %d SumMinus: %d\n", k, v.Quantity, v.TranPlus, v.TranMinus, v.SumPlus, v.SumMinus)
// 	// 	fmt.Println()
// 	// }
// }

// // ================================READERS======================================
// func readTransaction(data string) ([]models.Transaction, error) {
// 	var transactions []models.Transaction

// 	d, err := os.ReadFile(data)
// 	if err != nil {
// 		log.Printf("Error while Read data: %+v", err)
// 		return nil, err
// 	}
// 	err = json.Unmarshal(d, &transactions)
// 	if err != nil {
// 		log.Printf("Error while Unmarshal data: %+v", err)
// 		return nil, err
// 	}
// 	return transactions, nil
// }

// func readProduct(data string) ([]models.Products, error) {
// 	var productes []models.Products

// 	d, err := os.ReadFile(data)
// 	if err != nil {
// 		log.Printf("Error while Read data: %+v", err)
// 		return nil, err
// 	}
// 	err = json.Unmarshal(d, &productes)
// 	if err != nil {
// 		log.Printf("Error while Unmarshal data: %+v", err)
// 		return nil, err
// 	}
// 	return productes, nil
// }

// func readBranches(data string) ([]models.Branch, error) {
// 	var branches []models.Branch

// 	branch, err := os.ReadFile(data)
// 	if err != nil {
// 		log.Printf("Error while Read branch: %+v", err)
// 		return nil, err
// 	}
// 	err = json.Unmarshal(branch, &branches)
// 	if err != nil {
// 		log.Printf("Error while Unmarshal branch: %+v", err)
// 		return nil, err
// 	}

// 	return branches, nil
// }
