package print

// import (
// 	"encoding/json"
// 	"lesson_20/models"
// 	"log"
// 	"os"
// )

// // 6. Top ishchilarni chiqarish: berilgan vaqt oralig'ida
// // type dynamic:cashier bo'lsa top kassirlarni, shopAssistant bo'lsa top shop assistantlarni qaytarsin:
// //    Name      Branch        Earned_Sum
// //    Abbos     Chilonzor    10000000

// // "Id": "bc566634-45a5-40d9-abe1-9603d086194f",
// // "Type": "topup",
// // "Amount": 323.3,
// // "Source_type": "sales",
// // "Text": "comment",
// // "Sale_id": "5bd28370-f503-4a54-ade5-efe1eff3cbaa",
// // "Staff_id": "7fb8fe47-3948-447c-957e-f187736ec654",
// // "Created_at": "2023-08-11 18:19:55"

// func Six() {
// 	// branches, _ := readBranches("data/branch.json")
// 	staffes, _ := readStaffs("data/staff.json")
// 	transactions, _ := readTransaction("data/transaction.json")

// 	staffCountMap := make(map[string]string)
// 	branchNameMap := make(map[string]string)

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

// func readStaffs(data string) ([]models.Staff, error) {
// 	var staffes []models.Staff

// 	branch, err := os.ReadFile(data)
// 	if err != nil {
// 		log.Printf("Error while Read staff: %+v", err)
// 		return nil, err
// 	}
// 	err = json.Unmarshal(branch, &staffes)
// 	if err != nil {
// 		log.Printf("Error while Unmarshal staff: %+v", err)
// 		return nil, err
// 	}
// 	return staffes, nil
// }

// type transaction struct {
// 	Id          string
// 	Type        string //(withdraw,topup)
// 	Amount      int
// 	Source_type string //(sales,bonus)
// 	Text        string
// 	Sale_id     string
// 	Staff_id    string
// 	Created_at  string
// }

// func readTransaction(data string) ([]transaction, error) {
// 	var sales []transaction

// 	branch, err := os.ReadFile(data)
// 	if err != nil {
// 		log.Printf("Error while Read transaction: %+v", err)
// 		return nil, err
// 	}
// 	err = json.Unmarshal(branch, &sales)
// 	if err != nil {
// 		log.Printf("Error while Unmarshal transaction: %+v", err)
// 		return nil, err
// 	}
// 	return sales, nil
// }
