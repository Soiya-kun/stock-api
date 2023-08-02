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

type VolumePatternCreate struct {
	VolumePoint *float64
	IsOver      *bool
	IsMatchRank bool
}

type VolumePatternsCreate []VolumePatternCreate

type PricePatternCreate struct {
	ClosedPoint            *float64
	IsClosedPointOver      *bool
	IsClosedPointMatchRank bool
	OpenedPoint            *float64
	IsOpenedPointOver      *bool
	IsOpenedPointMatchRank bool
	HighPoint              *float64
	IsHighPointOver        *bool
	IsHighPointMatchRank   bool
	LowPoint               *float64
	IsLowPointOver         *bool
	IsLowPointMatchRank    bool
}

type PricePatternsCreate []PricePatternCreate

type MaXUpDownPatternCreate struct {
	MaX     int
	Pattern []bool // false, false, true, true, ...
}

type SearchConditionCreate struct {
	VolumePatterns    VolumePatternsCreate
	PricePatterns     PricePatternsCreate
	MaXUpDownPatterns []MaXUpDownPatternCreate
}

type SearchReq struct {
	SearchPatternID string
	EndDate         time.Time
}

type IStockUseCase interface {
	CreateStocks([]StockCreate) (entity.StocksWithSplits, error)
	FindByStockCode(string) (entity.StocksWithSplits, error)
	FindByRandom() (entity.StocksWithSplits, error)
	SaveStockCode(sc string, u entity.User) error
	ListSC() ([]string, error)
	CreateStockSplit(StockSplitCreate) error
	SaveSearchCondition(SearchConditionCreate, entity.User) error
	SearchByCondition(SearchReq) ([]string, error) // return stockCode
}
