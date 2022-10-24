package dao

import (
	"log"

	"github.com/daviddang91/filght-booking/booking/model"
	"gorm.io/gorm"
)

type IBookingService interface {
	GetByCustomerId(cusomterId string) ([]model.Booking, error)
	GetByFlightId(flightId string) ([]model.Booking, error)
	GetById(bookingId string) (*model.Booking, error)
	Create(booking *model.Booking) error
}

type BookingService struct {
	IBookingService,
	DB *gorm.DB
}

func NewBookingService(db *gorm.DB) BookingService {
	return BookingService{DB: db}
}

func (b *BookingService) GetByCustomerId(cusomterId string) ([]model.Booking, error) {
	var booking []model.Booking
	find := b.DB.Find(&booking, model.Booking{CustomerId: cusomterId})
	if find.Error != nil {
		return nil, find.Error
	}
	return booking, nil
}

func (b *BookingService) GetById(bookingId string) (*model.Booking, error) {
	booking := model.Booking{}
	find := b.DB.Find(&booking, "id = ?", bookingId)
	if find.Error != nil {
		return nil, find.Error
	}
	return &booking, nil
}

func (b *BookingService) Create(booking *model.Booking) error {
	if err := b.DB.Create(&booking).Error; err != nil {
		return err
	}
	return nil
}

func (b *BookingService) GetByFlightId(flightId string) ([]model.Booking, error) {
	var result []model.Booking
	find := b.DB.Find(&result, model.Booking{FlightId: flightId})
	if find.Error != nil {
		log.Println("Error occurred in GetByFlightId", find.Error)
		return nil, find.Error
	}
	return result, nil
}
