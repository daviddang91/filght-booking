package model

import (
	"log"
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	Id        uuid.UUID      `gorm:"type:uuid;primaryKey"`
	CreatedAt time.Time      `gorm:"column=created_at"`
	UpdatedAt time.Time      `gorm:"column=updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column=deleted_at;index"`
}

func (base *BaseModel) BeforeCreate(tx *gorm.DB) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}

	tx.Statement.SetColumn("Id", uuid.String())
	return nil
}
