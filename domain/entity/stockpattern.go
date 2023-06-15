package entity

type StockPattern struct {
	MaXUpDownPattern map[int][]bool // key:Maの日数 true: 上昇, false: 下降
}

// IsMatchMaXUpDownPattern
// StockPatternのMaXUpDownPatternに一致するかどうか
func (s *StockPattern) IsMatchMaXUpDownPattern(stocks []*StockCalc) bool {
	for i, pattern := range s.MaXUpDownPattern {
		verifiedStocks := stocks[len(stocks)-len(pattern)-1:]
		for j := range verifiedStocks {
			if j == 0 {
				continue
			}
			if (verifiedStocks[j].Ma[i] > verifiedStocks[j-1].Ma[i]) != pattern[j-1] {
				return false
			}
		}
	}
	return true
}
