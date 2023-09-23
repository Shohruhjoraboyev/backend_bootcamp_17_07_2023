package storage

import models "WareHouseProjects/models"

type StorageI interface {
	Branch() BranchesI
	Category() CategoriesI

	Close()
}

type BranchesI interface {
	CreateBranch(*models.CreateBranch) (string, error)
	GetBranch(*models.BranchIdRequest) (*models.Branch, error)
	GetAllBranch(*models.GetAllBranchRequest) (*models.GetAllBranchResponse, error)
	UpdateBranch(*models.UpdateBranch) (string, error)
	DeleteBranch(*models.BranchIdRequest) (string, error)
}

type CategoriesI interface {
	CreateCategory(*models.CreateCategory) (string, error)
	GetCategory(*models.CategoryIdRequest) (*models.Category, error)
	GetAllCategory(*models.GetAllCategoryRequest) (*models.GetAllCategoryResponse, error)
	UpdateCategory(*models.UpdateCategory) (string, error)
	DeleteCategory(*models.CategoryIdRequest) (string, error)
}
