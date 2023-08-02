package repository

import (
	"gorm.io/gorm"

	"gitlab.com/soy-app/stock-api/domain/entity"
	"gitlab.com/soy-app/stock-api/usecase/port"
)

type SearchStockPatternRepository struct {
	db *gorm.DB
}

func NewSearchStockPatternRepository(db *gorm.DB) port.SearchStockPatternRepository {
	return &SearchStockPatternRepository{
		db: db,
	}
}

func (s SearchStockPatternRepository) SaveSearchCondition(pattern entity.SearchStockPattern) error {
	return s.db.Create(&pattern).Error
}

func (s SearchStockPatternRepository) FindByID(id string) (entity.SearchStockPattern, error) {
	var pattern entity.SearchStockPattern
	err := s.db.Where("search_stock_pattern_id = ?", id).First(&pattern).Error
	if err != nil {
		return entity.SearchStockPattern{}, err
	}
	return pattern, nil
}
