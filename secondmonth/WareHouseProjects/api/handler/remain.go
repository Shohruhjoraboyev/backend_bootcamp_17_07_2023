package handler

import (
	"WareHouseProjects/models"
	"WareHouseProjects/pkg/logger"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateRemain godoc
// @Router       /remain [POST]
// @Summary      CREATE Remain
// @Description adds Remain data to db based on given info in body
// @Tags         remain
// @Accept       json
// @Produce      json
// @Param        data  body      models.CreateRemain  true  "Remain data"
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) CreateRemain(c *gin.Context) {
	var Remain models.CreateRemain
	err := c.ShouldBind(&Remain)
	if err != nil {
		h.log.Error("error while binding Remain:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}

	resp, err := h.storage.Remaining().CreateRemain(&Remain)
	if err != nil {
		h.log.Error("error Remain create:", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"code": http.StatusCreated, "message": "success", "resp": resp})
}

// GetRemain godoc
// @Router       /remain/{id} [GET]
// @Summary      GET BY ID
// @Description  gets Remain by ID
// @Tags         remain
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Remain ID" format(uuid)
// @Success      200  {object}  models.Remain
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetRemain(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.storage.Remaining().GetRemain(&models.RemainIdRequest{Id: id})
	if err != nil {
		h.log.Error("error get Remain:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetALlRemains godoc
// @Router       /remain [GET]
// @Summary      LIST Remain
// @Description  gets all Remain based on limit, page and search by name
// @Tags         remain
// @Accept       json
// @Produce      json
// @Param  		 limit         query     int        false  "limit"          minimum(1)     default(10)
// @Param  		 page          query     int        false  "page"           minimum(1)     default(1)
// @Param   	 branch_id        query     string     false  "branch_id"
// @Param   	 category_id        query     string     false  "category_id"
// @Param   	 barcode        query     string     false  "barcode"
// @Success      200  {object}  models.GetAllRemainRequest
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetAllRemain(c *gin.Context) {
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

	resp, err := h.storage.Remaining().GetAllRemain(&models.GetAllRemainRequest{
		Page:        page,
		Limit:       limit,
		Branch_id:   c.Query("search"),
		Category_id: c.Query("search"),
		Barcode:     c.Query("search"),
	})
	if err != nil {
		h.log.Error("error Remain GetAllRemain:", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateRemain godoc
// @Router       /remain/{id} [PUT]
// @Summary      UPDATE Remain
// @Description  UPDATES Remain BASED ON GIVEN DATA AND ID
// @Tags         remain
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of Remain" format(uuid)
// @Param        data  body      models.UpdateRemain  true  "Remain data"
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) UpdateRemain(c *gin.Context) {
	var Remain models.UpdateRemain

	err := c.ShouldBind(&Remain)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	Remain.ID = c.Param("id")
	resp, err := h.storage.Remaining().UpdateRemain(&Remain)
	if err != nil {
		h.log.Error("error Remain update:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "success", "resp": resp})
}

// DeleteRemain godoc
// @Router       /remain/{id} [DELETE]
// @Summary      DELETE Remain BY ID
// @Description  deletes Remain by id
// @Tags         remain
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of Remain" format(uuid)
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) DeleteRemain(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.storage.Remaining().DeleteRemain(&models.RemainIdRequest{Id: id})
	if err != nil {
		h.log.Error("error deleting Remain:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Remain"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Remain successfully deleted", "id": resp})
}
