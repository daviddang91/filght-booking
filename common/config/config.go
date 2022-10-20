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

func GetCustomerGrpcAddress() string {
	return os.Getenv("CUSTOMER_GRPC_ADDRESS")
}

func GetCustomerApiAddress() string {
	return os.Getenv("CUSTOMER_API_ADDRESS")
}

func GetFlightGrpcAddress() string {
	return os.Getenv("FLIGHT_GRPC_ADDRESS")
}

func GetFlightApiAddress() string {
	return os.Getenv("FLIGHT_API_ADDRESS")
}

func GetBookingGrpcAddress() string {
	return os.Getenv("BOOKING_GRPC_ADDRESS")
}

func GetBookingApiAddress() string {
	return os.Getenv("BOOKING_API_ADDRESS")
}
