package handler

import (
	"app/models"
	"app/pkg/logger"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateTransaction godoc
// @Router       /transaction [POST]
// @Summary      CREATES TRANSACTION
// @Description  CREATES TRANSACTION BASED ON GIVEN DATA
// @Tags         TRANSACTION
// @Accept       json
// @Produce      json
// @Param        data  body      models.CreateTransaction  true  "transaction data"
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) CreateTransaction(c *gin.Context) {
	var transaction models.CreateTransaction
	err := c.ShouldBind(&transaction)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}

	resp, err := h.storage.Transaction().CreateTransaction(c.Request.Context(), &transaction)
	if err != nil {
		fmt.Println("error from storage create transaction:", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "created", "id": resp})
}

// GetTransaction godoc
// @Router       /transaction/{id} [GET]
// @Summary      GET BY ID
// @Description  get transaction by ID
// @Tags         TRANSACTION
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Transaction ID" format(uuid)
// @Success      200  {object}  models.Transaction
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetTransaction(c *gin.Context) {
	response := models.Transaction{}
	id := c.Param("id")

	ok, err := h.redis.Cache().Get(c.Request.Context(), id, &response)
	if err != nil {
		fmt.Println("not found transaction in cache", id)
	}

	if ok {
		c.JSON(http.StatusOK, response)
		return
	}

	resp, err := h.storage.Transaction().GetTransaction(c.Request.Context(), &models.IdRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println("error from storage get transaction:", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": resp})

	err = h.redis.Cache().Create(c.Request.Context(), id, resp, time.Hour)
	if err != nil {
		fmt.Println("error from storage set transaction in cache:", err.Error())
	}
}

// ListTransaction godoc
// @Router       /transaction [GET]
// @Summary      GET  ALL TRANSACTION
// @Description  gets all transaction based on limit, page and search by name
// @Tags         TRANSACTION
// @Accept       json
// @Produce      json
// @Param   limit         query     int        false  "limit"          minimum(1)     default(10)
// @Param   page         query     int        false  "page"          minimum(1)     default(1)
// @Param   search         query     string        false  "search"
// @Success      200  {object}  models.GetAllTransactionResponse
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetAllTransaction(c *gin.Context) {
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
		h.log.Error("error from storage getAll transaction:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateTransaction godoc
// @Router       /transaction/{id} [PUT]
// @Summary      UPDATE TRANSACTION BY ID
// @Description  UPDATES TRANSACTION BASED ON GIVEN DATA AND ID
// @Tags         TRANSACTION
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of transaction" format(uuid)
// @Param        data  body      models.CreateTransaction  true  "transaction data"
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) UpdateTransaction(c *gin.Context) {
	var transaction models.Transaction
	err := c.ShouldBind(&transaction)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	transaction.Id = c.Param("id")
	resp, err := h.storage.Transaction().UpdateTransaction(c.Request.Context(), &transaction)
	if err != nil {
		h.log.Error("error transaction update:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "updated", "id": resp})

	err = h.redis.Cache().Delete(c.Request.Context(), transaction.Id)
	if err != nil {
		fmt.Println("error deleting transaction in redis: ", err)
	}
}

// DeleteTransaction godoc
// @Router       /transaction/{id} [DELETE]
// @Summary      DELETE TRANSACTION BY ID
// @Description  DELETES TRANSACTION BASED ON ID
// @Tags         TRANSACTION
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of transaction" format(uuid)
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) DeleteTransaction(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.storage.Transaction().DeleteTransaction(c.Request.Context(), &models.IdRequest{Id: id})
	if err != nil {
		h.log.Error("error deleting transaction:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "transaction successfully deleted", "id": resp})

	err = h.redis.Cache().Delete(c.Request.Context(), id)
	if err != nil {
		fmt.Println("error deleting transaction in redis: ", err)
	}
}

// func (h *Handler) GetTopStaffs(c *gin.Context) {

// 	resp, err := h.strg.Transaction().GetTopStaffs(models.TopWorkerRequest{
// 		Type:     Type,
// 		FromDate: fromData,
// 		ToDate:   ToData,
// 	})

// 	branchNamesMap, _ := h.strg.Branch().GetAllBranch(models.GetAllBranchRequest{})

// 	branchName := make(map[string]string)
// 	for _, b := range branchNamesMap.Branches {
// 		branchName[b.Id] = b.Name
// 	}
// 	for _, v := range resp {
// 		fmt.Printf("Branch: %s Staff: %s Earning: %d\n", branchName[v.BranchId], v.Name, v.Money)
// 	}

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }
