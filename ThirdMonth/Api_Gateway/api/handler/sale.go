package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	sale_service "gitlab.com/market3723841/api-gateway-service/genproto/sale-service"
)

// CreateSale godoc
// @Router       /v1/sales [post]
// @Summary      Create a new sale
// @Description  Create a new sale with the provided details
// @Tags         sales
// @Accept       json
// @Produce      json
// @Param        sale     body  sale_service.SaleCreateReq  true  "data of the sale"
// @Success      201  {object}  sale_service.SaleCreateResp
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) CreateSale(ctx *gin.Context) {
	var sale = sale_service.Sale{}

	err := ctx.ShouldBindJSON(&sale)
	if err != nil {
		h.handlerResponse(ctx, "CreateSale", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.SaleService().Create(ctx, &sale_service.SaleCreateReq{
		BranchId:        sale.BranchId,
		ShopAssistantId: sale.ShopAssistantId,
		CashierId:       sale.CashierId,
		Price:           sale.Price,
		PaymentType:     sale.PaymentType,
		Status:          sale.Status,
		ClientName:      sale.ClientName,
	})

	if err != nil {
		h.handlerResponse(ctx, "SaleService().Create", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "create sale response", http.StatusOK, resp)
}

// ListSales godoc
// @Router       /v1/sales [get]
// @Summary      List sales
// @Description  get sales
// @Tags         sales
// @Accept       json
// @Produce      json
// @Param        limit    query     int  false  "limit for response"  Default(10)
// @Param		 page     query     int  false  "page for response"   Default(1)
// @Param        branch_id     query     string false "search by branch_id"
// @Param        client_name     query     string false "search by client_name"
// @Param        payment_type     query     string false "search by payment_type"
// @Param        shop_assistant_id     query     string false "search by shop_assistant_id"
// @Param        cashier_id     query     string false "search by cashier_id"
// @Param        status     query     string false "search by status"
// @Param        created_at_from     query     string false "search by created_at_from"
// @Param        created_at_to     query     string false "search by created_at_to"
// @Param        price_from     query     string false "search by price_from"
// @Param        price_to     query     string false "search by price_to"
// @Success      200  {array}   sale_service.Sale
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) GetListSale(ctx *gin.Context) {
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

	priceFrom, err := strconv.ParseFloat(ctx.DefaultQuery("price_from", "0"), 64)
	if err != nil {
		h.handlerResponse(ctx, "error get price from", http.StatusBadRequest, err.Error())
		return
	}

	priceTo, err := strconv.ParseFloat(ctx.DefaultQuery("price_to", "999999999"), 64)
	if err != nil {
		h.handlerResponse(ctx, "error get price to", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.SaleService().GetList(ctx.Request.Context(), &sale_service.SaleGetListReq{
		Page:            int64(page),
		Limit:           int64(limit),
		BranchId:        ctx.Query("branch_id"),
		ClientName:      ctx.Query("client_name"),
		PaymentType:     ctx.Query("payment_type"),
		ShopAssistantId: ctx.Query("shop_assistant_id"),
		CashierId:       ctx.Query("cashier_id"),
		Status:          ctx.Query("status"),
		CreatedAtFrom:   ctx.DefaultQuery("created_at_from", "2000-01-01"),
		CreatedAtTo:     ctx.DefaultQuery("created_at_to", "2095-12-12"),
		PriceFrom:       float32(priceFrom),
		PriceTo:         float32(priceTo),
	})

	if err != nil {
		h.handlerResponse(ctx, "error GetListSale", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "get list sale response", http.StatusOK, resp)
}

// GetSale godoc
// @Router       /v1/sales/{id} [get]
// @Summary      Get a sale by ID
// @Description  Retrieve a sale by its unique identifier
// @Tags         sales
// @Accept       json
// @Produce      json
// @Param        id   path    string     true    "Sale ID to retrieve"
// @Success      200  {object}  sale_service.Sale
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) GetSale(ctx *gin.Context) {
	id := ctx.Param("id")

	resp, err := h.services.SaleService().GetById(ctx.Request.Context(), &sale_service.SaleIdReq{Id: id})
	if err != nil {
		h.handlerResponse(ctx, "error sale GetById", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "get sale response", http.StatusOK, resp)
}

// UpdateSale godoc
// @Router       /v1/sales/{id} [put]
// @Summary      Update an existing sale
// @Description  Update an existing sale with the provided details
// @Tags         sales
// @Accept       json
// @Produce      json
// @Param        id       path    string     true    "Sale ID to update"
// @Param        sale   body    sale_service.SaleUpdateReq  true    "Updated data for the sale"
// @Success      200  {object}  sale_service.SaleUpdateResp
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) UpdateSale(ctx *gin.Context) {
	var sale = sale_service.Sale{}
	sale.Id = ctx.Param("id")
	err := ctx.ShouldBindJSON(&sale)
	if err != nil {
		h.handlerResponse(ctx, "error while binding", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.SaleService().Update(ctx.Request.Context(), &sale_service.SaleUpdateReq{
		Id:              sale.Id,
		BranchId:        sale.BranchId,
		ShopAssistantId: sale.ShopAssistantId,
		CashierId:       sale.CashierId,
		Price:           sale.Price,
		PaymentType:     sale.PaymentType,
		Status:          sale.Status,
		ClientName:      sale.ClientName,
	})

	if err != nil {
		h.handlerResponse(ctx, "error sale Update", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "update sale response", http.StatusOK, resp.Msg)
}

// DeleteSale godoc
// @Router       /v1/sales/{id} [delete]
// @Summary      Delete a sale
// @Description  delete a sale by its unique identifier
// @Tags         sales
// @Accept       json
// @Produce      json
// @Param        id   path    string     true    "Sale ID to retrieve"
// @Success      200  {object}  sale_service.SaleDeleteResp
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) DeleteSale(ctx *gin.Context) {
	id := ctx.Param("id")

	resp, err := h.services.SaleService().Delete(ctx.Request.Context(), &sale_service.SaleIdReq{Id: id})
	if err != nil {
		h.handlerResponse(ctx, "error sale Delete", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "delete sale response", http.StatusOK, resp.Msg)
}
