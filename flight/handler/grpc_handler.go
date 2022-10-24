package handler

import (
	"context"

	"github.com/daviddang91/filght-booking/common/grpc/pb"
	"github.com/daviddang91/filght-booking/flight/dao"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	pb.FlightServiceServer
	FlightService dao.IFlightService
}

func (s *Server) DetailFlight(ctx context.Context, req *pb.FlightRequest) (*pb.FlightResponse, error) {
	found, err := s.FlightService.GetById(req.Id)

	if err != nil {
		return nil, err
	}

	return &pb.FlightResponse{
		Id:            found.Id.String(),
		Name:          found.Name,
		From:          found.From,
		To:            found.To,
		Code:          found.Code,
		TotalSlot:     found.TotalSlot,
		DepartureTime: timestamppb.New(found.DepartureTime),
		ArrivalTime:   timestamppb.New(found.ArrivalTime),
		Audit: &pb.Audit{
			CreatedAt: timestamppb.New(found.CreatedAt),
			UpdatedAt: timestamppb.New(found.UpdatedAt),
		},
	}, nil
}
