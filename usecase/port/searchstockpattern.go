package port

import "gitlab.com/soy-app/stock-api/domain/entity"

type SearchStockPatternRepository interface {
	SaveSearchCondition(pattern entity.SearchStockPattern) error
	FindByID(id string) (entity.SearchStockPattern, error)
}
