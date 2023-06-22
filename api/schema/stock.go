package schema

type StockRes struct {
	StockCode            string  `json:"stockCode"`
	StockName            string  `json:"stockName"`
	Market               string  `json:"market"`
	Industry             string  `json:"industry"`
	Date                 string  `json:"date"`
	ClosedPrice          float64 `json:"closedPrice"`
	OpenedPrice          float64 `json:"openedPrice"`
	HighPrice            float64 `json:"highPrice"`
	LowPrice             float64 `json:"lowPrice"`
	Volume               float64 `json:"volume"`
	TransactionPrice     float64 `json:"transactionPrice"`
	MarketCapitalization float64 `json:"marketCapitalization"`
	LowLimit             float64 `json:"lowLimit"`
	HighLimit            float64 `json:"highLimit"`
}

type StockCodeListRes struct {
	StockCodes []string `json:"stockCodes"`
}
