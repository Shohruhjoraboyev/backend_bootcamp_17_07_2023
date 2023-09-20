package handler

import (
	"fmt"
	"net/http"
	"projectWithGinAndSwagger/models"
	"projectWithGinAndSwagger/pkg/logger"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateStaff godoc
// @Router       /staff  [POST]
// @Summary      CREATE STAFF
// @Description  CREATE STAFF BY GIVEN DATA
// @Tags         staff
// @Accept       json
// @Produce      json
// @Param        data   body      models.CreateStaff  true  "STAFF DATA"
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
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

// GetStaff godoc
// @Router       /staff/{id} [GET]
// @Summary      GET STAFF
// @Description  GET STAFF BY ID
// @Tags         staff
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "STAFF ID" format(uuid)
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
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

// GetAllStaff godoc
// @Router       /staff  [GET]
// @Summary      GET ALL STAFF
// @Description  GET ALL  STAFF BY GIVEN LIMIT,PAGE AND SEARCH BY NAME
// @Tags         staff
// @Accept       json
// @Produce      json
// @Param        limit   query      int  false  "limit" minimum(1) default(10)
// @Param        page    query      int  false  "page" minimum(1) default(1)
// @Param        search  query      string  false  "search"
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
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

// UpdateStaff godoc
// @Router       /staff/{id} [PUT]
// @Summary      UPDATE STAFF
// @Description  UPDATE STAFF BY ID
// @Tags         staff
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "STAFF ID" format(uuid)
// @Param        data   body     models.CreateStaff   true  "STAFF DATA"
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
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

// DeleteStaff godoc
// @Router       /staff/{id} [DELETE]
// @Summary      DELETE STAFF
// @Description  DELETE STAFF BY ID
// @Tags         staff
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "STAFF ID" format(uuid)
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
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
