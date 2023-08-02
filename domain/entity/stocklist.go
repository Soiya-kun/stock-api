package entity

type StocksWithSplits struct {
	StockList   Stocks
	StockSplits []StockSplit
}

// Stocks
// 株式分割情報を反映した株価
func (s *StocksWithSplits) Stocks() Stocks {
	ret := make([]*Stock, len(s.StockList))
	copy(ret, s.StockList)
	for _, split := range s.StockSplits {
		for _, stock := range ret {
			stock.Stock().StockAfterApplyingSplit(split)
		}
	}
	return ret
}
