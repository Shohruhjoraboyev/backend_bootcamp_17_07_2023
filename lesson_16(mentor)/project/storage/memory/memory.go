package memory

import "lesson_15/storage"

type store struct {
	branches    *branchRepo
	staffes     *staffRepo
	sales       *saleRepo
	transaction *transactionRepo
}

func NewStorage(fileName string) storage.StorageI {
	return &store{
		branches:    NewBranchRepo(fileName),
		staffes:     NewStaffRepo(fileName),
		sales:       NewSaleRepo(fileName),
		transaction: NewTransactionRepo(fileName),
	}
}

func (s *store) Branch() storage.BranchesI {
	return s.branches
}

func (s *store) Staff() storage.StaffesI {
	return s.staffes
}

func (s *store) Sales() storage.SalesI {
	return s.sales
}

func (s *store) Transaction() storage.TransactionI {
	return s.transaction
}
