package grpc

import (
	branch_service "branch_service/genproto"
	grpc_client "branch_service/grpc/client"
	"branch_service/grpc/service"
	"branch_service/pkg/logger"
	"branch_service/storage"

	"google.golang.org/grpc"
)

func SetUpServer(log logger.LoggerI, strg storage.StorageI, grpcClient grpc_client.GrpcClientI) *grpc.Server {
	s := grpc.NewServer()
	branch_service.RegisterBranchServiceServer(s, service.NewBranchService(log, strg, grpcClient))
	return s
}
