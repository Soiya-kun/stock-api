package schema

import (
	"time"

	"gitlab.com/soy-app/stock-api/usecase/port"

	"gitlab.com/soy-app/stock-api/usecase/interactor"
)

type StockCreate struct {
	StockCode     string    `json:"stockCode"`
	StockName     string    `json:"stockName"`
	Market        string    `json:"market"`
	Industry      string    `json:"industry"`
	Date          time.Time `json:"date"`
	Price         *float64  `json:"price"`
	Change        *float64  `json:"change"`
	ChangePercent *float64  `json:"changePercent"`
	PreviousClose *float64  `json:"previousClose"`
	Open          *float64  `json:"open"`
	High          *float64  `json:"high"`
	Low           *float64  `json:"low"`
	Volume        *float64  `json:"volume"`
	TradingValue  *float64  `json:"tradingValue"`
	MarketCap     *float64  `json:"marketCap"`
	LowerLimit    *float64  `json:"lowerLimit"`
	UpperLimit    *float64  `json:"upperLimit"`
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

type StockRes struct {
	StockCode            string  `json:"stockCode"`
	StockName            string  `json:"stockName"`
	Market               string  `json:"market"`
	Industry             string  `json:"industry"`
	Date                 string  `json:"date"`
	ClosedPrice          float64 `json:"closedPrice"`
	OpenedPrice          float64 `json:"openedPrice"`
	HighPrice            float64 `json:"highPrice"`
	LowPrice             float64 `json:"lowPrice"`
	Volume               float64 `json:"volume"`
	TransactionPrice     float64 `json:"transactionPrice"`
	MarketCapitalization float64 `json:"marketCapitalization"`
	LowLimit             float64 `json:"lowLimit"`
	HighLimit            float64 `json:"highLimit"`
}

func StockResFromEntity(s port.Stock) StockRes {
	return StockRes{
		StockCode:            s.Stock().StockCode,
		StockName:            s.Stock().StockName,
		Market:               s.Stock().Market,
		Industry:             s.Stock().Industry,
		Date:                 s.Stock().Date.Format("2006-01-02"),
		ClosedPrice:          s.Stock().PriceVal(),
		OpenedPrice:          s.Stock().OpenedPriceVal(),
		HighPrice:            s.Stock().HighVal(),
		LowPrice:             s.Stock().LowVal(),
		Volume:               s.Stock().VolumeVal(),
		TransactionPrice:     s.Stock().TradingValueVal(),
		MarketCapitalization: s.Stock().MarketCapVal(),
		LowLimit:             s.Stock().LowerLimitVal(),
		HighLimit:            s.Stock().UpperLimitVal(),
	}
}

func StocksResFromEntity(s port.StockList) []StockRes {
	ret := make([]StockRes, len(s.Stocks()))
	for i, v := range s.Stocks() {
		ret[i] = StockResFromEntity(v)
	}
	return ret
}

type SaveSCReq struct {
	StockCode string `json:"stockCode"`
}

type StockCodeListRes struct {
	StockCodes []string `json:"stockCodes"`
}

type StockSplitReq struct {
	StockCode  string    `json:"stockCode"`
	Date       time.Time `json:"date"`
	SplitRatio float64   `json:"splitRatio"`
}
