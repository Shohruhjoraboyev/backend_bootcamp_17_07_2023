package handler

import (
	"project/config"
	"project/storage"
)

type handler struct {
	strg storage.StorageI
	cfg  config.Config
}

func NewHandler(strg storage.StorageI, conf config.Config) *handler {
	return &handler{strg: strg, cfg: conf}
}
