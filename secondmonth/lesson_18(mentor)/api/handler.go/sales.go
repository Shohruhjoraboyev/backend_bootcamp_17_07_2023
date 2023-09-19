package handler

import (
	"fmt"
	"net/http"
	"projectWithGinAndSwagger/models"
	"projectWithGinAndSwagger/pkg/logger"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateSale(c *gin.Context) {
	var sales models.CreateSale
	err := c.ShouldBind(&sales)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}
	resp, err := h.storage.Sales().CreateSale(&sales)
	if err != nil {
		fmt.Println("error Sales Create:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Sales successfully created", "id": resp})
}

func (h *Handler) GetSale(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.storage.Sales().GetSale(&models.IdRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		fmt.Println("error Sales Get:", err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetAllSale(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		h.log.Error("error get page Sale:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid page param")
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		h.log.Error("error get limit Sale:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid page param in Sales")
		return
	}
	resp, err := h.storage.Sales().GetAllSale(&models.GetAllSaleRequest{
		Page:        page,
		Limit:       limit,
		Client_name: c.Query("search"),
	})
	if err != nil {
		h.log.Error("error Sales GetAllSale:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	h.log.Warn("response to GetAllSale")
	c.JSON(http.StatusOK, resp)

}

func (h *Handler) UpdateSale(c *gin.Context) {
	var sales models.Sales
	err := c.ShouldBind(&sales)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	sales.Id = c.Param("id")
	resp, err := h.storage.Sales().UpdateSale(&sales)
	if err != nil {
		h.log.Error("error Sales Update:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Sales"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Sales successfully updated", "id": resp})

}
func (h *Handler) DeleteSale(c *gin.Context) {
	id := c.Param("id")
	resp, err := h.storage.Sales().DeleteSale(&models.IdRequest{Id: id})
	if err != nil {
		h.log.Error("error deleting Sales:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Sales"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sales successfully deleted", "id": resp})
}
