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

	return r
}
