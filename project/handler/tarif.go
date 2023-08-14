package handler

import (
	"fmt"
	"project/models"
)

func (h *handler) CreateStaffTarif(amountForCard, amountForCash int, name, types string) {
	resp, err := h.strg.StaffTarif().CreateStaffTarif(models.CreateStaffTarifRequest{

		Name:          name,
		Type:          types, // (fixed, percent)
		AmountForCash: amountForCard,
		AmountForCard: amountForCash,
	})
	if err != nil {
		fmt.Println("error from Create StaffTariff: ", err.Error())
		return
	}
	fmt.Println("Created new StaffTarif: ", resp)
}

func (h *handler) UpdateStaffTarif(id string, amountForCard, amountForCash int, types, name string) {
	resp, err := h.strg.StaffTarif().UpdateStaffTarif(models.StaffTarif{
		Id:            id,
		Name:          name,
		Type:          types, // (fixed, percent)
		AmountForCash: amountForCard,
		AmountForCard: amountForCash,
	})
	if err != nil {
		fmt.Println("error  from Updated  StaffTarif: ", err.Error())
		return
	}
	fmt.Println(resp)
}

func (h *handler) GetStaffTarif(id string) {
	resp, err := h.strg.StaffTarif().GetStaffTarif(models.IdRequest{
		Id: id,
	})

	if err != nil {
		fmt.Println("error from GetStaffTarif: ", err.Error())
		return
	}
	fmt.Println("found StaffTarif with id: ", resp)
}

func (h *handler) GetAllStaffTarif(page, limit int, search string) {
	if page < 1 {
		page = h.cfg.Page
	}
	if limit < 1 {
		limit = h.cfg.Limit
	}

	resp, err := h.strg.StaffTarif().GetAllStaffTarif(models.GetAllStaffTarifRequest{
		Page:  page,
		Limit: limit,
		Name:  search,
	})

	if err != nil {
		fmt.Println("error from GetAllStaff:", err.Error())
		return
	}
	fmt.Println(resp)
}

func (h *handler) DeleteStaffTarif(id string) {
	resp, err := h.strg.StaffTarif().DeleteStaffTarif(models.IdRequest{Id: id})
	if err != nil {
		fmt.Println("error from DeleteStaffTarif:", err.Error())
		return
	}
	fmt.Println(resp)
}
