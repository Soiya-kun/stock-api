package constructor

import (
	"gitlab.com/soy-app/stock-api/domain/entity"
	"gitlab.com/soy-app/stock-api/usecase/port"
)

func NewMaxVolumeInDaysIsOverAverageCreate(
	ulid port.ULID,
	day int,
	ratioOverAverage float64,
) *entity.MaxVolumeInDaysIsOverAverage {
	if day == 0 {
		return nil
	}
	return &entity.MaxVolumeInDaysIsOverAverage{
		MaxVolumeInDaysIsOverAverageID: ulid.New(),
		Day:                            day,
		RatioOverAverage:               ratioOverAverage,
	}
}

func NewPricePatternCreate(
	ulid port.ULID,
	priceRank *int,
	openedPriceRank *int,
	highRank *int,
	lowRank *int,
) entity.PricePattern {
	return entity.PricePattern{
		PricePatternID:  ulid.New(),
		PriceRank:       priceRank,
		OpenedPriceRank: openedPriceRank,
		HighRank:        highRank,
		LowRank:         lowRank,
	}
}

func NewMaXUpDownPatternCreate(
	ulid port.ULID,
	maX int,
	pattern []bool,
) entity.MaXUpDownPattern {
	patternStr := ""
	for _, v := range pattern {
		if v {
			patternStr += "1"
		} else {
			patternStr += "0"
		}
	}
	return entity.MaXUpDownPattern{
		MaXUpDownPatternID: ulid.New(),
		MaX:                maX,
		Pattern:            patternStr,
	}
}
