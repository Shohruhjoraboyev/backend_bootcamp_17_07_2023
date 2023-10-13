package service

import (
	branch_service "branch_service/genproto"
	grpc_client "branch_service/grpc/client"
	"branch_service/pkg/logger"
	"branch_service/storage"
	"context"
)

type BranchService struct {
	logger  logger.LoggerI
	storage storage.StorageI
	clients grpc_client.GrpcClientI
	branch_service.UnimplementedBranchServiceServer
}

func NewBranchService(log logger.LoggerI, strg storage.StorageI, grpcClients grpc_client.GrpcClientI) *BranchService {
	return &BranchService{
		logger:  log,
		storage: strg,
		clients: grpcClients,
	}
}

func (b *BranchService) Create(ctx context.Context, req *branch_service.CreateBranchRequest) (*branch_service.CreateBranchResponse, error) {
	id, err := b.storage.Branch().CreateBranch(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &branch_service.CreateBranchResponse{Id: id}, nil
}

func (b *BranchService) Get(ctx context.Context, req *branch_service.IdRequest) (*branch_service.GetBranchResponse, error) {
	branch, err := b.storage.Branch().GetBranch(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &branch_service.GetBranchResponse{Branch: branch}, nil
}

func (b *BranchService) List(ctx context.Context, req *branch_service.ListBranchRequest) (*branch_service.ListBranchResponse, error) {
	branches, err := b.storage.Branch().GetAllBranch(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &branch_service.ListBranchResponse{Branches: branches.Branches,
		Count: branches.Count}, nil
}

func (s *BranchService) Update(ctx context.Context, req *branch_service.UpdateBranchRequest) (*branch_service.Response, error) {
	resp, err := s.storage.Branch().UpdateBranch(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &branch_service.Response{Message: resp}, nil
}

func (s *BranchService) Delete(ctx context.Context, req *branch_service.IdRequest) (*branch_service.Response, error) {
	resp, err := s.storage.Branch().DeleteBranch(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &branch_service.Response{Message: resp}, nil
}
