package main

import (
	"github.com/daviddang91/filght-booking/booking/dao"
	"github.com/daviddang91/filght-booking/booking/handler"
	"github.com/daviddang91/filght-booking/common/config"
	"github.com/daviddang91/filght-booking/common/database"
	"github.com/daviddang91/filght-booking/common/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
}

func main() {
	db := database.OpenConnection()
	defer database.CloseConnection(db)

	bookingService := dao.NewBookingService(db)

	// routes
	g := gin.Default()
	h := handler.NewHandler(&bookingService)

	g.GET("/ping", h.HealthCheck)

	rg := g.Group("/bookings")
	rg.Use(middleware.Authenticate())
	rg.POST("/reserve", h.HandleReserveBooking)

	apiAddress := config.GetBookingApiAddress()
	err := g.Run(apiAddress)
	if err != nil {
		panic(err)
	}
}
