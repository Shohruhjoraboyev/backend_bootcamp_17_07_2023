package handler

import (
	"fmt"
	"net/http"
	"projectWithGinAndSwagger/models"
	"projectWithGinAndSwagger/pkg/logger"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateTransaction(c *gin.Context) {
	var Transaction models.CreateTransaction
	err := c.ShouldBind(&Transaction)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}

	resp, err := h.storage.Transaction().CreateTransaction(&Transaction)
	if err != nil {
		fmt.Println("error Transaction Create:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Transaction successfully created", "id": resp})
}
func (h *Handler) GetTransaction(c *gin.Context) {

	id := c.Param("id")

	resp, err := h.storage.Transaction().GetTransaction(&models.IdRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		fmt.Println("error Transaction Get:", err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

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

	resp, err := h.storage.Transaction().GetAllTransaction(&models.GetAllTransactionRequest{
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
func (h *Handler) UpdateTransaction(c *gin.Context) {
	var Transaction models.Transaction
	err := c.ShouldBind(&Transaction)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	Transaction.Id = c.Param("id")
	resp, err := h.storage.Transaction().UpdateTransaction(&Transaction)
	if err != nil {
		h.log.Error("error Transaction Update:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction successfully updated", "id": resp})
}
func (h *Handler) DeleteTransaction(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.storage.Transaction().DeleteTransaction(&models.IdRequest{Id: id})
	if err != nil {
		h.log.Error("error deleting Transaction:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction successfully deleted", "id": resp})
}
