package handler

import (
	"backend_bootcamp_17_07_2023/lesson_14/storage"
)

type handler struct {
	strg storage.StorageI
}

func NewHandler(strg storage.StorageI) *handler {
	return &handler{strg: strg}
}
