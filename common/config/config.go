package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func GetCustomerGrpcPort() string {
	return os.Getenv("CUSTOMER_GRPC_PORT")
}

func GetCustomerApiPort() string {
	return os.Getenv("CUSTOMER_API_PORT")
}

func GetFlightGrpcPort() string {
	return os.Getenv("FLIGHT_GRPC_PORT")
}

func GetFlightApiPort() string {
	return os.Getenv("FLIGHT_API_PORT")
}

func GetBookingGrpcPort() string {
	return os.Getenv("BOOKING_GRPC_PORT")
}

func GetBookingApiPort() string {
	return os.Getenv("BOOKING_API_PORT")
}
