package dao

import (
	"github.com/daviddang91/filght-booking/common/util"
	"github.com/daviddang91/filght-booking/customer/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ICustomerService interface {
	GetList(ctx *gin.Context) ([]model.Customer, *util.Pagination, error)
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

func (c *CustomerService) GetList(ctx *gin.Context) ([]model.Customer, *util.Pagination, error) {
	var customer model.Customer
	var query []model.Customer

	p := util.GeneratePagination(ctx, customer)
	pagination, err := util.Paginator(c.DB, &p, &query)

	return query, pagination, err
}

func (c *CustomerService) Create(customer *model.Customer) error {
	if err := c.DB.Create(&customer).Error; err != nil {
		return err
	}
	return nil
}

func (c *CustomerService) Update(customer *model.Customer) (*model.Customer, error) {
	save := c.DB.Save(customer)
	if save.Error != nil {
		return customer, save.Error
	}
	return customer, nil
}

func (c *CustomerService) GetById(customerId string) (*model.Customer, error) {
	query := model.Customer{}
	find := c.DB.First(&query, "id = ?", customerId)
	if find.Error != nil {
		return nil, find.Error
	}
	return &query, nil
}

func (c *CustomerService) GetByUsername(username string) (*model.Customer, error) {
	query := model.Customer{}
	find := c.DB.First(&query, model.Customer{Username: username})
	if find.Error != nil {
		return nil, find.Error
	}
	return &query, nil
}
