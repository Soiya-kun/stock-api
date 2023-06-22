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

type StockSplitCreate struct {
	StockCode  string
	Date       time.Time
	SplitRatio float64
}

type MaxVolumeInDaysIsOverAverageCreate struct {
	Day         int
	OverAverage float64
}

type PricePatternCreate struct {
	PriceRank       *int // "終値"の順位
	OpenedPriceRank *int // "始値"の順位
	HighRank        *int // "高値"の順位
	LowRank         *int // "安値"の順位
}

type MaXUpDownPatternCreate struct {
	MaX     int
	Pattern []bool // false, false, true, true, ...
}

type SearchConditionCreate struct {
	MaxVolumeInDaysIsOverAverage MaxVolumeInDaysIsOverAverageCreate
	PricePatterns                []PricePatternCreate
	MaXUpDownPatterns            []MaXUpDownPatternCreate
}

type IStockUseCase interface {
	CreateStocks([]StockCreate) (entity.StockList, error)
	FindByStockCode(string) (entity.StockList, error)
	FindByRandom() (entity.StockList, error)
	SaveStockCode(sc string, u entity.User) error
	ListSC() ([]string, error)
	CreateStockSplit(StockSplitCreate) error
	SaveSearchCondition(SearchConditionCreate, entity.User) error
}
