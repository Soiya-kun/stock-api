package constructor

import (
	"gitlab.com/soy-app/stock-api/domain/entity"
)

func NewSearchStockPatternCreate(
	ulid string,
	userID string,
	volumePatterns entity.VolumePatterns,
	pricePatterns entity.PricePatterns,
	MaXUpDownPatterns []*entity.MaXUpDownPattern,
) entity.SearchStockPattern {
	return entity.SearchStockPattern{
		SearchStockPatternID: ulid,
		UserID:               userID,
		VolumePatterns:       volumePatterns,
		PricePatterns:        pricePatterns,
		MaXUpDownPatterns:    MaXUpDownPatterns,
	}
}
