package port

import (
	"time"

	"gitlab.com/soy-app/stock-api/domain/entity"
)

type SearchedStockPatternRepository interface {
	Create(pattern entity.SearchedStockPattern) error
	FindBySearchStockPatternIDAndEndDate(
		SearchStockPatternID string, endDate time.Time,
	) (entity.SearchedStockPattern, error)
}
