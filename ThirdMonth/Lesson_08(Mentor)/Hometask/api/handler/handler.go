package handler

import (
	grpc_client "playground/cpp-bootcamp/grpc"
	"playground/cpp-bootcamp/pkg/logger"
	"playground/cpp-bootcamp/storage"
)

type Handler struct {
	psqlStrg   storage.StorageI
	redisStrg  storage.CacheI
	log        logger.LoggerI
	grpcClient grpc_client.GrpcClientI
}

func NewHandler(psql storage.StorageI, redisStorage storage.CacheI, loger logger.LoggerI, client grpc_client.GrpcClientI) *Handler {
	return &Handler{psqlStrg: psql, redisStrg: redisStorage, log: loger, grpcClient: client}
}
