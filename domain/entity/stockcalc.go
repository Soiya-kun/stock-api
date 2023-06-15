package entity

type StockCalc struct {
	Stock
	Ma map[int]float64
}

type StocksCalc struct {
	Stocks []*StockCalc
}

func (s *StocksCalc) getByIdxRange(start, end int) StocksCalc {
	return StocksCalc{Stocks: s.Stocks[start:end]}
}

func (s *StocksCalc) AveragePrice() float64 {
	var sum float64
	for _, stock := range s.Stocks {
		sum += stock.PriceVal()
	}
	return sum / float64(len(s.Stocks))
}

func (s *StocksCalc) CalcMA(day int) {
	for i, stock := range s.Stocks {
		if i < day {
			continue
		}
		ss := s.getByIdxRange(i-day, i)
		stock.Ma[day] = ss.AveragePrice()
	}
}
