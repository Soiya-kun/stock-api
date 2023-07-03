package interactor

import (
	"gitlab.com/soy-app/stock-api/domain/constructor"
	"gitlab.com/soy-app/stock-api/domain/entity"
	"gitlab.com/soy-app/stock-api/usecase/port"
)

type StockUseCase struct {
	ulid      port.ULID
	stockRepo port.StockRepository
}

func NewStockUseCase(ulid port.ULID, stockRepo port.StockRepository) IStockUseCase {
	return &StockUseCase{ulid: ulid, stockRepo: stockRepo}
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

func (s StockUseCase) SaveSearchCondition(create SearchConditionCreate, user entity.User) error {
	SearchCondition := constructor.NewSearchStockPatternCreate(
		s.ulid.New(),
		user.UserID,
		func() entity.VolumePatterns {
			ret := make(entity.VolumePatterns, len(create.VolumePatterns))
			for i, p := range create.VolumePatterns {
				pEnt := constructor.NewVolumePattern(
					s.ulid.New(),
					i,
					p.VolumePoint,
					p.IsOver,
					p.IsMatchRank,
				)
				ret[i] = pEnt
			}
			return ret
		}(),
		func() []*entity.PricePattern {
			ret := make([]*entity.PricePattern, len(create.PricePatterns)*4)
			for i, p := range create.PricePatterns {
				pOpen := constructor.NewPricePatternCreate(
					s.ulid.New(),
					i*4,
					p.OpenedPoint,
					p.IsClosedPointOver,
					p.IsClosedPointMatchRank,
				)
				ret[i*4] = &pOpen
				pHigh := constructor.NewPricePatternCreate(
					s.ulid.New(),
					i*4+1,
					p.HighPoint,
					p.IsHighPointOver,
					p.IsHighPointMatchRank,
				)
				ret[i*4+1] = &pHigh
				pLow := constructor.NewPricePatternCreate(
					s.ulid.New(),
					i*4+2,
					p.LowPoint,
					p.IsLowPointOver,
					p.IsLowPointMatchRank,
				)
				ret[i*4+2] = &pLow
				pClose := constructor.NewPricePatternCreate(
					s.ulid.New(),
					i*4+3,
					p.ClosedPoint,
					p.IsClosedPointOver,
					p.IsClosedPointMatchRank,
				)
				ret[i*4+3] = &pClose
			}
			return ret
		}(),
		func() []*entity.MaXUpDownPattern {
			ret := make([]*entity.MaXUpDownPattern, len(create.MaXUpDownPatterns))
			for i, p := range create.MaXUpDownPatterns {
				pEnt := constructor.NewMaXUpDownPatternCreate(
					s.ulid,
					p.MaX,
					p.Pattern,
				)
				ret[i] = &pEnt
			}
			return ret
		}(),
	)
	return s.stockRepo.SaveSearchCondition(SearchCondition)
}
