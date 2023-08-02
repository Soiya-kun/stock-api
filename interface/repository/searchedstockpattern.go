package repository

import (
	"time"

	"gorm.io/gorm"

	"gitlab.com/soy-app/stock-api/domain/entity"
	"gitlab.com/soy-app/stock-api/usecase/port"
)

type SearchedStockPatternRepository struct {
	db *gorm.DB
}

func NewSearchedStockPatternRepository(db *gorm.DB) port.SearchedStockPatternRepository {
	return &SearchedStockPatternRepository{
		db: db,
	}
}

func (s SearchedStockPatternRepository) Create(searchedPattern entity.SearchedStockPattern) error {
	return s.db.Create(&entity.SearchedStockPattern{
		SearchedStockPatternID: searchedPattern.SearchedStockPatternID,
		SearchStockPatternID:   searchedPattern.SearchStockPatternID,
		EndDate:                searchedPattern.EndDate,
	}).Error
}

func (s SearchedStockPatternRepository) FindBySearchStockPatternIDAndEndDate(SearchStockPatternID string, endDate time.Time) (entity.SearchedStockPattern, error) {
	var searchedPattern entity.SearchedStockPattern
	err := s.db.Where("search_stock_pattern_id = ? AND end_date = ?", SearchStockPatternID, endDate).
		First(&searchedPattern).
		Error
	if err != nil {
		return entity.SearchedStockPattern{}, err
	}
	return searchedPattern, nil
}
