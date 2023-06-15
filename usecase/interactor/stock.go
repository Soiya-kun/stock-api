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

func (s StockUseCase) CreateStocks(creates []StockCreate) (entity.StockList, error) {
	stocks := make([]*entity.Stock, len(creates))
	for i, c := range creates {
		stocks[i] = &entity.Stock{
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
	err := s.stockRepo.Create(entity.StockList{StockList: stocks})
	if err != nil {
		return entity.StockList{}, err
	}
	return entity.StockList{
		StockList: stocks,
	}, nil
}

func (s StockUseCase) FindByStockCode(sc string) (entity.StockList, error) {
	res, err := s.stockRepo.FindByStockCode(sc)
	if err != nil {
		return entity.StockList{}, err
	}

	splits, err := s.stockRepo.FindStockSplitsByStockCode(sc)
	if err != nil {
		return entity.StockList{}, err
	}

	return entity.StockList{
		StockList:   res.Stocks(),
		StockSplits: splits,
	}, nil
}

func (s StockUseCase) FindByRandom() (entity.StockList, error) {
	sc, err := s.stockRepo.FindRandomSC()
	if err != nil {
		return entity.StockList{}, err
	}

	res, err := s.stockRepo.FindByStockCode(sc)
	if err != nil {
		return entity.StockList{}, err
	}

	splits, err := s.stockRepo.FindStockSplitsByStockCode(sc)
	if err != nil {
		return entity.StockList{}, err
	}

	res.StockSplits = splits
	return res, nil
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
