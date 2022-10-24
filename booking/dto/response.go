package dto

import (
	"time"

	"github.com/daviddang91/filght-booking/booking/model"
	cmDto "github.com/daviddang91/filght-booking/common/dto"
	"github.com/daviddang91/filght-booking/common/grpc/pb"
	cDto "github.com/daviddang91/filght-booking/customer/dto"
	fDto "github.com/daviddang91/filght-booking/flight/dto"
)

type BookingnResponse struct {
	cmDto.BaseResponse
	Customer     cDto.CustomerResponse `json:"customer"`
	Flight       fDto.FlightResponse   `json:"flight"`
	Code         string                `json:"code"`
	Status       string                `json:"status"`
	ReservedSlot int32                 `json:"reserved_slot"`
	BookedDate   time.Time             `json:"booked_date"`
}

func (b *BookingnResponse) BindResponse(flight *pb.FlightResponse, customer *pb.CustomerResponse, booking *model.Booking) BookingnResponse {
	result := BookingnResponse{
		BaseResponse: cmDto.BaseResponse{
			Id:        booking.Id.String(),
			CreatedAt: booking.CreatedAt,
			UpdatedAt: booking.UpdatedAt,
			DeletedAt: booking.DeletedAt,
		},
		Customer: cDto.CustomerResponse{
			BaseResponse: cmDto.BaseResponse{
				Id:        customer.Id,
				CreatedAt: customer.Audit.CreatedAt.AsTime(),
				UpdatedAt: customer.Audit.UpdatedAt.AsTime(),
			},
			Username: customer.Username,
			FullName: customer.FullName,
			Email:    customer.Email,
		},
		Flight: fDto.FlightResponse{
			BaseResponse: cmDto.BaseResponse{
				Id:        flight.Id,
				CreatedAt: flight.Audit.CreatedAt.AsTime(),
				UpdatedAt: flight.Audit.UpdatedAt.AsTime(),
			},
			Name:          flight.Name,
			From:          flight.From,
			To:            flight.To,
			Code:          flight.Code,
			TotalSlot:     flight.TotalSlot,
			DepartureTime: flight.DepartureTime.AsTime(),
			ArrivalTime:   flight.ArrivalTime.AsTime(),
		},
		Code:         booking.Code,
		Status:       booking.Status,
		ReservedSlot: booking.ReservedSlot,
		BookedDate:   booking.BookedDate,
	}
	return result
}
