package handler

import (
	"fmt"
	"net/http"
	"projectWithGinAndSwagger/models"
	"projectWithGinAndSwagger/pkg/logger"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateSale godoc
// @Router       /sale [POST]
// @Summary      CREATE SALE
// @Description  CREATE SALE BASED ON GIVEN DATA
// @Tags         sale
// @Accept       json
// @Produce      json
// @Param        data  body      models.CreateSale  true  "sale data"
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) CreateSale(c *gin.Context) {
	var sales models.CreateSale
	err := c.ShouldBind(&sales)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}
	resp, err := h.storage.Sales().CreateSale(c.Request.Context(), &sales)
	if err != nil {
		fmt.Println("error Sales Create:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Sales successfully created", "id": resp})
}

// GetSale godoc
// @Router       /sale/{id} [GET]
// @Summary      GET SALE
// @Description  GET SALE BY ID
// @Tags         sale
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "SALE ID" format(uuid)
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetSale(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.storage.Sales().GetSale(c.Request.Context(), &models.IdRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		fmt.Println("error Sales Get:", err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetAllSale godoc
// @Router       /sale [GET]
// @Summary      GetAll Sale
// @Description  GET ALL SALES BASED ON GIVEN LIMIT,PAGE AND SEARCH BY CLIENT_NAME
// @Tags         sale
// @Accept       json
// @Produce      json
// @Param        limit   query     int        false  "limit"         minimum(1)     default(10)
// @Param        page    query     int        false  "page"          minimum(1)     default(1)
// @Param        search  query     string     false  "search"
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
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
	resp, err := h.storage.Sales().GetAllSale(c.Request.Context(), &models.GetAllSaleRequest{
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

// UpdateSale godoc
// @Router       /sale/{id} [PUT]
// @Summary      UPDATE SALE
// @Description  UPDATE SALE BY ID AND GIVEN DATA
// @Tags         sale
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "SALE ID" format(uuid)
// @Param        data  body      models.CreateSale  true  "sale data"
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) UpdateSale(c *gin.Context) {
	var sales models.Sales
	err := c.ShouldBind(&sales)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	sales.Id = c.Param("id")
	resp, err := h.storage.Sales().UpdateSale(c.Request.Context(), &sales)
	if err != nil {
		h.log.Error("error Sales Update:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Sales"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Sales successfully updated", "id": resp})
}

// DeleteSale godoc
// @Router       /sale/{id} [DELETE]
// @Summary      DELETE SALE
// @Description  DELETE SALE BY ID
// @Tags         sale
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "SALE ID" format(uuid)
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) DeleteSale(c *gin.Context) {
	id := c.Param("id")
	resp, err := h.storage.Sales().DeleteSale(c.Request.Context(), &models.IdRequest{Id: id})
	if err != nil {
		h.log.Error("error deleting Sales:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Sales"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sales successfully deleted", "id": resp})
}
