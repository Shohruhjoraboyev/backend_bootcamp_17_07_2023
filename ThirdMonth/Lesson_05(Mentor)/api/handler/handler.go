package handler

import (
	"projectWithGinAndSwagger/pkg/logger"
	"projectWithGinAndSwagger/storage"
)

type Handler struct {
	storage      storage.StorageI
	redisStorage storage.CacheI

	log logger.LoggerI
}

func NewHandler(strg storage.StorageI, redisStrg storage.CacheI, loger logger.LoggerI) *Handler {
	return &Handler{storage: strg, redisStorage: redisStrg, log: loger}
}
