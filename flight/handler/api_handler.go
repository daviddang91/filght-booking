package handler

import (
	"net/http"

	cDto "github.com/daviddang91/filght-booking/common/dto"
	"github.com/daviddang91/filght-booking/flight/dao"
	fDto "github.com/daviddang91/filght-booking/flight/dto"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	flightService dao.IFlightService
}

func NewHandler(service dao.IFlightService) Handler {
	return Handler{flightService: service}
}

func (h *Handler) HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (h *Handler) HandleGetList(ctx *gin.Context) {
	flights, pagination, err := h.flightService.GetList(ctx)
	if err != nil {
		response := cDto.BuildErrorResponse("Error occurred", err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	response := cDto.BuildListResponse(&flights, &pagination)
	ctx.JSON(200, &response)
}

func (h *Handler) HandleGetByCode(ctx *gin.Context) {
	code := ctx.Param("code")
	result, err := h.flightService.GetByCode(code)
	if err != nil {
		response := cDto.BuildErrorResponse("Error occurred", err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}
	response := cDto.BuildDetailResponse(&result)
	ctx.JSON(200, &response)
}

func (h *Handler) CreateFlight(ctx *gin.Context) {
	req := &fDto.FlightRequest{}
	if err := ctx.ShouldBind(&req); err != nil {
		if validateErrors, ok := err.(validator.ValidationErrors); ok {
			errMessages := make([]string, 0)
			for _, v := range validateErrors {
				errMessages = append(errMessages, v.Error())
			}

			response := cDto.BuildErrorResponse("Error binding request", errMessages)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		response := cDto.BuildErrorResponse("Error binding request", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	flight := req.BindRequest()
	if err := h.flightService.Create(&flight); err != nil {
		response := cDto.BuildErrorResponse("Failed to create flight", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	res := &fDto.FlightResponse{}
	response := cDto.BuildDetailResponse(res.BindResponse(&flight))
	ctx.JSON(201, &response)
}
