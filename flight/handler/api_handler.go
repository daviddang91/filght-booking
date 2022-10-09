package handler

import (
	"net/http"

	cDto "github.com/daviddang91/filght-booking/common/dto"
	"github.com/daviddang91/filght-booking/flight/dao"
	"github.com/gin-gonic/gin"
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
