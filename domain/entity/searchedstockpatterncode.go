package entity

type SearchedStockPatternCode struct {
	SearchedStockPatternID string `gorm:"primaryKey;size:26;not null"`
	SearchedStockPattern   SearchedStockPattern
	StockCode              string `gorm:"primaryKey"`
}
