package repository

import (
	"fmt"
	"gorm.io/gorm"
	"sort"

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
	err := s.db.
		Preload("VolumePatterns").
		Preload("PricePatterns").
		Preload("MaXUpDownPatterns").
		Where("search_stock_pattern_id = ?", id).
		First(&pattern).
		Order("arr_index desc").
		Error
	sort.Slice(pattern.PricePatterns, func(i, j int) bool {
		return pattern.PricePatterns[i].ArrIndex < pattern.PricePatterns[j].ArrIndex
	})
	for i := range pattern.PricePatterns {
		fmt.Println(pattern.PricePatterns[i].ArrIndex)
	}
	if err != nil {
		return entity.SearchStockPattern{}, err
	}
	return pattern, nil
}
