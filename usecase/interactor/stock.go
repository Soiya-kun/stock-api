package interactor

import (
	"fmt"
	"gitlab.com/soy-app/stock-api/domain/constructor"
	"gitlab.com/soy-app/stock-api/domain/entity"
	"gitlab.com/soy-app/stock-api/usecase/port"
)

type StockUseCase struct {
	ulid                           port.ULID
	stockRepo                      port.StockRepository
	searchStockPatternRepo         port.SearchStockPatternRepository
	searchedStockPatternRepository port.SearchedStockPatternRepository
}

func NewStockUseCase(
	ulid port.ULID,
	stockRepo port.StockRepository,
	searchStockPatternRepo port.SearchStockPatternRepository,
	searchedStockPatternRepository port.SearchedStockPatternRepository,
) IStockUseCase {
	return &StockUseCase{
		ulid:                           ulid,
		stockRepo:                      stockRepo,
		searchStockPatternRepo:         searchStockPatternRepo,
		searchedStockPatternRepository: searchedStockPatternRepository,
	}
}

func (s StockUseCase) CreateStocks(creates []StockCreate) (entity.StocksWithSplits, error) {
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
	err := s.stockRepo.Create(entity.StocksWithSplits{StockList: stocks})
	if err != nil {
		return entity.StocksWithSplits{}, err
	}
	return entity.StocksWithSplits{
		StockList: stocks,
	}, nil
}

func (s StockUseCase) FindByStockCode(sc string) (entity.StocksWithSplits, error) {
	res, err := s.stockRepo.FindByStockCode(sc)
	if err != nil {
		return entity.StocksWithSplits{}, err
	}

	splits, err := s.stockRepo.FindStockSplitsByStockCode(sc)
	if err != nil {
		return entity.StocksWithSplits{}, err
	}

	return entity.StocksWithSplits{
		StockList:   res.Stocks(),
		StockSplits: splits,
	}, nil
}

func (s StockUseCase) FindByRandom() (entity.StocksWithSplits, error) {
	sc, err := s.stockRepo.FindRandomSC()
	if err != nil {
		return entity.StocksWithSplits{}, err
	}

	res, err := s.stockRepo.FindByStockCode(sc)
	if err != nil {
		return entity.StocksWithSplits{}, err
	}

	splits, err := s.stockRepo.FindStockSplitsByStockCode(sc)
	if err != nil {
		return entity.StocksWithSplits{}, err
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
	return s.searchStockPatternRepo.SaveSearchCondition(SearchCondition)
}

func (s StockUseCase) SearchByCondition(req SearchReq) ([]string, error) {
	condition, err := s.searchStockPatternRepo.FindByID(req.SearchPatternID)
	if err != nil {
		return nil, err
	}

	codes, err := s.stockRepo.ListSC()
	if err != nil {
		return nil, err
	}

	var retCodes []string

	for _, c := range codes {
		res, err := s.stockRepo.FindByStockCode(c)
		if err != nil {
			return nil, err
		}

		isMatched := condition.IsMatchPricePatterns(res.StockList.StocksCalc())
		if isMatched {
			fmt.Println("matched")
			retCodes = append(retCodes, c)
		}
	}

	searchedStockPatternID := s.ulid.New()
	if err := s.searchedStockPatternRepository.Create(
		entity.SearchedStockPattern{
			SearchedStockPatternID: searchedStockPatternID,
			SearchStockPatternID:   req.SearchPatternID,
			SearchedStockPatternCodes: func() []*entity.SearchedStockPatternCode {
				ret := make([]*entity.SearchedStockPatternCode, len(retCodes))
				for i, c := range retCodes {
					ret[i] = &entity.SearchedStockPatternCode{
						SearchedStockPatternID: searchedStockPatternID,
						StockCode:              c,
					}
				}
				return ret
			}(),
			EndDate: req.EndDate,
		}); err != nil {
		return nil, err
	}

	return retCodes, nil
}
