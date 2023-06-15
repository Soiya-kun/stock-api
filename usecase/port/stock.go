package port

import (
	"encoding/csv"

	"gitlab.com/soy-app/stock-api/domain/entity"
)

type Stock interface {
	Stock() *entity.Stock
}

type StockList interface {
	Stocks() []*entity.Stock
}

type StockRepository interface {
	Create(list entity.StockList) error
	ReadCSV(reader *csv.Reader) (entity.StockList, error)
	FindByStockCode(string) (entity.StockList, error)
	FindRandomSC() (string, error)
	SaveStockCode(sc, userID string) error
	ListSC() ([]string, error)
	CreateStockSplit(entity.StockSplit) error
	FindStockSplitsByStockCode(string) ([]entity.StockSplit, error)
}
