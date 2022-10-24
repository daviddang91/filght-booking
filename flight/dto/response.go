package dto

import (
	"time"

	cmDto "github.com/daviddang91/filght-booking/common/dto"
	"github.com/daviddang91/filght-booking/flight/model"
)

type FlightResponse struct {
	cmDto.BaseResponse
	Name          string    `json:"name"`
	From          string    `json:"flight_from"`
	To            string    `json:"flight_to"`
	Code          string    `json:"code"`
	TotalSlot     int32     `json:"total_slot"`
	DepartureTime time.Time `json:"departure_time"`
	ArrivalTime   time.Time `json:"arrival_time"`
}

func (f *FlightResponse) BindResponse(flight *model.Flight) FlightResponse {
	result := FlightResponse{
		BaseResponse: cmDto.BaseResponse{
			Id:        flight.Id.String(),
			CreatedAt: flight.CreatedAt,
			UpdatedAt: flight.UpdatedAt,
			DeletedAt: flight.DeletedAt,
		},
		Name:          flight.Name,
		From:          flight.From,
		To:            flight.To,
		Code:          flight.Code,
		TotalSlot:     flight.TotalSlot,
		DepartureTime: flight.DepartureTime,
		ArrivalTime:   flight.ArrivalTime,
	}
	return result
}
