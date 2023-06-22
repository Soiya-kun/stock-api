package entity

type MaxVolumeInDaysIsOverAverage struct {
	MaxVolumeInDaysIsOverAverageID string `gorm:"primaryKey;size:30;not null"`
	SearchStockPatternID           string
	Day                            int     // N日間
	RatioOverAverage               float64 // 平均出来高の何倍か 1.5, 3.0...
}

type PricePattern struct {
	PricePatternID       string `gorm:"primaryKey;size:30;not null"`
	SearchStockPatternID string
	SearchStockPattern   SearchStockPattern
	PriceRank            *int // "終値"の順位
	OpenedPriceRank      *int // "始値"の順位
	HighRank             *int // "高値"の順位
	LowRank              *int // "安値"の順位
}

type MaXUpDownPattern struct {
	MaXUpDownPatternID   string `gorm:"primaryKey;size:30;not null"`
	SearchStockPatternID string
	SearchStockPattern   SearchStockPattern
	MaX                  int
	Pattern              string `gorm:"size:255;not null"` // 00110101001111 -> false, false, true, true, ...
}

func (m *MaXUpDownPattern) IndexBool(idx int) bool {
	s := m.Pattern[idx : idx+1]
	return s == "1"
}
