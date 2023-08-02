package port

import (
	"encoding/csv"

	"gitlab.com/soy-app/stock-api/domain/entity"
)

type Stock interface {
	Stock() *entity.Stock
}

type StockList interface {
	Stocks() entity.Stocks
}

type StockRepository interface {
	Create(list entity.StocksWithSplits) error
	ReadCSV(reader *csv.Reader) (entity.StocksWithSplits, error)
	FindByStockCode(string) (entity.StocksWithSplits, error)
	FindRandomSC() (string, error)
	SaveStockCode(sc, userID string) error
	ListSC() ([]string, error)
	CreateStockSplit(entity.StockSplit) error
	FindStockSplitsByStockCode(string) ([]entity.StockSplit, error)
}
