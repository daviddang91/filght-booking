package dao

import (
	"log"

	"github.com/daviddang91/filght-booking/booking/model"
	"gorm.io/gorm"
)

type IBookingService interface {
	GetList(username string) ([]model.Booking, error)
	GetById(id int) (*model.Booking, error)
	Create(booking *model.Booking) (*model.Booking, error)
	GetByFlightId(flightId int) ([]model.Booking, error)
}

type BookingService struct {
	IBookingService,
	DB *gorm.DB
}

func NewBookingService(db *gorm.DB) BookingService {
	return BookingService{DB: db}
}

func (b *BookingService) GetList(username string) ([]model.Booking, error) {
	var booking []model.Booking
	find := b.DB.Find(&booking, model.Booking{CustomerUsername: username})
	if find.Error != nil {
		return nil, find.Error
	}
	return booking, nil
}

func (b *BookingService) GetById(id int) (*model.Booking, error) {
	booking := model.Booking{}
	find := b.DB.Find(&booking, model.Booking{ID: id})
	if find.Error != nil {
		return nil, find.Error
	}
	return &booking, nil
}

func (b *BookingService) Create(booking *model.Booking) (*model.Booking, error) {
	create := b.DB.Create(booking)
	if create.Error != nil {
		return nil, create.Error
	}
	return booking, nil
}

func (b *BookingService) GetByFlightId(flightId int) ([]model.Booking, error) {
	var result []model.Booking
	find := b.DB.Find(&result, model.Booking{FlightId: flightId})
	if find.Error != nil {
		log.Println("Error occurred in GetByFlightId", find.Error)
		return nil, find.Error
	}
	return result, nil
}
