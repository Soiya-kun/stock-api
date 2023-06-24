package entity

type VolumePatterns []*VolumePattern
type PricePatterns []*PricePattern

type SearchStockPattern struct {
	SearchStockPatternID string `gorm:"primaryKey;size:30;not null"`
	UserID               string
	User                 User `gorm:"constraint:OnDelete:CASCADE"`
	// Volume
	VolumePatterns VolumePatterns
	// ClosedPoint
	PricePatterns PricePatterns
	// Ma
	MaXUpDownPatterns []*MaXUpDownPattern
}

// IsMatchMaXUpDownPattern
// StockPatternのMaXUpDownPatternに一致するかどうか
func (s *SearchStockPattern) IsMatchMaXUpDownPattern(sc StocksCalc) bool {
	for i, p := range s.MaXUpDownPatterns {
		verifiedStocks := sc.Stocks[len(sc.Stocks)-len(p.Pattern)-1:]
		for j := range verifiedStocks {
			if j == 0 {
				continue
			}
			if (verifiedStocks[j].Ma[i] > verifiedStocks[j-1].Ma[i]) != p.IndexBool(j-1) {
				return false
			}
		}
	}
	return true
}
