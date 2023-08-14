package handler

import (
	"fmt"
	"project/models"
)

func (h *handler) CreateStaff(branchId, tariffId int, staffType models.StaffType, name string, balance float64) {
	resp, err := h.strg.Staff().CreateStaff(models.CreateStaffRequest{
		BranchId: branchId,
		TariffId: tariffId,
		Type:     staffType,
		Name:     name,
		Balance:  balance,
	})
	if err != nil {
		fmt.Println("error from Create Staff: ", err.Error())
		return
	}
	fmt.Println("Created new Staff: ", resp)
}

func (h *handler) UpdateStaff(id string, branchId, tariffId int, types, name string, balance float64) {
	resp, err := h.strg.Staff().UpdateStaff(models.Staff{
		Id:       id,
		BranchId: branchId,
		TariffId: tariffId,
		Type:     types,
		Name:     name,
		Balance:  balance,
	})
	if err != nil {
		fmt.Println("error  from Updated  Staff: ", err.Error())
		return
	}
	fmt.Println(resp)
}

func (h *handler) GetStaff(id string) {
	resp, err := h.strg.Staff().GetStaff(models.GetStaffByIdRequest{
		Id: id,
	})
	if err != nil {
		fmt.Println("error  from Get Staff: ", err.Error())
		return
	}
	fmt.Println(resp)
}

func (h *handler) GetAllStaff(page, limit int, types models.StaffType, name string, balanceFrom, balanceTo float64) {

	if page < 1 {
		page = h.cfg.Page
	}
	if limit < 1 {
		limit = h.cfg.Limit
	}
	resp, err := h.strg.Staff().GetAllStaff(models.GetAllStaffRequest{
		Type:        types,
		Name:        name,
		BalanceFrom: balanceFrom,
		BalanceTo:   balanceTo,
		Page:        page,
		Limit:       limit,
	})
	if err != nil {
		fmt.Println("error  from Get All Staff: ", err.Error())
		return
	}
	fmt.Println(resp)

}

func (h *handler) DeleteStaff(id string) {
	resp, err := h.strg.Staff().DeleteStaff(models.GetStaffByIdRequest{
		Id: id,
	})
	if err != nil {
		fmt.Println("error  from Delete Branch: ", err.Error())
		return
	}
	fmt.Println(resp)

}
