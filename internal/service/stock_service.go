package service

import (
	"sort"
	"stock-app/internal/models"
	"stock-app/internal/repository"
)

type StockService interface {
	GetAll() ([]models.Stock, error)
	Recommend() ([]models.Stock, error)
}

type stockService struct {
	repo repository.StockRepository
}

func NewStockService(r repository.StockRepository) StockService {
	return &stockService{repo: r}
}

func (s *stockService) GetAll() ([]models.Stock, error) {
	return s.repo.GetAll()
}

func (s *stockService) Recommend() ([]models.Stock, error) {
	stocks, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	// Simple algoritmo de recomendación: top 3 por precio
	sort.Slice(stocks, func(i, j int) bool {
		return stocks[i].Price > stocks[j].Price
	})

	if len(stocks) > 3 {
		stocks = stocks[:3]
	}

	return stocks, nil
}
