package schema

import (
	"time"

	"gitlab.com/soy-app/stock-api/domain/entity"
	"gitlab.com/soy-app/stock-api/usecase/interactor"
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

type StockCreateListReq struct {
	Stocks []StockCreate `json:"stocks"`
}

func (r *StockCreateListReq) StockCreate() []interactor.StockCreate {
	ret := make([]interactor.StockCreate, len(r.Stocks))
	for i, v := range r.Stocks {
		ret[i] = interactor.StockCreate{
			StockCode:     v.StockCode,
			StockName:     v.StockName,
			Market:        v.Market,
			Industry:      v.Industry,
			Date:          v.Date,
			Price:         v.Price,
			Change:        v.Change,
			ChangePercent: v.ChangePercent,
			PreviousClose: v.PreviousClose,
			Open:          v.Open,
			High:          v.High,
			Low:           v.Low,
			Volume:        v.Volume,
			TradingValue:  v.TradingValue,
			MarketCap:     v.MarketCap,
			LowerLimit:    v.LowerLimit,
			UpperLimit:    v.UpperLimit,
		}
	}
	return ret
}

func StockListResFromEntity(stocks []entity.Stock) []StockCreate {
	ret := make([]StockCreate, len(stocks))
	for i, v := range stocks {
		ret[i] = StockCreate{
			StockCode:     v.StockCode,
			StockName:     v.StockName,
			Market:        v.Market,
			Industry:      v.Industry,
			Date:          v.Date,
			Price:         v.Price,
			Change:        v.Change,
			ChangePercent: v.ChangePercent,
			PreviousClose: v.PreviousClose,
			Open:          v.Open,
			High:          v.High,
			Low:           v.Low,
			Volume:        v.Volume,
			TradingValue:  v.TradingValue,
			MarketCap:     v.MarketCap,
			LowerLimit:    v.LowerLimit,
			UpperLimit:    v.UpperLimit,
		}
	}
	return ret
}
