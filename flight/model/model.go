package model

import (
	"time"

	bm "github.com/daviddang91/filght-booking/common/model"
)

type Flight struct {
	bm.BaseModel
	Name          string    `gorm:"column:name;not null"`
	From          string    `gorm:"column:flight_from;not null"`
	To            string    `gorm:"column:flight_to;not null"`
	Code          string    `gorm:"column=code;not null"`
	TotalSlot     int32     `gorm:"column=total_slot;not null"`
	DepartureTime time.Time `gorm:"column=departure_time;not null"`
	ArrivalTime   time.Time `gorm:"column=arrival_time;not null"`
}

func (f *Flight) TableName() string {
	return "flight"
}
