package port

import (
	"encoding/csv"

	"gitlab.com/soy-app/stock-api/domain/entity"
)

type StockRepository interface {
	Create([]entity.Stock) error
	ReadCSV(reader *csv.Reader) ([]entity.Stock, error)
	FindByStockCode(string) ([]entity.Stock, error)
	FindRandomSC() (string, error)
}
