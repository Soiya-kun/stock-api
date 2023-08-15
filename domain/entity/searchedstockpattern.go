package entity

import "time"

type SearchedStockPattern struct {
	SearchedStockPatternID    string `gorm:"primaryKey;size:26;not null"`
	SearchStockPatternID      string
	SearchStockPattern        SearchStockPattern
	SearchedStockPatternCodes []*SearchedStockPatternCode
	EndDate                   time.Time
}
