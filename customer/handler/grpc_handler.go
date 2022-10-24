package handler

import (
	"context"

	"github.com/daviddang91/filght-booking/common/grpc/pb"
	"github.com/daviddang91/filght-booking/customer/dao"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	pb.CustomerServiceServer
	CustomerService dao.ICustomerService
}

func (s *Server) DetailCustomer(ctx context.Context, req *pb.CustomerRequest) (*pb.CustomerResponse, error) {
	found, err := s.CustomerService.GetById(req.Id)

	if err != nil {
		return nil, err
	}

	return &pb.CustomerResponse{
		Id:       found.Id.String(),
		Username: found.Username,
		Email:    found.Email,
		FullName: found.FullName,
		Audit: &pb.Audit{
			CreatedAt: timestamppb.New(found.CreatedAt),
			UpdatedAt: timestamppb.New(found.UpdatedAt),
		},
	}, nil
}
