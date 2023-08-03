package storage

import "backend_bootcamp_17_07_2023/lesson_14/models"

type StorageI interface {
	Branch() BranchesI
	Staff() StaffesI
	Product() ProductI
}

type BranchesI interface {
	CreateBranch(models.CreateBranch) (int, error)
	UpdateBranch(models.Branch) (string, error)
	GetBranch(models.IdRequest) (models.Branch, error)
	GetAllBranch(models.GetAllBranchRequest) (models.GetAllBranch, error)
	DeleteBranch(models.IdRequest) (string, error)
}

type StaffesI interface {
	CreateStaff(models.CreateStaff) (int, error)
	UpdateStaff(models.Staff) (string, error)
	GetStaff(models.IdRequest) (models.Staff, error)
	GetAllStaff(models.GetAllStaffRequest) (models.GetAllStaff, error)
	DeleteStaff(models.IdRequest) (string, error)
}

type ProductI interface {
	CreateProduct(models.CreateProduct) (int, error)
	UpdateProduct(models.Product) (string, error)
	GetProduct(models.IdRequest) (models.Product, error)
	GetAllProduct(models.GetAllProductRequest) (models.GetAllProductResponse, error)
	DeleteProduct(models.IdRequest) (string, error)
}
