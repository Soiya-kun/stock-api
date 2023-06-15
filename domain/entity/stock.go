package entity

import "time"

type Stock struct {
	StockCode     string    `gorm:"primaryKey;size:30;not null"` // "SC"
	StockName     string    `gorm:"size:255;not null"`           // "名称"
	Market        string    `gorm:"size:255;not null"`           // "市場"
	Industry      string    `gorm:"size:255;not null"`           // "業種"
	Date          time.Time `gorm:"primaryKey;not null"`         // "日付"
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

func (s *Stock) Stock() *Stock {
	return s
}

func (s *Stock) PriceVal() float64 {
	if s.Price == nil {
		return 0
	}
	return *s.Price
}

func (s *Stock) ChangeVal() float64 {
	if s.Change == nil {
		return 0
	}
	return *s.Change
}

func (s *Stock) ChangePercentVal() float64 {
	if s.ChangePercent == nil {
		return 0
	}
	return *s.ChangePercent
}

func (s *Stock) PreviousCloseVal() float64 {
	if s.PreviousClose == nil {
		return 0
	}
	return *s.PreviousClose
}

func (s *Stock) OpenedPriceVal() float64 {
	if s.OpenedPrice == nil {
		return 0
	}
	return *s.OpenedPrice
}

func (s *Stock) HighVal() float64 {
	if s.High == nil {
		return 0
	}
	return *s.High
}

func (s *Stock) LowVal() float64 {
	if s.Low == nil {
		return 0
	}
	return *s.Low
}

func (s *Stock) VolumeVal() float64 {
	if s.Volume == nil {
		return 0
	}
	return *s.Volume
}

func (s *Stock) TradingValueVal() float64 {
	if s.TradingValue == nil {
		return 0
	}
	return *s.TradingValue
}

func (s *Stock) MarketCapVal() float64 {
	if s.MarketCap == nil {
		return 0
	}
	return *s.MarketCap
}

func (s *Stock) LowerLimitVal() float64 {
	if s.LowerLimit == nil {
		return 0
	}
	return *s.LowerLimit
}

func (s *Stock) UpperLimitVal() float64 {
	if s.UpperLimit == nil {
		return 0
	}
	return *s.UpperLimit
}

func (s *Stock) StockAfterApplyingSplit(split StockSplit) {
	if s.Date.After(split.Date) {
		return
	}
	if s.OpenedPrice != nil {
		*s.OpenedPrice /= split.SplitRatio
	}
	if s.High != nil {
		*s.High /= split.SplitRatio
	}
	if s.Low != nil {
		*s.Low /= split.SplitRatio
	}
	if s.Price != nil {
		*s.Price /= split.SplitRatio
	}
	if s.PreviousClose != nil {
		*s.PreviousClose /= split.SplitRatio
	}
	if s.LowerLimit != nil {
		*s.LowerLimit /= split.SplitRatio
	}
	if s.UpperLimit != nil {
		*s.UpperLimit /= split.SplitRatio
	}
}
