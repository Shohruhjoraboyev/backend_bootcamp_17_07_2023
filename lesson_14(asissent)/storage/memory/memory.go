package memory

import "backend_bootcamp_17_07_2023/lesson_14/storage"

type store struct {
	branches *branchRepo
	staffes  *staffRepo
	products *productRepo
}

func NewStorage() storage.StorageI {
	return &store{
		branches: NewBranchRepo(),
		staffes:  NewStaffRepo(),
		products: NewProductRepo(),
	}
}

func (s *store) Branch() storage.BranchesI {
	return s.branches
}

func (s *store) Staff() storage.StaffesI {
	return s.staffes
}

func (s *store) Product() storage.ProductI {
	return s.products
}
