package handler

import (
	"fmt"
	"lesson_15/models"
)

func (h *handler) CreateSale(BranchId, TariffId, TypeId int, Name string, Balance float64) {
	resp, err := h.strg.Sales().CreateSale(models.CreateSales{})
	if err != nil {
		fmt.Println("error from CreateSales: ", err.Error())
		return
	}
	fmt.Println("created new staff with id: ", resp)
}

func (h *handler) UpdateSale(BranchId, TariffId, TypeId int, Name string, Balance float64) {
	resp, err := h.strg.Sales().UpdateSale(models.Sales{})

	if err != nil {
		fmt.Println("error from UpdateSales: ", err.Error())
		return
	}
	fmt.Println("Updated staff with id: ", resp)
}

func (h *handler) GetSale(id string) {
	resp, err := h.strg.Sales().GetSale(models.IdRequest{
		Id: id,
	})

	if err != nil {
		fmt.Println("error from GetSales: ", err.Error())
		return
	}
	fmt.Println("found staff with id: ", resp)
}

func (h *handler) GetAllSale(page, limit int, search string) {
	if page < 1 {
		page = h.cfg.Page
	}
	if limit < 1 {
		limit = h.cfg.Limit
	}

	resp, err := h.strg.Sales().GetAllSale(models.GetAllSalesRequest{
		Page:  page,
		Limit: limit,
		Name:  search,
	})

	if err != nil {
		fmt.Println("error from GetAllSales: ", err.Error())
		return
	}
	fmt.Println("found all Saleses based on filter: ", resp)
}

func (h *handler) DeleteSale(id string) {
	resp, err := h.strg.Sales().DeleteSale(models.IdRequest{
		Id: id,
	})

	if err != nil {
		fmt.Println("error from DeleteSales: ", err.Error())
		return
	}
	fmt.Println("deleted staff with id: ", resp)
}
