package constructor

import (
	"gitlab.com/soy-app/stock-api/domain/entity"
	"gitlab.com/soy-app/stock-api/usecase/port"
)

func NewVolumePattern(
	ulid string,
	ArrIndex int,
	VolumePoint *float64,
	IsOver *bool,
	IsMatchRank bool,
) *entity.VolumePattern {
	return &entity.VolumePattern{
		VolumePatternID: ulid,
		ArrIndex:        ArrIndex,
		VolumePoint:     VolumePoint,
		IsOver:          IsOver,
		IsMatchRank:     IsMatchRank,
	}
}

func NewPricePatternCreate(
	ulid string,
	ArrIndex int,
	PricePoint *float64,
	IsOver *bool,
	IsMatchRank bool,
) entity.PricePattern {
	return entity.PricePattern{
		PricePatternID: ulid,
		ArrIndex:       ArrIndex,
		PricePoint:     PricePoint,
		IsOver:         IsOver,
		IsMatchRank:    IsMatchRank,
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
