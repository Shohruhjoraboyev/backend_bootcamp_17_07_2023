package api

import (
	"projectWithGinAndSwagger/api/handler"
	"projectWithGinAndSwagger/pkg/helper"

	_ "projectWithGinAndSwagger/api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

/*
func NewServer(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	// r.Use(helper.StartMiddleware)
	// r.Use(helper.EndMiddleware)

	r.Use(helper.Logger)
	r.POST("/login", h.Login)
	v1 := r.Group("v1")
	v1.Use(helper.AuthMiddleware)
	v1.POST("/person", helper.Validation, h.CreatePerson)
	v1.GET("/person", h.GetAllPersons)
	v1.GET("/person/:id", h.GetPerson)
	r.PUT("/person/:id", helper.Validation, h.UpdatePerson)
	r.DELETE("/person/:id", h.DeletePerson)
	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return r
}

*/

func NewServer(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	r.Use(helper.Logger)
	v1 := r.Group("v1")
	v1.Use(helper.AuthMiddleware)
	//Branches
	r.POST("/branch", h.CreateBranch)
	r.GET("/branch/:id", h.GetBranch)
	r.GET("/branch", h.GetAllBranch)
	r.PUT("/branch/:id", h.UpdateBranch)
	r.DELETE("/branch/:id", h.DeleteBranch)
	//Sales
	r.POST("/sale", h.CreateSale)
	r.GET("/sale/:id", h.GetSale)
	r.GET("/sale", h.GetAllSale)
	r.PUT("/sale/:id", h.UpdateSale)
	r.DELETE("/sale/:id", h.DeleteSale)
	//Staffs
	r.POST("/staff", helper.PasswordMiddleware(), h.CreateStaff)
	r.GET("/staff/:id", h.GetStaff)
	r.GET("/staff", h.GetAllStaff)
	r.PUT("/staff/:id", h.UpdateStaff)
	r.DELETE("/staff/:id", h.DeleteStaff)
	r.POST("/login", h.Login)
	r.PUT("/login", h.ChangePassword)
	//StaffTariffs
	r.POST("/staff_tarif", h.CreateStaffTarif)
	r.GET("/staff_tarif/:id", h.GetStaffTarif)
	r.GET("/staff_tarif", h.GetAllStaffTarif)
	r.PUT("/staff_tarif/:id", h.UpdateStaffTarif)
	r.DELETE("/staff_tarif/:id", h.DeleteStaffTarif)
	//Transactions
	r.POST("/transaction", h.CreateTransaction)
	r.GET("/transaction/:id", h.GetTransaction)
	r.GET("/transaction", h.GetAllTransaction)
	r.PUT("/transaction/:id", h.UpdateTransaction)
	r.DELETE("/transaction/:id", h.DeleteTransaction)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return r
}
