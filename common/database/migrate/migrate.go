package main

import (
	bm "github.com/daviddang91/filght-booking/booking/model"
	"github.com/daviddang91/filght-booking/common/config"
	"github.com/daviddang91/filght-booking/common/database"
	cm "github.com/daviddang91/filght-booking/customer/model"
	fm "github.com/daviddang91/filght-booking/flight/model"
)

func init() {
	config.LoadEnvVariables()
}

func main() {
	db := database.OpenConnection()
	defer database.CloseConnection(db)

	// Auto migrate models to database
	db.AutoMigrate(
		&cm.Customer{},
		&fm.Flight{},
		&bm.Booking{},
	)
}
