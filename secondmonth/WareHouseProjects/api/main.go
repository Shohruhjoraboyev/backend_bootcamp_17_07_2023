package api

import (
	_ "WareHouseProjects/api/docs"
	"WareHouseProjects/api/handler"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewServer(h *handler.Handler) *gin.Engine {
	r := gin.Default()
	//Branches
	r.POST("/branch", h.CreateBranch)
	r.GET("/branch/:id", h.GetBranch)
	r.GET("/branch", h.GetAllBranch)
	r.PUT("/branch/:id", h.UpdateBranch)
	r.DELETE("/branch/:id", h.DeleteBranch)

	//Categories
	r.POST("/category", h.CreateCategory)
	r.GET("/category/:id", h.GetCategory)
	r.GET("/category", h.GetAllCategory)
	r.PUT("/category/:id", h.UpdateCategory)
	r.DELETE("/category/:id", h.DeleteCategory)

	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return r
}
