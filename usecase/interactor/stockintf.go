package interactor

import (
	"time"

	"gitlab.com/soy-app/stock-api/domain/entity"
)

type StockCreate struct {
	StockCode     string
	StockName     string
	Market        string
	Industry      string
	Date          time.Time
	Price         *float64
	Change        *float64
	ChangePercent *float64
	PreviousClose *float64
	Open          *float64
	High          *float64
	Low           *float64
	Volume        *float64
	TradingValue  *float64
	MarketCap     *float64
	LowerLimit    *float64
	UpperLimit    *float64
}

type IStockUseCase interface {
	CreateStocks([]StockCreate) ([]entity.Stock, error)
	FindByStockCode(string) ([]entity.Stock, error)
}
