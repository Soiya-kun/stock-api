package port

import (
	"encoding/csv"
	"fmt"

	"gitlab.com/soy-app/stock-api/domain/entity"
)

var (
	ErrCreateStock = fmt.Errorf("failed to create stock")
)

type StockRepository interface {
	Create([]entity.Stock) error
	ReadCSV(reader *csv.Reader) ([]entity.Stock, error)
	FindByStockCode(string) ([]entity.Stock, error)
}
