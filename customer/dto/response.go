package dto

import (
	cmDto "github.com/daviddang91/filght-booking/common/dto"
	cuM "github.com/daviddang91/filght-booking/customer/model"
)

type CustomerResponse struct {
	cmDto.BaseResponse
	Username string     `json:"username"`
	FullName string     `json:"firstName"`
	Gender   cuM.Gender `json:"gender"`
	Email    string     `json:"email"`
}

func (c *CustomerResponse) BindResponse(customer *cuM.Customer) CustomerResponse {
	result := CustomerResponse{
		BaseResponse: cmDto.BaseResponse{
			ID:        customer.ID,
			CreatedAt: customer.CreatedAt,
			UpdatedAt: customer.UpdatedAt,
			DeletedAt: customer.DeletedAt,
		},
		Username: customer.Username,
		FullName: customer.FullName,
		Email:    customer.Email,
		Gender:   customer.Gender,
	}
	return result
}
