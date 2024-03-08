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
	VolumePatterns []struct {
		VolumePoint *float64 `json:"volumePoint"`
		IsOver      *bool    `json:"isOver"`
		IsMatchRank bool     `json:"isMatchRank"`
	} `json:"volumePatterns"`

	PricePatterns []struct {
		ClosedPoint            *float64 `json:"closedPoint"`
		IsClosedPointOver      *bool    `json:"isClosedPointOver"`
		IsClosedPointMatchRank bool     `json:"isClosedPointMatchRank"`
		OpenedPoint            *float64 `json:"openedPoint"`
		IsOpenedPointOver      *bool    `json:"isOpenedPointOver"`
		IsOpenedPointMatchRank bool     `json:"isOpenedPointMatchRank"`
		HighPoint              *float64 `json:"highPoint"`
		IsHighPointOver        *bool    `json:"isHighPointOver"`
		IsHighPointMatchRank   bool     `json:"isHighPointMatchRank"`
		LowPoint               *float64 `json:"lowPoint"`
		IsLowPointOver         *bool    `json:"isLowPointOver"`
		IsLowPointMatchRank    bool     `json:"isLowPointMatchRank"`
	} `json:"pricePatterns"`

	MaXUpDownPatterns []struct {
		MaX     int    `json:"maX"`
		Pattern []bool `json:"pattern"`
	} `json:"maXUpDownPatterns"`
}

type StockCodeByThresholdReq struct {
	MinTradeValue int       `json:"minTradeValue"`
	Date          time.Time `json:"date"`
}
