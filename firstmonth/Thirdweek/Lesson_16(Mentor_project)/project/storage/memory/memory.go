package memory

import (
	"project/models"
	"project/storage"
)

type store struct {
	branches     *branchRepo
	staffs       *staffRepo
	sales        *salesRepo
	transactions *transactionRepo
}

func NewStorage(files models.FileNames) storage.StorageI {
	return &store{
		branches:     NewBranchRepo(files.BranchFile),
		staffs:       NewStaffRepo(files.StaffFile),
		sales:        NewSalesRepo(files.SaleFile),
		transactions: NewTransactionRepo(files.TransactionFile),
	}
}
func (s *store) Branch() storage.BranchesI {
	return s.branches
}

func (s *store) Staff() storage.StaffI {
	return s.staffs
}

func (s *store) Sales() storage.SalesI {
	return s.sales
}
func (s *store) Transaction() storage.TransactionI {
	return s.transactions
}
