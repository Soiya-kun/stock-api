package schema

import "gitlab.com/soy-app/stock-api/usecase/interactor"

func (s *SaveSearchConditionReq) UsecaseArg() interactor.SearchConditionCreate {
	return func() interactor.SearchConditionCreate {
		return interactor.SearchConditionCreate{
			MaxVolumeInDaysIsOverAverage: interactor.MaxVolumeInDaysIsOverAverageCreate{
				Day:         s.MaxVolumeInDaysIsOverAverage.Day,
				OverAverage: s.MaxVolumeInDaysIsOverAverage.OverAverage,
			},
			PricePatterns: func() []interactor.PricePatternCreate {
				ret := make([]interactor.PricePatternCreate, len(s.PricePatterns))
				for i, v := range s.PricePatterns {
					ret[i] = interactor.PricePatternCreate{
						PriceRank:       v.PriceRank,
						OpenedPriceRank: v.OpenedPriceRank,
						HighRank:        v.HighRank,
						LowRank:         v.LowRank,
					}
				}
				return ret
			}(),
			MaXUpDownPatterns: func() []interactor.MaXUpDownPatternCreate {
				ret := make([]interactor.MaXUpDownPatternCreate, len(s.MaXUpDownPatterns))
				for i, v := range s.MaXUpDownPatterns {
					ret[i] = interactor.MaXUpDownPatternCreate{
						MaX:     v.MaX,
						Pattern: v.Pattern,
					}
				}
				return ret
			}(),
		}
	}()
}
