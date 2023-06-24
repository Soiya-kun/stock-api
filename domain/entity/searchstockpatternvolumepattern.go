package entity

import (
	"sort"
)

func (v *VolumePatterns) RankIndex(idx int) int {
	volumePatterns := *v
	sort.Slice(volumePatterns, func(i, j int) bool {
		return *volumePatterns[i].VolumePoint > *volumePatterns[j].VolumePoint
	})

	for i := range *v {
		if volumePatterns[i].ArrIndex == idx {
			sameCount := 0
			for j := i - 1; j > 0; j-- {
				if *volumePatterns[i].VolumePoint == *volumePatterns[j].VolumePoint {
					sameCount++
					continue
				}
				break
			}
			return i - sameCount
		}
	}
	return -1
}

func (s *SearchStockPattern) IsMatchVolumePatterns(sc StocksCalc) bool {
	stocksStruct := sc.getByIdxRange(
		len(sc.Stocks)-len(s.VolumePatterns),
		len(sc.Stocks),
	)

	for _, stock := range stocksStruct.Stocks {
		*stock.Volume = *stock.Volume / stocksStruct.MaxVolume()
	}

	type IndexedVolumeRank struct {
		volume    float64
		rankIndex int
		index     int
	}
	indexedVolumeRanks := make([]IndexedVolumeRank, len(stocksStruct.Stocks))
	for i, stock := range stocksStruct.Stocks {
		indexedVolumeRanks[i] = IndexedVolumeRank{volume: *stock.Volume, index: i}
	}
	sort.Slice(indexedVolumeRanks, func(i, j int) bool {
		return indexedVolumeRanks[i].volume > indexedVolumeRanks[j].volume
	})

	for i, v := range indexedVolumeRanks {
		sameCount := 0
		for j := i - 1; j > 0; j-- {
			if v.volume == indexedVolumeRanks[j].volume {
				sameCount++
				continue
			}
			break
		}
		indexedVolumeRanks[i].rankIndex = i - sameCount
	}
	sort.Slice(indexedVolumeRanks, func(i, j int) bool {
		return indexedVolumeRanks[i].index < indexedVolumeRanks[j].index
	})

	passCount := 0
	for i, pattern := range s.VolumePatterns {
		if pattern.IsMatchRank == false {
			passCount++
			continue
		}
		if indexedVolumeRanks[i].rankIndex != s.VolumePatterns.RankIndex(i) {
			return false
		}
		if pattern.IsOver == nil {
			continue
		}
		if (*pattern.IsOver && indexedVolumeRanks[i].rankIndex < s.VolumePatterns.RankIndex(i)) ||
			(!*pattern.IsOver && indexedVolumeRanks[i].rankIndex > s.VolumePatterns.RankIndex(i)) {
			return false
		}
	}
	return true
}
