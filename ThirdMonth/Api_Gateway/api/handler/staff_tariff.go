package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	staff_service "gitlab.com/market3723841/api-gateway-service/genproto/staff-service"
)

// CreateStaffTariff godoc
// @Router       /v1/staff-tariffs [post]
// @Summary      Create a new staff tariff
// @Description  Create a new staff tariff with the provided details
// @Tags         staff-tariffs
// @Accept       json
// @Produce      json
// @Param        tariff     body  staff_service.TariffCreateReq  true  "data of the staff tariff"
// @Success      201  {object}  staff_service.TariffCreateResp
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) CreateStaffTariff(ctx *gin.Context) {
	var staffTariff = staff_service.StaffTariff{}

	err := ctx.ShouldBindJSON(&staffTariff)
	if err != nil {
		h.handlerResponse(ctx, "CreateStaffTariff", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.StaffTariffService().Create(ctx, &staff_service.TariffCreateReq{
		Name:          staffTariff.Name,
		Typ:           staffTariff.Typ,
		AmountForCash: staffTariff.AmountForCash,
		AmountForCard: staffTariff.AmountForCard,
	})

	if err != nil {
		h.handlerResponse(ctx, "StaffTariffService().Create", http.StatusBadRequest, err.Error())

		return
	}

	h.handlerResponse(ctx, "create staff tariff response", http.StatusOK, resp)
}

// ListStaffTariff godoc
// @Router       /v1/staff-tariffs [get]
// @Summary      List staff-tariffs
// @Description  get staff-tariffs
// @Tags         staff-tariffs
// @Accept       json
// @Produce      json
// @Param        limit    query     int  false  "limit for response"  Default(10)
// @Param		 page     query     int  false  "page for response"   Default(1)
// @Param        type     query     string false "search by type"
// @Success      200  {array}   staff_service.StaffTariff
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) GetListStaffTariff(ctx *gin.Context) {
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

	resp, err := h.services.StaffTariffService().GetList(ctx.Request.Context(), &staff_service.TariffGetListReq{
		Page:  int64(page),
		Limit: int64(limit),
		Typ:   ctx.Query("type"),
	})

	if err != nil {
		h.handlerResponse(ctx, "error GetListStaffTariff", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "get list staff tariff response", http.StatusOK, resp)
}

// GetStaffTariff godoc
// @Router       /v1/staff-tariffs/{id} [get]
// @Summary      Get a staff-tariff by ID
// @Description  Retrieve a staff-tariff by its unique identifier
// @Tags         staff-tariffs
// @Accept       json
// @Produce      json
// @Param        id   path    string     true    "Staff Tariff ID to retrieve"
// @Success      200  {object}  staff_service.StaffTariff
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) GetStaffTariff(ctx *gin.Context) {
	id := ctx.Param("id")

	resp, err := h.services.StaffTariffService().GetById(ctx.Request.Context(), &staff_service.TariffIdReq{Id: id})
	if err != nil {
		h.handlerResponse(ctx, "error staff tariff GetById", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "get staff tariff response", http.StatusOK, resp)
}

// UpdateStaffTariff godoc
// @Router       /v1/staff-tariffs/{id} [put]
// @Summary      Update an existing staff-tariff
// @Description  Update an existing staff-tariff with the provided details
// @Tags         staff-tariffs
// @Accept       json
// @Produce      json
// @Param        id       path    string     true    "Staff Tariff ID to update"
// @Param        staff-tariff   body    staff_service.TariffUpdateReq  true    "Updated data for the staff-tariff"
// @Success      200  {object}  staff_service.TariffUpdateResp
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) UpdateStaffTariff(ctx *gin.Context) {
	var staffTariff = staff_service.StaffTariff{}
	staffTariff.Id = ctx.Param("id")
	err := ctx.ShouldBindJSON(&staffTariff)
	if err != nil {
		h.handlerResponse(ctx, "error while binding", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.StaffTariffService().Update(ctx.Request.Context(), &staff_service.TariffUpdateReq{
		Id:            staffTariff.Id,
		Name:          staffTariff.Name,
		Typ:           staffTariff.Typ,
		AmountForCash: staffTariff.AmountForCash,
		AmountForCard: staffTariff.AmountForCard,
	})

	if err != nil {
		h.handlerResponse(ctx, "error staff tariff Update", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "update staff tariff response", http.StatusOK, resp.Msg)
}

// DeleteStaffTariff godoc
// @Router       /v1/staff-tariffs/{id} [delete]
// @Summary      Delete a staff-tariff
// @Description  delete a staff-tariff by its unique identifier
// @Tags         staff-tariffs
// @Accept       json
// @Produce      json
// @Param        id   path    string     true    "StaffTariff ID to retrieve"
// @Success      200  {object}  staff_service.TariffDeleteResp
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) DeleteStaffTariff(ctx *gin.Context) {
	id := ctx.Param("id")

	resp, err := h.services.StaffTariffService().Delete(ctx.Request.Context(), &staff_service.TariffIdReq{Id: id})
	if err != nil {
		h.handlerResponse(ctx, "error staff tariff Delete", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "delete staff tariff response", http.StatusOK, resp.Msg)
}
