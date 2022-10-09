package main

import (
	"fmt"

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

	// rg := g.Group("/customers")
	// rg.GET("/:id", h.DetailCustomer)
	// rg.PUT("/:id", h.UpdateCustomer)
	// rg.PUT("/:id/change-password", h.ChangePassword)

	apiAddress := fmt.Sprintf("0.0.0.0: %s", config.GetFlightApiPort())
	err := g.Run(apiAddress)
	if err != nil {
		panic(err)
	}
}
