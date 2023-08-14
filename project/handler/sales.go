package handler

import (
	"fmt"
	"project/models"
)

func (h *handler) CreateSale(clientName string, branchId, shop_asissentId, cashierId int, price float64, Payment models.Payment, Status models.Status) {
	resp, err := h.strg.Sales().CreateSale(models.CreateSales{
		ClientName:      clientName,
		BranchId:        branchId,
		Shop_asissentId: shop_asissentId,
		CashierId:       cashierId,
		Price:           price,
		PaymentType:     Payment,
		Status:          Status,
	})
	if err != nil {
		fmt.Println("error from Create Sale: ", err.Error())
		return
	}
	fmt.Println("Created new Sale: ", resp)
}

func (h *handler) UpdateSale(Id, clientName string, branchId, shop_asissentId, cashierId int, price float64, Payment models.Payment, Status models.Status) {
	resp, err := h.strg.Sales().UpdateSale(models.Sales{
		Id:              Id,
		ClientName:      clientName,
		BranchId:        branchId,
		Shop_asissentId: shop_asissentId,
		CashierId:       cashierId,
		Price:           price,
		PaymentType:     Payment,
		Status:          Status,
	})
	if err != nil {
		fmt.Println("error  from Updated  Sale: ", err.Error())
		return
	}
	fmt.Println(resp)
}

func (h *handler) GetSale(id string) {
	resp, err := h.strg.Sales().GetSale(models.IdRequest{
		Id: id,
	})
	if err != nil {
		fmt.Println("error  from Get Sale: ", err.Error())
		return
	}
	fmt.Println(resp)
}

func (h *handler) GetAllSale(page, limit int, search string) {
	if page < 1 {
		page = h.cfg.Page
	}
	if limit < 1 {
		limit = h.cfg.Limit
	}

	resp, err := h.strg.Sales().GetAllSale(models.GetAllSalesRequest{
		Page:       page,
		Limit:      limit,
		ClientName: search,
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
