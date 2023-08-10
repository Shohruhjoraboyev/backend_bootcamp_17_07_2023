package handler

import (
	"fmt"
	"project/models"
)

func (h *handler) CreateBranch(name, address, year string) {
	resp, err := h.strg.Branch().CreateBranch(models.CreateBranch{
		Name:      name,
		Address:   address,
		FoundedAt: year,
	})
	if err != nil {
		fmt.Println("error  from Create Branch: ", err.Error())
		return
	}
	fmt.Println("Created new Branch: ", resp)

}

func (h *handler) UpdateBranch(id string, name, address, foundedAt string) {
	resp, err := h.strg.Branch().UpdateBranch(models.Branch{
		Id:        id,
		Name:      name,
		Address:   address,
		FoundedAt: foundedAt,
	})
	if err != nil {
		fmt.Println("error  from Updated  Branch: ", err.Error())
		return
	}
	fmt.Println(resp)

}

func (h *handler) GetBranch(id string) {
	resp, err := h.strg.Branch().GetBranch(models.IdRequest{
		Id: id,
	})
	if err != nil {
		fmt.Println("error  from Get Branch: ", err.Error())
		return
	}
	fmt.Println(resp)

}

func (h *handler) GetAllBranch(page, limit int, search string) {
	resp, err := h.strg.Branch().GetAllBranch(models.GetAllBranchRequest{
		Page:  page,
		Limit: limit,
		Name:  search,
	})
	if err != nil {
		fmt.Println("error  from Get All Branch: ", err.Error())
		return
	}
	fmt.Println(resp)

}

func (h *handler) DeleteBranch(id string) {
	resp, err := h.strg.Branch().DeleteBranch(models.IdRequest{
		Id: id,
	})
	if err != nil {
		fmt.Println("error  from Delete Branch: ", err.Error())
		return
	}
	fmt.Println(resp)

}
