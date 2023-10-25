package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	staff_service "gitlab.com/market3723841/api-gateway-service/genproto/staff-service"
)

// CreateStaff godoc
// @Router       /v1/staffs [post]
// @Summary      Create a new staff
// @Description  Create a new staff with the provided details
// @Tags         staffs
// @Accept       json
// @Produce      json
// @Param        staff     body  staff_service.StaffCreateReq  true  "data of the staff"
// @Success      201  {object}  staff_service.StaffCreateResp
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) CreateStaff(ctx *gin.Context) {
	var staff = staff_service.Staff{}

	err := ctx.ShouldBindJSON(&staff)
	if err != nil {
		h.handlerResponse(ctx, "CreateStaff", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.StaffService().Create(ctx, &staff_service.StaffCreateReq{
		BranchId:  staff.BranchId,
		TariffId:  staff.TariffId,
		Name:      staff.Name,
		Typ:       staff.Typ,
		Balance:   staff.Balance,
		BirthDate: staff.BirthDate,
	})

	if err != nil {
		h.handlerResponse(ctx, "StaffService().Create", http.StatusBadRequest, err.Error())

		return
	}

	h.handlerResponse(ctx, "create staff response", http.StatusOK, resp)
}

// ListStaffs godoc
// @Router       /v1/staffs [get]
// @Summary      List staffs
// @Description  get staffs
// @Tags         staffs
// @Accept       json
// @Produce      json
// @Param        limit    query     int  false  "limit for response"  Default(10)
// @Param		 page     query     int  false  "page for response"   Default(1)
// @Param        name     query     string false "search by name"
// @Param        branch_id     query     string false "search by branch_id"
// @Param        tariff_id     query     string false "search by tariff_id"
// @Success      200  {array}   staff_service.Staff
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) GetListStaff(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil {
		h.handlerResponse(ctx, "error get page", http.StatusBadRequest, err.Error())
		return
	}

	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	if err != nil {
		h.handlerResponse(ctx, "error get limit", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.StaffService().GetList(ctx.Request.Context(), &staff_service.StaffGetListReq{
		Page:     int64(page),
		Limit:    int64(limit),
		Name:     ctx.Query("name"),
		BranchId: ctx.Query("branch_id"),
		TarifId:  ctx.Query("tariff_id"),
	})

	if err != nil {
		h.handlerResponse(ctx, "error GetListStaff", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "get list staff response", http.StatusOK, resp)
}

// GetStaff godoc
// @Router       /v1/staffs/{id} [get]
// @Summary      Get a staff by ID
// @Description  Retrieve a staff by its unique identifier
// @Tags         staffs
// @Accept       json
// @Produce      json
// @Param        id   path    string     true    "Staff ID to retrieve"
// @Success      200  {object}  staff_service.Staff
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) GetStaff(ctx *gin.Context) {
	id := ctx.Param("id")

	resp, err := h.services.StaffService().GetById(ctx.Request.Context(), &staff_service.StaffIdReq{Id: id})
	if err != nil {
		h.handlerResponse(ctx, "error staff GetById", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "get staff response", http.StatusOK, resp)
}

// UpdateStaff godoc
// @Router       /v1/staffs/{id} [put]
// @Summary      Update an existing staff
// @Description  Update an existing staff with the provided details
// @Tags         staffs
// @Accept       json
// @Produce      json
// @Param        id       path    string     true    "Staff ID to update"
// @Param        staff   body    staff_service.StaffUpdateReq  true    "Updated data for the staff"
// @Success      200  {object}  staff_service.StaffUpdateResp
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) UpdateStaff(ctx *gin.Context) {
	var staff = staff_service.Staff{}
	staff.Id = ctx.Param("id")
	err := ctx.ShouldBindJSON(&staff)
	if err != nil {
		h.handlerResponse(ctx, "error while binding", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.StaffService().Update(ctx.Request.Context(), &staff_service.StaffUpdateReq{
		Id:        staff.Id,
		BranchId:  staff.BranchId,
		TariffId:  staff.TariffId,
		Name:      staff.Name,
		Typ:       staff.Typ,
		Balance:   staff.Balance,
		BirthDate: staff.BirthDate,
	})

	if err != nil {
		h.handlerResponse(ctx, "error staff Update", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "update staff response", http.StatusOK, resp.Msg)
}

// DeleteStaff godoc
// @Router       /v1/staffs/{id} [delete]
// @Summary      Delete a staff
// @Description  delete a staff by its unique identifier
// @Tags         staffs
// @Accept       json
// @Produce      json
// @Param        id   path    string     true    "Staff ID to retrieve"
// @Success      200  {object}  staff_service.StaffDeleteResp
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) DeleteStaff(ctx *gin.Context) {
	id := ctx.Param("id")

	resp, err := h.services.StaffService().Delete(ctx.Request.Context(), &staff_service.StaffIdReq{Id: id})
	if err != nil {
		h.handlerResponse(ctx, "error staff Delete", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "delete staff response", http.StatusOK, resp.Msg)
}
