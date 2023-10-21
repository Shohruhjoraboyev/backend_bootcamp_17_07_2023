package storage

import (
	"context"
	pb "sale_service/genproto"
	"time"
)

type StorageI interface {
	Sale() SaleI
}
type CacheI interface {
	Cache() RedisI
}

type RedisI interface {
	Create(ctx context.Context, key string, obj interface{}, ttl time.Duration) error
	Get(ctx context.Context, key string, response interface{}) (bool, error)
	Delete(ctx context.Context, key string) error
}

type SaleI interface {
	CreateSale(context.Context, *pb.CreateSaleRequest) (string, error)
	GetSale(context.Context, *pb.IdRequest) (*pb.Sale, error)
	GetAllSale(context.Context, *pb.ListSaleRequest) (*pb.ListSaleResponse, error)
	UpdateSale(context.Context, *pb.UpdateSaleRequest) (string, error)
	DeleteSale(context.Context, *pb.IdRequest) (string, error)
}
