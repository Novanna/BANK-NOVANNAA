package handler

import (
	"Trial/BANK-NOVANNA/internal/domain/entity"
	"Trial/BANK-NOVANNA/internal/domain/service"
	"Trial/BANK-NOVANNA/pkg/jwttoken"
	"Trial/BANK-NOVANNA/pkg/response"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	adminService service.IAdminService
}

func NewAdminHandler(adminService service.IAdminService) *AdminHandler {
	var adminHandler = AdminHandler{}
	adminHandler.adminService = adminService
	return &adminHandler
}

func (h *AdminHandler) RegisterAdmin(c *gin.Context) {
	var registerAdmin entity.ReqisterViewModel
	err := c.ShouldBindJSON(&registerAdmin)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	registerAdminError := registerAdmin.Validate()
	if len(registerAdminError) > 0 {
		response.ResponseCustomError(c, registerAdminError, http.StatusBadRequest)
		return
	}

	result, err := h.adminService.SaveAdmin(&registerAdmin)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	response.ResponseCreated(c, result)
}

func (h *AdminHandler) Login(c *gin.Context) {
	var loginVM entity.LoginViewModel

	err := c.ShouldBindJSON(&loginVM)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	validateAdmin, err := h.adminService.GetAdminByEmailPassword(loginVM)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	if validateAdmin == nil {
		validateAdmin = &entity.Admin{}
	}

	// Generete JWT
	token, err := jwttoken.CreateToken(int64(validateAdmin.ID))
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	adminData := map[string]interface{}{
		"access_token": token.AccessToken,
		"expired":      token.ExpiredToken,
		"admin_id":     validateAdmin.ID,
		"admin":        fmt.Sprintf("%s %s", validateAdmin.FirstName, validateAdmin.LastName),
	}

	response.ResponseOKWithData(c, adminData)
}
