package entity

type Stocks []*Stock

func (s Stocks) StocksCalc() StocksCalc {
	stocksCalc := make([]*StockCalc, len(s))
	for _, stock := range s {
		stocksCalc = append(stocksCalc, &StockCalc{
			Stock: *stock,
			Ma:    make(map[int]float64),
		})
	}
	return StocksCalc{
		Stocks: stocksCalc,
	}
}
