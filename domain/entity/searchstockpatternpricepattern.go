package entity

import (
	"sort"
)

func (v *PricePatterns) RankIndex(idx int) int {
	pricePatterns := *v
	sort.Slice(pricePatterns, func(i, j int) bool {
		return *pricePatterns[i].PricePoint < *pricePatterns[j].PricePoint
	})

	for i := range pricePatterns {
		if pricePatterns[i].ArrIndex == idx {
			sameCount := 0
			for j := i - 1; j > 0; j-- {
				if *pricePatterns[i].PricePoint == *pricePatterns[j].PricePoint {
					sameCount++
					continue
				}
				break
			}
			return len(pricePatterns) - 1 - i + sameCount
		}
	}
	return -1
}

func (s *SearchStockPattern) IsMatchPricePatterns(sc StocksCalc) bool {
	// 実際の価格のrank算出
	type IndexedPriceRank struct {
		price      float64
		pricePoint float64
		rankIndex  int
		index      int
	}
	indexedPriceRanks := make([]IndexedPriceRank, len(sc.Stocks)*4)
	for i, s := range sc.getByIdxRange(
		len(sc.Stocks)-len(s.PricePatterns)/4,
		len(sc.Stocks)).Stocks {
		indexedPriceRanks[i*4] = IndexedPriceRank{price: s.OpenedPriceVal(), index: i * 4}
		indexedPriceRanks[i*4+1] = IndexedPriceRank{price: s.HighVal(), index: i*4 + 1}
		indexedPriceRanks[i*4+2] = IndexedPriceRank{price: s.LowVal(), index: i*4 + 2}
		indexedPriceRanks[i*4+3] = IndexedPriceRank{price: s.PriceVal(), index: i*4 + 3}
	}
	sort.Slice(indexedPriceRanks, func(i, j int) bool {
		return indexedPriceRanks[i].price > indexedPriceRanks[j].price
	})
	var maxPrice float64
	for i := range indexedPriceRanks {
		sameCount := 0
		for j := i - 1; j > 0; j-- {
			if indexedPriceRanks[i].price == indexedPriceRanks[j].price {
				sameCount++
				continue
			}
			break
		}
		indexedPriceRanks[i].rankIndex = i - sameCount

		if indexedPriceRanks[i].price > maxPrice {
			maxPrice = indexedPriceRanks[i].price
		}
	}
	for i := range indexedPriceRanks {
		indexedPriceRanks[i].pricePoint = indexedPriceRanks[i].price / maxPrice
	}
	sort.Slice(indexedPriceRanks, func(i, j int) bool {
		return indexedPriceRanks[i].index < indexedPriceRanks[j].index
	})

	// 価格パターンのrank算出
	passCount := 0
	for i, pattern := range s.PricePatterns {
		if !pattern.IsMatchRank {
			passCount++
			continue
		}
		if indexedPriceRanks[i].rankIndex != s.PricePatterns.RankIndex(i) {
			return false
		}
		if pattern.IsOver == nil {
			continue
		}
		if (*pattern.IsOver && indexedPriceRanks[i].pricePoint < *pattern.PricePoint) ||
			(!*pattern.IsOver && indexedPriceRanks[i].pricePoint > *pattern.PricePoint) {
			return false
		}
	}
	return true
}
