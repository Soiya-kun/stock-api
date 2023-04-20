package entity

import "time"

type Stock struct {
	ID            uint      `gorm:"primaryKey;autoIncrement"`
	StockCode     string    `gorm:"size:30;not null"`  // "SC"
	StockName     string    `gorm:"size:255;not null"` // "名称"
	Market        string    `gorm:"size:255;not null"` // "市場"
	Industry      string    `gorm:"size:255;not null"` // "業種"
	Date          time.Time `gorm:"size:255;not null"` // "日付"
	Price         *float64  // "株価"
	Change        *float64  // "前日比"
	ChangePercent *float64  // "前日比（％）"
	PreviousClose *float64  // "前日終値"
	OpenedPrice   *float64  // "始値"
	High          *float64  // "高値"
	Low           *float64  // "安値"
	Volume        *float64  // "出来高"
	TradingValue  *float64  // "売買代金（千円）"
	MarketCap     *float64  // "時価総額（百万円）"
	LowerLimit    *float64  // "値幅下限"
	UpperLimit    *float64  // "値幅上限"
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
