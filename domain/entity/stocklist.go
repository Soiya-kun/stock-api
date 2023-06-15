package entity

type StockList struct {
	StockList   []*Stock
	StockSplits []StockSplit
}

// Stocks
// 株式分割情報を反映した株価
func (s *StockList) Stocks() []*Stock {
	var ret []*Stock
	copy(ret, s.StockList)
	for _, split := range s.StockSplits {
		for _, stock := range ret {
			stock.Stock().StockAfterApplyingSplit(split)
		}
	}
	return ret
}
