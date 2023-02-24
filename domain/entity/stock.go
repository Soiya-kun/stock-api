package entity

import "time"

// Stock "SC","名称","市場","業種","日付","株価","前日比","前日比（％）","前日終値","始値","高値","安値","出来高","売買代金（千円）","時価総額（百万円）","値幅下限","値幅上限"
type Stock struct {
	ID            uint      `gorm:"primaryKey;autoIncrement"`
	StockCode     string    `gorm:"primaryKey;size:30;not null"`
	StockName     string    `gorm:"size:255;not null"`
	Market        string    `gorm:"size:255;not null"`
	Industry      string    `gorm:"size:255;not null"`
	Date          time.Time `gorm:"size:255;not null"`
	Price         *float64
	Change        *float64
	ChangePercent *float64
	PreviousClose *float64
	Open          *float64
	High          *float64
	Low           *float64
	Volume        *float64
	TradingValue  *float64
	MarketCap     *float64
	LowerLimit    *float64
	UpperLimit    *float64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
