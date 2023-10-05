package handler

import (
	"fmt"
	"net/http"
	"projectWithGinAndSwagger/api/handler/response"
	"projectWithGinAndSwagger/config"
	"projectWithGinAndSwagger/models"
	"projectWithGinAndSwagger/pkg/helper"
	"projectWithGinAndSwagger/pkg/logger"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// create person handler
// @Router       /login [post]
// @Summary      create person
// @Description  api for create persons
// @Tags         persons
// @Accept       json
// @Produce      json
// @Param        person    body     models.LoginRequest  true  "data ofStaff"
// @Success      200  {object}  models.LoginRequest
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) Login(c *gin.Context) {
	var req models.LoginRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		res := response.ErrorResp{Code: "BAD REQUEST", Message: "invalid fields in body"}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err != nil {

		resp, err := h.storage.Staff().GetByLogin(c.Request.Context(), &models.LoginRequest{
			Login:    req.Login,
			Password: req.Password,
		})
		if err != nil {
			fmt.Println("error Person GetByLogin:", err.Error())
			res := response.ErrorResp{Code: "INTERNAL ERROR", Message: "internal server error"}
			c.JSON(http.StatusInternalServerError, res)
			return
		}
		if req.Password != resp.Password {
			h.log.Error("error while binding:", logger.Error(err))
			res := response.ErrorResp{Code: "INVALID Password", Message: "invalid password"}
			c.JSON(http.StatusBadRequest, res)
			return
		}

		m := make(map[string]interface{})
		m["user_id"] = resp.Login
		m["branch_id"] = resp.BranchId
		token, err := helper.GenerateJWT(m, config.TokenExpireTime, config.JWTSecretKey)
		c.JSON(http.StatusCreated, models.LoginRes{Token: token})
	}
}

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
	hashedPass, err := helper.GeneratePasswordHash(staff.Password)

	if err != nil {
		h.log.Error("error while generating hash password:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}
	staff.Password = string(hashedPass)

	resp, err := h.storage.Staff().CreateStaff(c.Request.Context(), &staff)
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
	response := models.Staff{}
	id := c.Param("id")

	ok, err := h.redisStorage.Cache().Get(c.Request.Context(), id, &response)
	if err != nil {
		fmt.Println("not found staff in redis cache")
	}

	if ok {
		c.JSON(http.StatusOK, response)
		return
	}
	resp, err := h.storage.Staff().GetStaff(c.Request.Context(), &models.IdRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		fmt.Println("error Staff Get: ", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": resp})

	err = h.redisStorage.Cache().Create(c.Request.Context(), id, resp, time.Hour)
	if err != nil {
		fmt.Println("error staff Create in redis cache:", err.Error())
	}
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

	resp, err := h.storage.Staff().GetAllStaff(c.Request.Context(), &models.GetAllStaffRequest{
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
	resp, err := h.storage.Staff().UpdateStaff(c.Request.Context(), &staff)
	if err != nil {
		h.log.Error("Error Staff Update: ", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Update Staff"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Staff successfully updated", "id": resp})
	err = h.redisStorage.Cache().Delete(c.Request.Context(), staff.Id)
	if err != nil {
		fmt.Println("error staff Delete in redis cache:", err.Error())
	}
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
	resp, err := h.storage.Staff().DeleteStaff(c.Request.Context(), &models.IdRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		fmt.Println("error Staff Delete: ", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Staff successfully deleted", "id": resp})
	err = h.redisStorage.Cache().Delete(c.Request.Context(), id)
	if err != nil {
		fmt.Println("error staff Delete in redis cache:", err.Error())
	}
}

// ChangeStaffPassword  godoc
// @Router       /login/{id} [PUT]
// @Summary      Change STAFF Password
// @Description  Change STAFF Password by id and old password
// @Tags         persons
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "STAFF ID" format(uuid)
// @Param        data   body    models.ChangePassword   true  "STAFF OLD AND NEW PASSWORDS"
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) ChangePassword(c *gin.Context) {
	var req models.ChangePassword
	id := c.Param("id")
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.log.Error("Error While binding", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	Id := c.Param("id")

	staff, err := h.storage.Staff().GetStaff(c.Request.Context(), &models.IdRequest{Id: Id})
	if err != nil {
		h.log.Error("Error Staff get in change password: ", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get Staff fro deleting change balance"})
		return
	}
	err = helper.ComparePasswords([]byte(staff.Password), []byte(req.OldPassword))
	if err != nil {
		h.log.Error("Error Staff Password : ", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to match password"})
		return
	}
	newHashPass, err := helper.GeneratePasswordHash(req.NewPassword)
	if err != nil {
		h.log.Error("Error Staff Password hashing : ", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Hashing password"})
		return
	}

	resp, err := h.storage.Staff().ChangePassword(c.Request.Context(), &models.ReqNewPassword{
		Id:       Id,
		Password: string(newHashPass),
	})
	if err != nil {
		h.log.Error("Error To sending hashed password and id: ", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed sending staff;s new password and id"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password successfully Changed", "id": resp})
	err = h.redisStorage.Cache().Delete(c.Request.Context(), id)
	if err != nil {
		fmt.Println("error staff Delete in redis cache:", err.Error())
	}
}
