package dao

import (
	"github.com/daviddang91/filght-booking/common/util"
	"github.com/daviddang91/filght-booking/flight/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IFlightService interface {
	GetList(ctx *gin.Context) ([]model.Flight, *util.Pagination, error)
	GetByCode(code string) (*model.Flight, error)
	GetById(id string) (*model.Flight, error)
}

type FlightService struct {
	IFlightService,
	DB *gorm.DB
}

func NewFlightService(db *gorm.DB) FlightService {
	return FlightService{DB: db}
}

func (f *FlightService) GetList(ctx *gin.Context) ([]model.Flight, *util.Pagination, error) {
	var flight model.Flight
	var query []model.Flight

	p := util.GeneratePagination(ctx, flight)
	pagination, err := util.Paginator(f.DB, &p, &query)

	return query, pagination, err
}

func (f *FlightService) GetByCode(code string) (*model.Flight, error) {
	query := model.Flight{}
	find := f.DB.Find(&query, model.Flight{Code: code})
	if find.Error != nil {
		return nil, find.Error
	}
	return &query, nil
}

func (f *FlightService) GetById(id string) (*model.Flight, error) {
	query := model.Flight{}
	find := f.DB.Find(&query, "id = ?", id)
	if find.Error != nil {
		return nil, find.Error
	}
	return &query, nil
}
