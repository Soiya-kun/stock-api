package schema

import (
	"gitlab.com/soy-app/stock-api/usecase/interactor"
	"gitlab.com/soy-app/stock-api/usecase/port"
)

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
	ret := make([]StockRes, s.Len())
	for i, v := range s.Stocks() {
		ret[i] = StockResFromEntity(v)
	}
	return ret
}
