package storage

import (
	"context"
	sale_service "example-grpc-server/genproto"
	"time"
)

type StorageI interface {
	Branch() BranchI
}
type CacheI interface {
	Cache() RedisI
}

type RedisI interface {
	Create(ctx context.Context, key string, obj interface{}, ttl time.Duration) error
	Get(ctx context.Context, key string, response interface{}) (bool, error)
	Delete(ctx context.Context, key string) error
}

type BranchI interface {
	CreateBranch(context.Context, *sale_service.CreateBranchRequest) (string, error)
	// GetBranch(context.Context, *models.IdRequest) (*models.Branch, error)
	GetAllBranch(context.Context, *sale_service.ListBranchRequest) (*sale_service.ListBranchResponse, error)
	// UpdateBranch(context.Context, *models.Branch) (string, error)
	// DeleteBranch(context.Context, *models.IdRequest) (string, error)
}
