package entity

import "time"

type StockSplit struct {
	StockCode  string    `gorm:"primaryKey;size:30;not null"` // "SC"
	Date       time.Time `gorm:"primaryKey;not null"`         // "日付"
	SplitRatio float64   // "分割"
}
