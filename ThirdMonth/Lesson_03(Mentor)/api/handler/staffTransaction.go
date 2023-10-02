package handler

import (
	"fmt"
	"net/http"
	"projectWithGinAndSwagger/models"
	"projectWithGinAndSwagger/pkg/logger"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateStaffTRansaction godoc
// @Router       /transaction  [POST]
// @Summary      CREATE STAFF TRANSACTION
// @Description  CREATE STAFF TRANSACTION BY GIVEN DATA
// @Tags         transaction
// @Accept       json
// @Produce      json
// @Param        data   body      models.CreateTransaction  true  "STAFF TRANSACTION DATA"
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) CreateTransaction(c *gin.Context) {
	var Transaction models.CreateTransaction
	err := c.ShouldBind(&Transaction)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}

	resp, err := h.storage.Transaction().CreateTransaction(c.Request.Context(), &Transaction)
	if err != nil {
		fmt.Println("error Transaction Create:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Transaction successfully created", "id": resp})
}

// GetStaffTransaction godoc
// @Router       /transaction/{id}  [GET]
// @Summary      GET STAFF TRANSACTION
// @Description  GET STAFF TRANSACTION BY GIVEN ID
// @Tags         transaction
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "TRANSACTION ID" format(uuid)
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetTransaction(c *gin.Context) {

	id := c.Param("id")

	resp, err := h.storage.Transaction().GetTransaction(c.Request.Context(), &models.IdRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		fmt.Println("error Transaction Get:", err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetAllStaffTransaction godoc
// @Router       /transaction   [GET]
// @Summary      GET  ALL STAFF TRANSACTION
// @Description  GET  ALL STAFF TRANSACTION BY GIVEN LIMIT,PAGE AND SEARCH BY TEXT
// @Tags         transaction
// @Accept       json
// @Produce      json
// @Param        limit   query      int  false  "limit" minimum(1) default(10)
// @Param        page    query      int  false  "page" minimum(1) default(1)
// @Param        serch   query      string  false  "search"
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetAllTransaction(c *gin.Context) {
	h.log.Info("request GetAllTransaction")
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		h.log.Error("error get page:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid page param")
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		h.log.Error("error get limit:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid page param")
		return
	}

	resp, err := h.storage.Transaction().GetAllTransaction(c.Request.Context(), &models.GetAllTransactionRequest{
		Page:  page,
		Limit: limit,
		Text:  c.Query("search"),
	})
	if err != nil {
		h.log.Error("error Transaction GetAllTransaction:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	h.log.Warn("response to GetAllTransaction")
	c.JSON(http.StatusOK, resp)
}

// UpdateStaffTransaction godoc
// @Router       /transaction  [PUT]
// @Summary      Update STAFF TRANSACTION
// @Description  Update STAFF TRANSACTION BY GIVEN ID and DATA
// @Tags         transaction
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "TRANSACTION ID" format(uuid)
// @Param        data   body      models.CreateTransaction  true  "STAFF TRANSACTION DATA"
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) UpdateTransaction(c *gin.Context) {
	var Transaction models.Transaction
	err := c.ShouldBind(&Transaction)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	Transaction.Id = c.Param("id")
	resp, err := h.storage.Transaction().UpdateTransaction(c.Request.Context(), &Transaction)
	if err != nil {
		h.log.Error("error Transaction Update:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction successfully updated", "id": resp})
}

// DeleteStaffTransaction godoc
// @Router       /transaction/{id}  [DELETE]
// @Summary      DELETE STAFF TRANSACTION
// @Description  DELETE STAFF TRANSACTION BY GIVEN ID
// @Tags         transaction
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "TRANSACTION ID" format(uuid)
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) DeleteTransaction(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.storage.Transaction().DeleteTransaction(c.Request.Context(), &models.IdRequest{Id: id})
	if err != nil {
		h.log.Error("error deleting Transaction:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction successfully deleted", "id": resp})
}
