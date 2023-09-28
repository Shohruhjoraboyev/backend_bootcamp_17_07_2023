package handler

import (
	"fmt"
	"net/http"
	"projectWithGinAndSwagger/models"

	"github.com/gin-gonic/gin"
)

// GetTopStaff godoc
// @Router       /top_staff [GET]
// @Summary      GET TOP STAFF
// @Description  GET  TOP STAFF
// @Tags         topsatff
// @Accept       json
// @Produce      json
// @Param        from_date   query      string  true  "from_date"
// @Param        to_date   query      string  true  "to_date"
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetTopStaff(c *gin.Context) {
	FromDate := c.Query("from_date")
	Todate := c.Query("to_date")

	resp, err := h.storage.BiznesLogic().GetTopStaffs(&models.TopStaffRequest{FromDate: FromDate, Todate: Todate})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		fmt.Println("error  Top Staff Get: ", err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)

}
