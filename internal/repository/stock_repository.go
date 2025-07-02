package repository

import (
	"stock-app/internal/models"

	"gorm.io/gorm"
)

type StockRepository interface {
	Save(stock *models.Stock) error
	GetAll() ([]models.Stock, error)
	FindBySymbol(symbol string) (*models.Stock, error)
}

type stockRepository struct {
	db *gorm.DB
}

func NewStockRepository(db *gorm.DB) StockRepository {
	return &stockRepository{db}
}

func (r *stockRepository) Save(stock *models.Stock) error {
	return r.db.Create(stock).Error
}

func (r *stockRepository) GetAll() ([]models.Stock, error) {
	var stocks []models.Stock
	err := r.db.Order("price DESC").Find(&stocks).Error
	return stocks, err
}

func (r *stockRepository) FindBySymbol(symbol string) (*models.Stock, error) {
	var stock models.Stock
	err := r.db.Where("symbol = ?", symbol).First(&stock).Error
	if err != nil {
		return nil, err
	}
	return &stock, nil
}
