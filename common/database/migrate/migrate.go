package main

import (
	"github.com/daviddang91/filght-booking/common/config"
	"github.com/daviddang91/filght-booking/common/database"
	cm "github.com/daviddang91/filght-booking/customer/model"
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
	)
}
