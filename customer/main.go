package main

import (
	"log"
	"net"
	"time"

	"github.com/daviddang91/filght-booking/common/config"
	"github.com/daviddang91/filght-booking/common/database"
	cmMd "github.com/daviddang91/filght-booking/common/middleware"
	"github.com/daviddang91/filght-booking/customer/dao"
	"github.com/daviddang91/filght-booking/customer/grpc/pb"
	"github.com/daviddang91/filght-booking/customer/handler"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func init() {
	config.LoadEnvVariables()
}

func main() {
	db := database.OpenConnection()
	defer database.CloseConnection(db)

	customerService := dao.NewCustomerService(db)

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
		server := &handler.Server{CustomerService: &customerService}
		pb.RegisterCustomerServiceServer(s, server)
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
		log.Println("Serving grpc for customer-service...")
	}

	// routes
	g := gin.Default()
	h := handler.NewHandler(&customerService)

	g.GET("/ping", h.HealthCheck)
	g.POST("/login", h.Login)
	g.POST("/register", h.CreateCustomer)

	rg := g.Group("/customers")
	rg.Use(cmMd.Authenticate())
	rg.GET("/:id", h.DetailCustomer)
	rg.PUT("/:id", h.UpdateCustomer)
	rg.PUT("/:id/change-password", h.ChangePassword)

	apiAddress := config.GetCustomerApiAddress()
	err := g.Run(apiAddress)
	if err != nil {
		panic(err)
	}
}
