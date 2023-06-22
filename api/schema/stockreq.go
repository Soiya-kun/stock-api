package schema

import "time"

type StockCreateReq struct {
	StockCode     string    `json:"stockCode"`
	StockName     string    `json:"stockName"`
	Market        string    `json:"market"`
	Industry      string    `json:"industry"`
	Date          time.Time `json:"date"`
	Price         *float64  `json:"price"`
	Change        *float64  `json:"change"`
	ChangePercent *float64  `json:"changePercent"`
	PreviousClose *float64  `json:"previousClose"`
	Open          *float64  `json:"open"`
	High          *float64  `json:"high"`
	Low           *float64  `json:"low"`
	Volume        *float64  `json:"volume"`
	TradingValue  *float64  `json:"tradingValue"`
	MarketCap     *float64  `json:"marketCap"`
	LowerLimit    *float64  `json:"lowerLimit"`
	UpperLimit    *float64  `json:"upperLimit"`
}

type StockCreateListReq struct {
	Stocks []StockCreateReq `json:"stocks"`
}

type SaveSCReq struct {
	StockCode string `json:"stockCode"`
}

type StockSplitReq struct {
	StockCode  string    `json:"stockCode"`
	Date       time.Time `json:"date"`
	SplitRatio float64   `json:"splitRatio"`
}

type SaveSearchConditionReq struct {
	MaxVolumeInDaysIsOverAverage struct {
		Day         int     `json:"day"`
		OverAverage float64 `json:"overAverage"`
	} `json:"maxVolumeInDaysIsOverAverage"`

	PricePatterns []struct {
		PriceRank       *int `json:"priceRank"`
		OpenedPriceRank *int `json:"openedPriceRank"`
		HighRank        *int `json:"highRank"`
		LowRank         *int `json:"lowRank"`
	} `json:"pricePatterns"`

	MaXUpDownPatterns []struct {
		MaX     int    `json:"MaX"`
		Pattern []bool `json:"pattern"`
	} `json:"maXUpDownPatterns"`
}
