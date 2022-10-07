package dao

import (
	"github.com/daviddang91/filght-booking/customer/model"
	"gorm.io/gorm"
)

type ICustomerService interface {
	Create(customer *model.Customer) error
	Update(customer *model.Customer) (*model.Customer, error)
	GetById(customerId string) (*model.Customer, error)
	GetByUsername(username string) (*model.Customer, error)
	ChangePassword(customer *model.Customer, password string) bool
}

type CustomerService struct {
	ICustomerService
	DB *gorm.DB
}

func NewCustomerService(db *gorm.DB) CustomerService {
	return CustomerService{DB: db}
}

func (cs *CustomerService) Create(customer *model.Customer) error {
	if err := cs.DB.Create(&customer).Error; err != nil {
		return err
	}
	return nil
}

func (cs *CustomerService) Update(customer *model.Customer) (*model.Customer, error) {
	save := cs.DB.Save(customer)
	if save.Error != nil {
		return customer, save.Error
	}
	return customer, nil
}

func (cs *CustomerService) GetById(customerId string) (*model.Customer, error) {
	query := model.Customer{}
	find := cs.DB.First(&query, customerId)
	if find.Error != nil {
		return nil, find.Error
	}
	return &query, nil
}

func (cs *CustomerService) GetByUsername(username string) (*model.Customer, error) {
	query := model.Customer{}
	find := cs.DB.First(&query, model.Customer{Username: username})
	if find.Error != nil {
		return nil, find.Error
	}
	return &query, nil
}
