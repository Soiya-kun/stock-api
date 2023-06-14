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

func (s StockUseCase) FindByStockCode(sc string) ([]*entity.Stock, error) {
	res, err := s.stockRepo.FindByStockCode(sc)
	if err != nil {
		return nil, err
	}

	splits, err := s.stockRepo.FindStockSplitsByStockCode(sc)
	if err != nil {
		return nil, err
	}
	if len(splits) == 0 {
		return res, nil
	}

	stocks := entity.StockList{
		Stocks:      res,
		StockSplits: splits,
	}
	return stocks.GetStocksAfterApplyingSplit(), nil
}

func (s StockUseCase) FindByRandom() ([]*entity.Stock, error) {
	sc, err := s.stockRepo.FindRandomSC()
	if err != nil {
		return nil, err
	}

	return s.stockRepo.FindByStockCode(sc)
}

func (s StockUseCase) SaveStockCode(sc string, user entity.User) error {
	return s.stockRepo.SaveStockCode(sc, user.UserID)
}

func (s StockUseCase) ListSC() ([]string, error) {
	return s.stockRepo.ListSC()
}

func (s StockUseCase) CreateStockSplit(create StockSplitCreate) error {
	return s.stockRepo.CreateStockSplit(
		entity.StockSplit{
			StockCode:  create.StockCode,
			Date:       create.Date,
			SplitRatio: create.SplitRatio,
		})
}
