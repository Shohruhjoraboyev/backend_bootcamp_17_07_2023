package storage

import models "WareHouseProjects/models"

type StorageI interface {
	Branch() BranchesI

	Close()
}

type BranchesI interface {
	CreateBranch(*models.CreateBranch) (string, error)
	GetBranch(*models.BranchIdRequest) (*models.Branch, error)
	GetAllBranch(*models.GetAllBranchRequest) (*models.GetAllBranchResponse, error)
	UpdateBranch(*models.UpdateBranch) (string, error)
	DeleteBranch(*models.BranchIdRequest) (string, error)
}
