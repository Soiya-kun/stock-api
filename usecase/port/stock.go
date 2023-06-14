package port

import (
	"encoding/csv"

	"gitlab.com/soy-app/stock-api/domain/entity"
)

type StockRepository interface {
	Create([]entity.Stock) error
	ReadCSV(reader *csv.Reader) ([]entity.Stock, error)
	FindByStockCode(string) ([]*entity.Stock, error)
	FindRandomSC() (string, error)
	SaveStockCode(sc, userID string) error
	ListSC() ([]string, error)
	CreateStockSplit(entity.StockSplit) error
	FindStockSplitsByStockCode(string) ([]entity.StockSplit, error)
}
