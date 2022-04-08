package handler

import (
	"Trial/BANK-NOVANNA/internal/domain/entity"
	"Trial/BANK-NOVANNA/internal/domain/service"
	"Trial/BANK-NOVANNA/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	customerService service.ICustomerService
}

func NewCustomerHandler(customerService service.ICustomerService) *CustomerHandler {
	var customerHandler = CustomerHandler{}
	customerHandler.customerService = customerService
	return &customerHandler
}

func (h *CustomerHandler) SaveCustomer(c *gin.Context) {
	//
	var customer entity.CustomerViewModel
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	//
	saveCustomerError := customer.Validate()
	if len(saveCustomerError) > 0 {
		response.ResponseCustomError(c, saveCustomerError, http.StatusBadRequest)
		return
	}

	result, err := h.customerService.SaveCustomer(&customer)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
	}

	response.ResponseCreated(c, result)
}

func (h *CustomerHandler) GetAllCustomer(c *gin.Context) {
	result, err := h.customerService.GetAllCustomer()
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	if result == nil {
		result = []entity.CustomerViewModel{}
	}

	response.ResponseOKWithData(c, result)
}

func (h *CustomerHandler) GetDetailCustomer(c *gin.Context) {
	NoKTP, err := strconv.Atoi(c.Param("no_ktp"))
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.customerService.GetDetailCustomer(NoKTP)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	response.ResponseOKWithData(c, result)
}

func (h *CustomerHandler) UpdateCustomer(c *gin.Context) {

	custId, err := strconv.Atoi(c.Param("cust_id"))
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	//

	var customer entity.CustomerViewModel
	err = c.ShouldBindJSON(&customer)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	customer.ID = custId

	//

	saveInvoiceError := customer.Validate()
	if len(saveInvoiceError) > 0 {
		response.ResponseCustomError(c, saveInvoiceError, http.StatusBadRequest)
		return
	}

	result, err := h.customerService.UpdateCustomer(&customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if result == nil {
		result = &entity.CustomerViewModel{}
	}

	c.JSON(http.StatusOK, result)
}

func (h *CustomerHandler) DeleteCustomer(c *gin.Context) {
	noKTP, err := strconv.Atoi(c.Param("no_ktp"))
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.customerService.DeleteCustomer(noKTP)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	response.ResponseOK(c, "Successfully Deleted")
}
