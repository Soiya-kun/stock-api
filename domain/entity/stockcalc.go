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

func (s *StocksCalc) AverageVolume(start, end int) float64 {
	stocksRanged := s.getByIdxRange(start, end)
	var sum float64
	for _, stock := range stocksRanged.Stocks {
		sum += stock.VolumeVal()
	}
	return sum / float64(len(s.Stocks))
}

func (s *StocksCalc) MaxVolume(start, end int) float64 {
	stocksRanged := s.getByIdxRange(start, end)
	var max float64
	for _, stock := range stocksRanged.Stocks {
		if max < stock.VolumeVal() {
			max = stock.VolumeVal()
		}
	}
	return max
}

// MapStocks
// 1: 1日目のstartPrice
// 2: 1日目のHigh
// 3: 1日目のLow
// 4: 1日目のendPrice
// 5: 2日目のstartPrice
// 6: 2日目のHigh
// ...
func (s *StocksCalc) MapStocks(start, end int) map[int]float64 {
	stocksRanged := s.getByIdxRange(start, end)
	m := make(map[int]float64)
	for i, stock := range stocksRanged.Stocks {
		m[i*4+1] = stock.OpenedPriceVal()
		m[i*4+2] = stock.HighVal()
		m[i*4+3] = stock.LowVal()
		m[i*4+4] = stock.PriceVal()
	}
	return m
}
