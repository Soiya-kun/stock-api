package interactor

import (
	"gitlab.com/soy-app/stock-api/domain/entity"
	"gitlab.com/soy-app/stock-api/usecase/port"
)

type StockUseCase struct {
	stockRepo port.StockRepository
}

func NewStockUseCase(stockRepo port.StockRepository) IStockUseCase {
	return &StockUseCase{stockRepo: stockRepo}
}

func (s StockUseCase) CreateStocks(creates []StockCreate) ([]entity.Stock, error) {
	stocks := make([]entity.Stock, len(creates))
	for i, c := range creates {
		stocks[i] = entity.Stock{
			StockCode:     c.StockCode,
			StockName:     c.StockName,
			Market:        c.Market,
			Industry:      c.Industry,
			Date:          c.Date,
			Price:         c.Price,
			Change:        c.Change,
			ChangePercent: c.ChangePercent,
			PreviousClose: c.PreviousClose,
			OpenedPrice:   c.Open,
			High:          c.High,
			Low:           c.Low,
			Volume:        c.Volume,
			TradingValue:  c.TradingValue,
			MarketCap:     c.MarketCap,
			LowerLimit:    c.LowerLimit,
			UpperLimit:    c.UpperLimit,
		}
	}
	err := s.stockRepo.Create(stocks)
	if err != nil {
		return nil, err
	}
	return stocks, nil
}

func (s StockUseCase) FindByStockCode(sc string) ([]entity.Stock, error) {
	return s.stockRepo.FindByStockCode(sc)
}
