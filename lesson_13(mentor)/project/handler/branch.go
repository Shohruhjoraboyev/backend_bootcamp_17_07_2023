package handler

import (
	"backend_bootcamp_17_07_2023/lesson_8/project/models"
	"fmt"
)

func (h *handler) CreateBranch(name, adress string) {
	resp, err := h.strg.Branch().CreateBranch(models.CreateBranch{
		Name:   name,
		Adress: adress,
	})

	if err != nil {
		fmt.Println("error from CreatBranch: ", err.Error())
		return
	}
	fmt.Println("created new branch with id: ", resp)
}

func (h *handler) UpdateBranch(id int, name, adress string) {
	resp, err := h.strg.Branch().UpdateBranch(models.Branch{
		Id:     1,
		Name:   name,
		Adress: adress,
	})

	if err != nil {
		fmt.Println("error from UpdateBranch: ", err.Error())
		return
	}
	fmt.Println("Updated branch with id: ", resp)
}

func (h *handler) GetBranch(id int) {
	resp, err := h.strg.Branch().GetBranch(models.IdRequest{
		Id: 1,
	})

	if err != nil {
		fmt.Println("error from GetBranch: ", err.Error())
		return
	}
	fmt.Println("found branch with id: ", resp)
}

func (h *handler) GetAllBranch(page, limit int, search string) {
	resp, err := h.strg.Branch().GetAllBranch(models.GetAllBranchRequest{
		Page:  page,
		Limit: limit,
		Name:  search,
	})

	if err != nil {
		fmt.Println("error from GetAllBranch: ", err.Error())
		return
	}
	fmt.Println("found all branchs based on filter: ", resp)
}

func (h *handler) DeleteBranch(id int) {
	resp, err := h.strg.Branch().DeleteBranch(models.IdRequest{
		Id: id,
	})

	if err != nil {
		fmt.Println("error from DeleteBranch: ", err.Error())
		return
	}
	fmt.Println("deleted branch with id: ", resp)
}