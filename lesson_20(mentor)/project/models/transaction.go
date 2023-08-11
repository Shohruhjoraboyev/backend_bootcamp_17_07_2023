package models

type CreateTransaction struct {
	Type        string //(withdraw,topup)
	Amount      float64
	Source_type string //(sales,bonus)
	Text        string
	Sale_id     string
	Staff_id    string
	Created_at  string
}

type Transaction struct {
	Id          string
	Type        string //(withdraw,topup)
	Amount      float64
	Source_type string //(sales,bonus)
	Text        string
	Sale_id     string
	Staff_id    string
	Created_at  string
}

type GetAllTransactionRequest struct {
	Page  int
	Limit int
	Text  string
}

type GetAllTransactionResponse struct {
	Transactions []Transaction
	Count        int
}

type TopWorkerRequest struct {
	Type     string
	FromDate string
	ToDate   string
}

type TopWorkerRespond struct {
	Staff []TopWorker
}

type TopWorker struct {
	BranchName string
	StaffName  string
	StaffType  string
	EarnedSum  int
	Time       string
}

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
