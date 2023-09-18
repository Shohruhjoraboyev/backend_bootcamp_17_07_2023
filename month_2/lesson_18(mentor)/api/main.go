package api

import (
	"app/api/handler"

	"github.com/gin-gonic/gin"
)

func NewServer(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	r.Use()
	r.POST("/branch", h.CreateBranch)
	r.GET("/branch/:id", h.GetBranch)
	r.GET("/branch", h.GetAllBranch)
	r.PUT("/branch/:id", h.UpdateBranch)
	r.DELETE("/branch/:id", h.DeleteBranch)

	r.POST("/tariff", h.CreateStaffTarif)
	r.GET("/tariff/:id", h.GetStaffTarif)
	r.GET("/tariff", h.GetAllStaffTarif)
	r.PUT("/tariff/:id", h.UpdateStaffTarif)
	r.DELETE("/tariff/:id", h.DeleteStaffTarif)
	return r
}
