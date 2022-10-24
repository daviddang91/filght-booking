package handler

import (
	"net/http"

	"github.com/daviddang91/filght-booking/booking/dao"
	"github.com/daviddang91/filght-booking/booking/dto"
	cDto "github.com/daviddang91/filght-booking/common/dto"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	bookingService dao.IBookingService
	grpcClient     GrpcClient
}

func NewHandler(service dao.IBookingService) Handler {
	return Handler{bookingService: service, grpcClient: NewGrpcClient()}
}

func (h *Handler) HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (h *Handler) HandleReserveBooking(ctx *gin.Context) {
	req := &dto.BookingRequest{}
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

	flight, err := h.grpcClient.GetFlightDetail(req.FlightId)
	if err != nil {
		response := cDto.BuildErrorResponse("Error binding request", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	customer, err := h.grpcClient.GetCustomerDetail(req.CustomerId)
	if err != nil {
		response := cDto.BuildErrorResponse("Error binding request", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	booking := req.BindRequest()
	if err := h.bookingService.Create(&booking); err != nil {
		response := cDto.BuildErrorResponse("Cannot save booking", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	res := dto.BookingnResponse{}
	response := cDto.BuildDetailResponse(res.BindResponse(flight, customer, &booking))
	ctx.JSON(201, &response)
}

// req := dto.BookFlightRequest{}
// err := ctx.ShouldBind(&req)
// if err != nil {
// 	log.Println("Error binding request", err)
// 	util.SendDefaultBadRequestError(ctx)
// 	return
// }
// err = h.validator.Struct(&req)
// if err != nil {
// 	util.SendValidationError(ctx, err)
// 	return
// }

// 	success := h.grpcClient.ValidatePasswordUserService(req.Username, req.Password)
// 	if !success {
// 		log.Println("Invalid cred", err)
// 		ctx.AbortWithStatusJSON(http.StatusForbidden, err)
// 		return
// 	}

// 	detail, err := h.grpcClient.GetFlightDetail(req.FlightId)
// 	if err != nil {
// 		log.Println("Error get result from grpc flight service", err)
// 		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
// 		return
// 	}
// 	reservedFlights, err := h.bookingService.GetByFlightId(req.FlightId)
// 	if err != nil {
// 		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
// 		return
// 	}
// 	seatsAllocated := 0
// 	for _, f := range reservedFlights {
// 		seatsAllocated += f.ReservedSlot
// 	}
// 	if req.NumberOfSlot+seatsAllocated > int(detail.TotalSlot) {
// 		log.Printf("Cannot allocate %v seats; Now: %v/%v", req.NumberOfSlot, seatsAllocated, detail.TotalSlot)
// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, nil)
// 		return
// 	}
// 	create, err := h.bookingService.Create(&database.Booking{
// 		CustomerUsername: req.Username,
// 		FlightId:         req.FlightId,
// 		Code:             detail.Code,
// 		Status:           database.BookingStatusCreated,
// 		ReservedSlot:     req.NumberOfSlot,
// 		BookedDate:       time.Now(),
// 		CreatedAt:        time.Now(),
// 		UpdatedAt:        time.Now(),
// 	})
// 	if err != nil {
// 		log.Println("Cannot save booking", err)
// 		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, create)
// }

// func (h *Handler) HandlerGetListBooking(ctx *gin.Context) {
// 	req := dto.GetListBookingRequest{}
// 	err := ctx.ShouldBind(&req)
// 	if err != nil {
// 		log.Println("Error binding request", err)
// 		util.SendDefaultBadRequestError(ctx)
// 		return
// 	}
// 	err = h.validator.Struct(&req)
// 	if err != nil {
// 		util.SendValidationError(ctx, err)
// 		return
// 	}
// 	list, err := h.bookingService.GetList(req.Username)
// 	if err != nil {
// 		log.Println("Error occurred", err)
// 		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, list)
// }

// func (h *Handler) HandleGetBookingById(ctx *gin.Context) {
// 	id, err := strconv.Atoi(ctx.Param("id"))
// 	if err != nil {
// 		log.Println("ID cannot be parsed", err)
// 		ctx.AbortWithStatusJSON(http.StatusInternalServerError, nil)
// 		return
// 	}
// 	byId, err := h.bookingService.GetById(id)
// 	if err != nil {
// 		log.Println("Error getting booking by ID", err)
// 		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, byId)
// }
