package main

import (
	"log"
	"net"
	"time"

	"github.com/daviddang91/filght-booking/common/config"
	"github.com/daviddang91/filght-booking/common/database"
	"github.com/daviddang91/filght-booking/common/grpc/pb"
	"github.com/daviddang91/filght-booking/common/middleware"
	"github.com/daviddang91/filght-booking/flight/dao"
	"github.com/daviddang91/filght-booking/flight/handler"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func init() {
	config.LoadEnvVariables()
}

func main() {
	db := database.OpenConnection()
	defer database.CloseConnection(db)

	flightService := dao.NewFlightService(db)

	// init grpc
	grpcAddr := config.GetFlightGrpcAddress()
	c := make(chan bool)
	go func() {
		list, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			log.Fatalf("Failed to start listener %v", err)
		}
		log.Printf("Listening on %v\n", grpcAddr)
		s := grpc.NewServer()
		server := &handler.Server{FlightService: &flightService}
		pb.RegisterFlightServiceServer(s, server)
		if err = s.Serve(list); err != nil {
			c <- false
			log.Fatalf("Failed to serve %v\n", err)
		}
	}()
	select {
	case success := <-c:
		if !success {
			panic("Cannot init grpc")
		}
	case _ = <-time.After(1 * time.Second):
		log.Println("Serving grpc for flight-service...")
	}

	// routes
	g := gin.Default()
	h := handler.NewHandler(&flightService)

	g.GET("/ping", h.HealthCheck)

	g.GET("/flights", h.HandleGetList)
	g.GET("/flights/:code", h.HandleGetByCode)

	rg := g.Group("/admin")
	rg.Use(middleware.Authenticate())
	rg.POST("flights/", h.CreateFlight)

	apiAddress := config.GetFlightApiAddress()
	err := g.Run(apiAddress)
	if err != nil {
		panic(err)
	}
}
