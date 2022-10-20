package main

import (
	"github.com/daviddang91/filght-booking/common/config"
	"github.com/daviddang91/filght-booking/common/database"
	"github.com/daviddang91/filght-booking/flight/dao"
	"github.com/daviddang91/filght-booking/flight/handler"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
}

func main() {
	db := database.OpenConnection()
	defer database.CloseConnection(db)

	flightService := dao.NewFlightService(db)

	// routes
	g := gin.Default()
	h := handler.NewHandler(&flightService)

	g.GET("/ping", h.HealthCheck)

	g.GET("/flights", h.HandleGetList)
	g.GET("/flights/:code", h.HandleGetByCode)

	apiAddress := config.GetFlightApiAddress()
	err := g.Run(apiAddress)
	if err != nil {
		panic(err)
	}
}
