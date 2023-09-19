package handler

import (
	"fmt"
	"net/http"
	"projectWithGinAndSwagger/models"
	"projectWithGinAndSwagger/pkg/logger"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateStaff(c *gin.Context) {
	var staff models.CreateStaff
	err := c.ShouldBindJSON(&staff)

	if err != nil {
		h.log.Error("Error While Binding: ", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body"})
		return
	}

	resp, err := h.storage.Staff().CreateStaff(&staff)
	if err != nil {
		h.log.Error("Error Staff Create:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Staff successfully created", "id": resp})
}

func (h *Handler) GetStaff(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.storage.Staff().GetStaff(&models.IdRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		fmt.Println("error Staff Get: ", err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}
func (h *Handler) GetAllStaff(c *gin.Context) {
	h.log.Info("Request GetAllStaff")
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

	resp, err := h.storage.Staff().GetAllStaff(&models.GetAllStaffRequest{
		Page:  page,
		Limit: limit,
		Name:  c.Query("search"),
	})
	if err != nil {
		h.log.Error("error Staff GetAllStaff:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	h.log.Warn("response to GetAllStaff")
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateStaff(c *gin.Context) {
	var staff models.Staff
	err := c.ShouldBindJSON(&staff)
	if err != nil {
		h.log.Error("Error While binding", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	staff.Id = c.Param("id")
	resp, err := h.storage.Staff().UpdateStaff(&staff)
	if err != nil {
		h.log.Error("Error Staff Update: ", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Update Staff"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Staff successfully updated", "id": resp})
}
func (h *Handler) DeleteStaff(c *gin.Context) {
	id := c.Param("id")
	resp, err := h.storage.Staff().DeleteStaff(&models.IdRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		fmt.Println("error Staff Delete: ", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Staff successfully deleted", "id": resp})
}
