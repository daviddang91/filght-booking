package dto

import (
	"time"

	"gorm.io/gorm"
)

type BaseResponse struct {
	Id        string         `json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type DetailResponse struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

type ListResponse struct {
	Status     bool        `json:"status"`
	Data       interface{} `json:"data"`
	Pagination interface{} `json:"pagination"`
}

type ErrorResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type EmptyObj struct{}

func BuildDetailResponse(data interface{}) DetailResponse {
	res := DetailResponse{
		Status: true,
		Data:   data,
	}
	return res
}

func BuildListResponse(data interface{}, pagination interface{}) ListResponse {
	res := ListResponse{
		Status:     true,
		Data:       data,
		Pagination: pagination,
	}
	return res
}

func BuildErrorResponse(message string, err interface{}) ErrorResponse {
	res := ErrorResponse{
		Status:  false,
		Message: message,
		Errors:  err,
	}
	return res
}
