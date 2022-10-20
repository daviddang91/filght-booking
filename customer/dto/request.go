package dto

import (
	cuM "github.com/daviddang91/filght-booking/customer/model"
)

type CustomerRequest struct {
	Username string `json:"username" validate:"required" binding:"required"`
	Password string `json:"password" validate:"min=3" binding:"required"`
	FullName string `json:"full_name" validate:"min=3" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email" `
}

func (reg *CustomerRequest) BindRequest() cuM.Customer {
	user := cuM.Customer{
		Username: reg.Username,
		Password: reg.Password,
		FullName: reg.FullName,
		Email:    reg.Email,
	}
	return user
}

func (reg *CustomerRequest) BindUpdateRequest(customer *cuM.Customer) {
	customer.Username = reg.Username
	customer.Password = reg.Password
	customer.FullName = reg.FullName
	customer.Email = reg.Email
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
