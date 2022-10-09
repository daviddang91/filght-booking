package model

import (
	"time"

	bm "github.com/daviddang91/filght-booking/common/model"
)

type Flight struct {
	bm.BaseModel
	Code          string    `gorm:"column=code;not null"`
	TotalSlot     int       `gorm:"column=total_slot;not null"`
	DepartureTime time.Time `gorm:"column=departure_time;not null"`
	ArrivalTime   time.Time `gorm:"column=arrival_time;not null"`
}

func (f *Flight) TableName() string {
	return "flight"
}
