package handler

import (
	"net/http"

	cmDto "github.com/daviddang91/filght-booking/common/dto"
	"github.com/daviddang91/filght-booking/customer/dao"
	"github.com/daviddang91/filght-booking/customer/dto"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	customerService dao.ICustomerService
}

func NewHandler(service dao.ICustomerService) Handler {
	return Handler{customerService: service}
}

func (h *Handler) HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (h *Handler) CreateCustomer(ctx *gin.Context) {
	req := &dto.CustomerRequest{}
	if err := ctx.ShouldBind(&req); err != nil {
		if validateErrors, ok := err.(validator.ValidationErrors); ok {
			errMessages := make([]string, 0)
			for _, v := range validateErrors {
				errMessages = append(errMessages, v.Error())
			}

			response := cmDto.BuildErrorResponse("Error binding request", errMessages)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		response := cmDto.BuildErrorResponse("Error binding request", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	customer := req.BindRequest()
	if err := customer.HashPassword(customer.Password); err != nil {
		response := cmDto.BuildErrorResponse("Failed to register user", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if err := h.customerService.Create(&customer); err != nil {
		response := cmDto.BuildErrorResponse("Failed to create user", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	res := dto.CustomerResponse{}
	response := cmDto.BuildDetailResponse(res.BindResponse(&customer))
	ctx.JSON(201, &response)
}

func (h *Handler) DetailCustomer(ctx *gin.Context) {
	customerId := ctx.Param("id")
	customer, err := h.customerService.GetById(customerId)
	if err != nil {
		response := cmDto.BuildErrorResponse("Customer not found", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	res := dto.CustomerResponse{}
	response := cmDto.BuildDetailResponse(res.BindResponse(customer))
	ctx.JSON(200, &response)
}

func (h *Handler) UpdateCustomer(ctx *gin.Context) {
	req := &dto.CustomerRequest{}
	if err := ctx.ShouldBind(&req); err != nil {
		if validateErrors, ok := err.(validator.ValidationErrors); ok {
			errMessages := make([]string, 0)
			for _, v := range validateErrors {
				errMessages = append(errMessages, v.Error())
			}

			response := cmDto.BuildErrorResponse("Error binding request", errMessages)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		response := cmDto.BuildErrorResponse("Error binding request", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	customerId := ctx.Param("id")
	customer, err := h.customerService.GetById(customerId)
	if err != nil {
		response := cmDto.BuildErrorResponse("Customer not found", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	req.BindUpdateRequest(customer)
	updated, err := h.customerService.Update(customer)
	if err != nil {
		response := cmDto.BuildErrorResponse("Error occurred when update customer", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	res := dto.CustomerResponse{}
	response := cmDto.BuildDetailResponse(res.BindResponse(updated))
	ctx.JSON(200, &response)
}

func (h *Handler) ChangePassword(ctx *gin.Context) {
	req := &dto.ChangePasswordRequest{}
	if err := ctx.ShouldBind(&req); err != nil {
		if validateErrors, ok := err.(validator.ValidationErrors); ok {
			errMessages := make([]string, 0)
			for _, v := range validateErrors {
				errMessages = append(errMessages, v.Error())
			}

			response := cmDto.BuildErrorResponse("Error binding request", errMessages)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		response := cmDto.BuildErrorResponse("Error binding request", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	customerId := ctx.Param("id")
	customer, err := h.customerService.GetById(customerId)
	if err != nil {
		response := cmDto.BuildErrorResponse("Customer not found", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	checkPwd := customer.CheckPassword(req.OldPassword)
	if !checkPwd {
		response := cmDto.BuildErrorResponse("Invalid password", cmDto.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	customer.HashPassword(req.NewPassword)

	updated, err := h.customerService.Update(customer)
	if err != nil {
		response := cmDto.BuildErrorResponse("Error occurred when update customer", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	res := dto.CustomerResponse{}
	response := cmDto.BuildDetailResponse(res.BindResponse(updated))
	ctx.JSON(200, &response)
}
