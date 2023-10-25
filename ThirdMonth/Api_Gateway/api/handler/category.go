package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	product_service "gitlab.com/market3723841/api-gateway-service/genproto/protuct-service"
)

// CreateCategory godoc
// @Router       /v1/categories [post]
// @Summary      Create a new category
// @Description  Create a new category with the provided details
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        category     body  product_service.CategoryCreateReq  true  "data of the category"
// @Success      201  {object}  product_service.CategoryCreateResp
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) CreateCategory(ctx *gin.Context) {
	var category = product_service.Category{}

	err := ctx.ShouldBindJSON(&category)
	if err != nil {
		h.handlerResponse(ctx, "CreateCategory", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.CategoryService().Create(ctx, &product_service.CategoryCreateReq{
		Name:     category.Name,
		ParentId: category.ParentId,
	})

	if err != nil {
		h.handlerResponse(ctx, "CategoryService().Create", http.StatusBadRequest, err.Error())

		return
	}

	h.handlerResponse(ctx, "create category response", http.StatusOK, resp)
}

// ListCategories godoc
// @Router       /v1/categories [get]
// @Summary      List categories
// @Description  get categories
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        limit    query     int  false  "limit for response"  Default(10)
// @Param		 page     query     int  false  "page for response"   Default(1)
// @Param        name     query     string false "search by name"
// @Success      200  {array}   product_service.Category
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) GetListCategory(ctx *gin.Context) {
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

	resp, err := h.services.CategoryService().GetList(ctx.Request.Context(), &product_service.CategoryGetListReq{
		Page:  int64(page),
		Limit: int64(limit),
		Name:  ctx.Query("name"),
	})

	if err != nil {
		h.handlerResponse(ctx, "error GetListCategory", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "get list category response", http.StatusOK, resp)
}

// GetCategory godoc
// @Router       /v1/categories/{id} [get]
// @Summary      Get a category by ID
// @Description  Retrieve a category by its unique identifier
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        id   path    string     true    "Category ID to retrieve"
// @Success      200  {object}  product_service.Category
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) GetCategory(ctx *gin.Context) {
	id := ctx.Param("id")

	resp, err := h.services.CategoryService().GetById(ctx.Request.Context(), &product_service.CategoryIdReq{Id: id})
	if err != nil {
		h.handlerResponse(ctx, "error category GetById", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "get category response", http.StatusOK, resp)
}

// UpdateCategory godoc
// @Router       /v1/categories/{id} [put]
// @Summary      Update an existing category
// @Description  Update an existing category with the provided details
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        id       path    string     true    "Category ID to update"
// @Param        category   body    product_service.CategoryUpdateReq  true    "Updated data for the category"
// @Success      200  {object}  product_service.CategoryUpdateResp
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) UpdateCategory(ctx *gin.Context) {
	var category = product_service.Category{}
	category.Id = ctx.Param("id")
	err := ctx.ShouldBindJSON(&category)
	if err != nil {
		h.handlerResponse(ctx, "error while binding", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.CategoryService().Update(ctx.Request.Context(), &product_service.CategoryUpdateReq{
		Id:       category.Id,
		Name:     category.Name,
		ParentId: category.ParentId,
	})

	if err != nil {
		h.handlerResponse(ctx, "error category Update", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "update category response", http.StatusOK, resp.Msg)
}

// DeleteCategory godoc
// @Router       /v1/categories/{id} [delete]
// @Summary      Delete a category
// @Description  delete a category by its unique identifier
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        id   path    string     true    "Category ID to retrieve"
// @Success      200  {object}  product_service.CategoryDeleteResp
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *Handler) DeleteCategory(ctx *gin.Context) {
	id := ctx.Param("id")

	resp, err := h.services.CategoryService().Delete(ctx.Request.Context(), &product_service.CategoryIdReq{Id: id})
	if err != nil {
		h.handlerResponse(ctx, "error category Delete", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "delete category response", http.StatusOK, resp.Msg)
}
