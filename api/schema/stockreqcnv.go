package schema

import "gitlab.com/soy-app/stock-api/usecase/interactor"

func (s *SaveSearchConditionReq) UsecaseArg() interactor.SearchConditionCreate {
	return func() interactor.SearchConditionCreate {
		return interactor.SearchConditionCreate{
			VolumePatterns: func() interactor.VolumePatternsCreate {
				ret := make(interactor.VolumePatternsCreate, len(s.VolumePatterns))
				for i, v := range s.VolumePatterns {
					ret[i] = interactor.VolumePatternCreate{
						VolumePoint: v.VolumePoint,
						IsOver:      v.IsOver,
						IsMatchRank: v.IsMatchRank,
					}
				}
				return ret
			}(),
			PricePatterns: func() []interactor.PricePatternCreate {
				ret := make([]interactor.PricePatternCreate, len(s.PricePatterns))
				for i, v := range s.PricePatterns {
					ret[i] = interactor.PricePatternCreate{
						ClosedPoint:            v.ClosedPoint,
						IsClosedPointOver:      v.IsClosedPointOver,
						IsClosedPointMatchRank: v.IsClosedPointMatchRank,
						OpenedPoint:            v.OpenedPoint,
						IsOpenedPointOver:      v.IsOpenedPointOver,
						IsOpenedPointMatchRank: v.IsOpenedPointMatchRank,
						HighPoint:              v.HighPoint,
						IsHighPointOver:        v.IsHighPointOver,
						IsHighPointMatchRank:   v.IsHighPointMatchRank,
						LowPoint:               v.LowPoint,
						IsLowPointOver:         v.IsLowPointOver,
						IsLowPointMatchRank:    v.IsLowPointMatchRank,
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
