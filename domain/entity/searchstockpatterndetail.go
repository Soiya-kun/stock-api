package entity

type VolumePattern struct {
	VolumePatternID      string `gorm:"primaryKey;size:26;not null"`
	SearchStockPatternID string
	SearchStockPattern   SearchStockPattern
	ArrIndex             int
	VolumePoint          *float64
	IsOver               *bool
	IsMatchRank          bool
}

type PricePattern struct {
	PricePatternID       string `gorm:"primaryKey;size:30;not null"`
	SearchStockPatternID string
	SearchStockPattern   SearchStockPattern
	ArrIndex             int // %4=0: OpenedPrice, %4=1: HighPrice, %4=2: LowPrice, %4=3: ClosedPrice
	PricePoint           *float64
	IsOver               *bool
	IsMatchRank          bool
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
