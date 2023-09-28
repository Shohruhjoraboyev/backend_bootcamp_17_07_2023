package handler

import (
	"fmt"
	"net/http"
	"projectWithGinAndSwagger/models"
	"projectWithGinAndSwagger/pkg/logger"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateStaffTarif godoc
// @Router       /staff_tarif  [POST]
// @Summary      CREATE STAFF TARIFF
// @Description  CREATE STAFF TARIFF BY GIVEN DATA
// @Tags         tariff
// @Accept       json
// @Produce      json
// @Param        data   body      models.CreateStaffTarif  true  "STAFF TARIFF DATA"
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) CreateStaffTarif(c *gin.Context) {
	var tariff models.CreateStaffTarif
	err := c.ShouldBind(&tariff)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.storage.StaffTarif().CreateStaffTarif(c.Request.Context(), &tariff)
	if err != nil {
		fmt.Println("error Tariff Create:", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Tariff successfully created", "id": resp})
}

// GetStaffTariff godoc
// @Router       /staff_tarif/{id}  [GET]
// @Summary      GET STAFF TARIFF
// @Description  GET STAFF TARIFF BY GIVEN ID
// @Tags         tariff
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "TARIFF ID" format(uuid)
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetStaffTarif(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.storage.StaffTarif().GetStaffTarif(c.Request.Context(), &models.IdRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println("error tariff Get:", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": resp})
}

// GetAllStaffTariff godoc
// @Router       /staff_tarif  [GET]
// @Summary      GET ALL STAFF TARIFF
// @Description  gets all tariff based on limit, page and search by name
// @Tags         tariff
// @Accept       json
// @Produce      json
// @Param   limit         query     int        false  "limit"          minimum(1)     default(10)
// @Param   page          query     int        false  "page"          minimum(1)     default(1)
// @Param   search        query     string     false  "search"
func (h *Handler) GetAllStaffTarif(c *gin.Context) {
	h.log.Info("request GetAllTarif")
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

	resp, err := h.storage.StaffTarif().GetAllStaffTarif(c.Request.Context(), &models.GetAllStaffTarifRequest{
		Page:  page,
		Limit: limit,
		Name:  c.Query("search"),
	})
	if err != nil {
		h.log.Error("error StaffTarif GetAllTarif:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	h.log.Warn("response to GetAllTariff")
	c.JSON(http.StatusOK, resp)
}

// UpdateStaffTariff godoc
// @Router       /staff_tarif/{id}  [PUT]
// @Summary      UPDATE STAFF TARIFF
// @Description  UPDATE STAFF TARIFF BY GIVEN ID AND DATA
// @Tags         tariff
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "TARIFF ID" format(uuid)
// @Param        data   body    models.CreateStaffTarif  true  "STAFF TARIFF DATA"
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) UpdateStaffTarif(c *gin.Context) {
	var tariff models.StaffTarif

	err := c.ShouldBind(&tariff)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	tariff.Id = c.Param("id")
	resp, err := h.storage.StaffTarif().UpdateStaffTarif(c.Request.Context(), &tariff)
	if err != nil {
		h.log.Error("error StaffTariff Update:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update StaffTarif"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "StaffTarif successfully updated", "id": resp})
}

// DeleteStaffTariff godoc
// @Router       /staff_tarif/{id}  [DELETE]
// @Summary      DELETE STAFF TARIFF
// @Description  DELETE STAFF TARIFF BY GIVEN ID
// @Tags         tariff
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "TARIFF ID" format(uuid)
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) DeleteStaffTarif(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.storage.StaffTarif().DeleteStaffTarif(c.Request.Context(), &models.IdRequest{Id: id})
	if err != nil {
		h.log.Error("error deleting StaffTarif:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete StaffTarif"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "StaffTarif successfully deleted", "id": resp})
}
