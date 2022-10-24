package model

import (
	"time"

	bm "github.com/daviddang91/filght-booking/common/model"
	cm "github.com/daviddang91/filght-booking/customer/model"
	fm "github.com/daviddang91/filght-booking/flight/model"
)

type Booking struct {
	bm.BaseModel
	CustomerId   string      `gorm:"column=customer_id;not null"`
	Customer     cm.Customer `gorm:"foreignKey:CustomerId"`
	FlightId     string      `gorm:"column=flight_id;not null"`
	Flight       fm.Flight   `gorm:"foreignKey:FlightId"`
	Code         string      `gorm:"column=code;not null"`
	Status       string      `gorm:"column=status;not null;default:success"`
	ReservedSlot int32       `gorm:"reserved_slot;not null"`
	BookedDate   time.Time   `gorm:"column=booked_date;autoCreateTime;not null"`
}

func (u *Booking) TableName() string {
	return "booking"
}
