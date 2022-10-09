package main

import (
	"fmt"

	"github.com/daviddang91/filght-booking/common/config"
	"github.com/daviddang91/filght-booking/common/database"
	cmMd "github.com/daviddang91/filght-booking/common/middleware"
	"github.com/daviddang91/filght-booking/customer/dao"
	"github.com/daviddang91/filght-booking/customer/handler"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
}

func main() {
	db := database.OpenConnection()
	defer database.CloseConnection(db)

	customerService := dao.NewCustomerService(db)

	// routes
	g := gin.Default()
	h := handler.NewHandler(&customerService)

	g.GET("/customer-ping", h.HealthCheck)
	g.POST("/login", h.Login)
	g.POST("/register", h.CreateCustomer)

	rg := g.Group("/customers")
	rg.Use(cmMd.Authenticate())
	rg.GET("/:id", h.DetailCustomer)
	rg.PUT("/:id", h.UpdateCustomer)
	rg.PUT("/:id/change-password", h.ChangePassword)

	apiAddress := fmt.Sprintf("0.0.0.0: %s", config.GetCustomerApiPort())
	err := g.Run(apiAddress)
	if err != nil {
		panic(err)
	}
}
