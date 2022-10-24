package handler

import (
	"context"
	"log"

	"github.com/daviddang91/filght-booking/common/config"
	"github.com/daviddang91/filght-booking/common/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClient struct {
	customerClient pb.CustomerServiceClient
	flightClient   pb.FlightServiceClient
}

func NewGrpcClient() GrpcClient {
	customerConn, err := grpc.Dial(config.GetCustomerGrpcAddress(), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error dial %v\n", err)
	}

	flightConn, err := grpc.Dial(config.GetFlightGrpcAddress(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error dial %v\n", err)
	}

	return GrpcClient{
		customerClient: pb.NewCustomerServiceClient(customerConn),
		flightClient:   pb.NewFlightServiceClient(flightConn),
	}
}

func (g *GrpcClient) GetFlightDetail(flightId string) (*pb.FlightResponse, error) {
	flight, err := g.flightClient.DetailFlight(context.Background(), &pb.FlightRequest{Id: flightId})
	if err != nil {
		log.Println("Error occurred when calling grpc request to flight")
		return nil, err
	}
	return flight, nil
}

func (g *GrpcClient) GetCustomerDetail(customerId string) (*pb.CustomerResponse, error) {
	customer, err := g.customerClient.DetailCustomer(context.Background(), &pb.CustomerRequest{Id: customerId})
	if err != nil {
		log.Println("Error occurred when calling grpc request to customer")
		return nil, err
	}
	return customer, nil
}
