package port

import (
	"fmt"

	"gitlab.com/soy-app/stock-api/domain/entity"
)

var (
	ErrCreateStock = fmt.Errorf("failed to create stock")
)

type StockRepository interface {
	Create([]entity.Stock) error
}
