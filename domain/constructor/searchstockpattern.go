package constructor

import (
	"gitlab.com/soy-app/stock-api/domain/entity"
	"gitlab.com/soy-app/stock-api/usecase/port"
)

func NewSearchStockPatternCreate(
	ulid port.ULID,
	userID string,
	MaxVolumeInDaysIsOverAverage *entity.MaxVolumeInDaysIsOverAverage,
	PricePatterns []*entity.PricePattern,
	MaXUpDownPatterns []*entity.MaXUpDownPattern,
) entity.SearchStockPattern {
	return entity.SearchStockPattern{
		SearchStockPatternID:         ulid.New(),
		UserID:                       userID,
		MaxVolumeInDaysIsOverAverage: MaxVolumeInDaysIsOverAverage,
		PricePatterns:                PricePatterns,
		MaXUpDownPatterns:            MaXUpDownPatterns,
	}
}
