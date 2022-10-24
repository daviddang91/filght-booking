package dto

import (
	"time"

	"github.com/daviddang91/filght-booking/flight/model"
)

type FlightRequest struct {
	Name          string    `json:"name" binding:"required"`
	From          string    `json:"from" binding:"required"`
	To            string    `json:"to" binding:"required"`
	Code          string    `json:"code" binding:"required"`
	TotalSlot     int32     `json:"total_slot" binding:"required"`
	DepartureTime time.Time `json:"departure_time" binding:"required"`
	ArrivalTime   time.Time `json:"arrival_time" binding:"required"`
}

func (f *FlightRequest) BindRequest() model.Flight {
	flight := model.Flight{
		Name:          f.Name,
		From:          f.From,
		To:            f.To,
		Code:          f.Code,
		TotalSlot:     f.TotalSlot,
		DepartureTime: f.DepartureTime,
		ArrivalTime:   f.ArrivalTime,
	}
	return flight
}
