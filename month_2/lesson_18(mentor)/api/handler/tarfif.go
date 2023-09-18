package handler

import (
	"app/models"
	"app/pkg/logger"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateStaffTarif(c *gin.Context) {
	var tariff models.CreateStaffTarif
	err := c.ShouldBind(&tariff)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}

	resp, err := h.storage.Tariff().CreateStaffTarif(&tariff)
	if err != nil {
		fmt.Println("error Tariff Create:", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Tariff successfully created", "id": resp})
}

func (h *Handler) GetStaffTarif(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.storage.Tariff().GetStaffTarif(&models.IdRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println("error tariff Get:", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": resp})
}

func (h *Handler) UpdateStaffTarif(c *gin.Context) {

}

func (h *Handler) GetAllStaffTarif(c *gin.Context) {

}

func (h *Handler) DeleteStaffTarif(c *gin.Context) {

}
