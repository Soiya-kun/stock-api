package entity

import (
	"fmt"
	"sort"
)

type StockPattern struct {
	// Volume
	MaxVolumeInDaysIsOverAverage struct {
		Day         int     // N日間
		OverAverage float64 // 平均出来高の何倍か
	} // N日間のうちの最大出来高が平均出来高の何倍か

	// PriceRank
	PricePattern []struct {
		PriceRank       *int // "終値"の順位
		OpenedPriceRank *int // "始値"の順位
		HighRank        *int // "高値"の順位
		LowRank         *int // "安値"の順位
	}

	// Ma
	MaXUpDownPattern map[int][]bool // key:Maの日数 true: 上昇, false: 下降
}

// IsMaxVolumeInDaysOverAverage
// N日間のうちの最大出来高が平均出来高の何倍か
func (s *StockPattern) IsMaxVolumeInDaysOverAverage(sc StocksCalc) bool {
	maxVolume := sc.MaxVolume(
		len(sc.Stocks)-s.MaxVolumeInDaysIsOverAverage.Day,
		len(sc.Stocks)-1,
	)
	return maxVolume > sc.AverageVolume(
		len(sc.Stocks)-s.MaxVolumeInDaysIsOverAverage.Day,
		len(sc.Stocks)-1,
	)*s.MaxVolumeInDaysIsOverAverage.OverAverage
}

// IsMatchPricePattern
// StockPatternのPricePatternに一致するかどうか
// indexedStocksのindexは以下の通り
// 1: 1日目のstartPrice
// 2: 1日目のHigh
// 3: 1日目のLow
// 4: 1日目のendPrice
// 5: 2日目のstartPrice
// 6: 2日目のHigh...
func (s *StockPattern) IsMatchPricePattern(sc StocksCalc) bool {
	type indexedStock struct {
		price float64
		index int
	}
	indexedStocks := make([]indexedStock, len(s.PricePattern)*4)
	for i, s := range sc.getByIdxRange(
		len(sc.Stocks)-len(s.PricePattern),
		len(sc.Stocks)).Stocks {
		indexedStocks[i*4] = indexedStock{price: s.OpenedPriceVal(), index: i * 4}
		indexedStocks[i*4+2] = indexedStock{price: s.LowVal(), index: i*4 + 2}
		indexedStocks[i*4+1] = indexedStock{price: s.HighVal(), index: i*4 + 1}
		indexedStocks[i*4+3] = indexedStock{price: s.PriceVal(), index: i*4 + 3}
	}
	sort.Slice(indexedStocks, func(i, j int) bool {
		return indexedStocks[i].price > indexedStocks[j].price
	})

	type indexedRank struct {
		rank  *int
		index int
	}
	indexedRanks := make([]indexedRank, len(s.PricePattern)*4)
	for i, pattern := range s.PricePattern {
		indexedRanks[i*4] = indexedRank{rank: pattern.OpenedPriceRank, index: i * 4}
		indexedRanks[i*4+1] = indexedRank{rank: pattern.HighRank, index: i*4 + 1}
		indexedRanks[i*4+2] = indexedRank{rank: pattern.LowRank, index: i*4 + 2}
		indexedRanks[i*4+3] = indexedRank{rank: pattern.PriceRank, index: i*4 + 3}
	}
	sort.Slice(indexedRanks, func(i, j int) bool {
		if indexedRanks[i].rank == nil {
			return false
		}
		if indexedRanks[j].rank == nil {
			return true
		}
		return *indexedRanks[i].rank < *indexedRanks[j].rank
	})

	for i, rank := range indexedRanks {
		fmt.Println("rank", rank, "stock", indexedStocks[i])
		if rank.rank == nil {
			continue
		}
		if rank.index != indexedStocks[i].index {
			return false
		}
	}
	return true
}

// IsMatchMaXUpDownPattern
// StockPatternのMaXUpDownPatternに一致するかどうか
func (s *StockPattern) IsMatchMaXUpDownPattern(sc StocksCalc) bool {
	for i, pattern := range s.MaXUpDownPattern {
		verifiedStocks := sc.Stocks[len(sc.Stocks)-len(pattern)-1:]
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
