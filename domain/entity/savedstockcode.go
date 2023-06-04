package entity

type SavedStockCode struct {
	StockCode string `gorm:"primaryKey;size:30;not null"` // "SC"
	UserID    string `gorm:"primaryKey;size:30;not null"`
	User      User
}
