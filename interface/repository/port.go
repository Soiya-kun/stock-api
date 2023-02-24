package repository

import (
	"fmt"

	"gorm.io/gorm"

	"gitlab.com/soy-app/stock-api/domain/entity"
	"gitlab.com/soy-app/stock-api/usecase/port"
)

type StockRepository struct {
	db *gorm.DB
}

func NewStockRepository(db *gorm.DB) port.StockRepository {
	return &StockRepository{db: db}
}

func (r *StockRepository) Create(stocks []entity.Stock) error {
	err := r.db.Create(stocks).Error
	if err != nil {
		fmt.Println("failed to create stocks: %w", err)
		return port.ErrCreateStock
	}
	return nil
}
