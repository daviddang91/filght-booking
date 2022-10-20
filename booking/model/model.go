package model

import "time"

type Booking struct {
	ID               int       `gorm:"column=id;not null;primaryKey"`
	CustomerUsername string    `gorm:"column=customer_username;not null"`
	FlightId         int       `gorm:"column=flight_id;not null"`
	Code             string    `gorm:"column=code;not null"`
	Status           string    `gorm:"column=status;not null"`
	ReservedSlot     int       `gorm:"reserved_slot;not null"`
	BookedDate       time.Time `gorm:"column=booked_date;not null"`
	CreatedAt        time.Time `gorm:"column=created_at;not null"`
	UpdatedAt        time.Time `gorm:"column=updated_at;not null"`
}
