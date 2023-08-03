package handler

import "playground/project/storage"

type handler struct {
	strg storage.StorageI
}

func NewHandler(strg storage.StorageI) *handler {
	return &handler{strg: strg}
}
