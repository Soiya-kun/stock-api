package entity

type StockList struct {
	Stocks      []*Stock
	StockSplits []StockSplit
}

func (s StockList) GetStocksAfterApplyingSplit() []*Stock {
	for _, stock := range s.Stocks {
		for _, split := range s.StockSplits {
			stock.StockAfterApplyingSplit(split)
		}
	}
	return s.Stocks
}
