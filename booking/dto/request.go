package dto

import "github.com/daviddang91/filght-booking/booking/model"

type BookingRequest struct {
	CustomerId   string `json:"customer_id" binding:"required"`
	FlightId     string `json:"flight_id" binding:"required"`
	Code         string `json:"code" binding:"required"`
	ReservedSlot int32  `json:"reserved_slot" binding:"required"`
}

func (b *BookingRequest) BindRequest() model.Booking {
	booking := model.Booking{
		CustomerId:   b.CustomerId,
		FlightId:     b.FlightId,
		Code:         b.Code,
		ReservedSlot: b.ReservedSlot,
	}
	return booking
}
