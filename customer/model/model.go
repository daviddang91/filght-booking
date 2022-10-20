package model

import (
	bm "github.com/daviddang91/filght-booking/common/model"
	"golang.org/x/crypto/bcrypt"
)

type Customer struct {
	bm.BaseModel
	Username string `column=username;gorm:"uniqueIndex;not null"`
	Password string `gorm:"column=password;not null"`
	Email    string `gorm:"column=email;not null"`
	FullName string `gorm:"column=full_namenot null"`
}

func (customer *Customer) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	customer.Password = string(bytes)
	return nil
}

func (customer *Customer) CheckPassword(providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(providedPassword))
	return err == nil
}

func (u *Customer) TableName() string {
	return "customer"
}
