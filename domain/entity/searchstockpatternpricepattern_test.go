package entity_test

import (
	"testing"

	"gitlab.com/soy-app/stock-api/domain/entity"
)

func TestSearchStockPattern_IsMatchPricePatterns(t *testing.T) {
	type fields struct {
		PricePatterns entity.PricePatterns
	}
	type args struct {
		sc entity.StocksCalc
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "正常系",
			fields: fields{
				PricePatterns: entity.PricePatterns{
					&entity.PricePattern{
						ArrIndex:    0,
						PricePoint:  func() *float64 { f := 0.2; return &f }(),
						IsOver:      nil,
						IsMatchRank: true,
					},
					&entity.PricePattern{
						ArrIndex:    1,
						PricePoint:  func() *float64 { f := 0.4; return &f }(),
						IsOver:      nil,
						IsMatchRank: true,
					},
					&entity.PricePattern{
						ArrIndex:    2,
						PricePoint:  func() *float64 { f := 0.1; return &f }(),
						IsOver:      nil,
						IsMatchRank: true,
					},
					&entity.PricePattern{
						ArrIndex:    3,
						PricePoint:  func() *float64 { f := 0.3; return &f }(),
						IsOver:      nil,
						IsMatchRank: false,
					},
					&entity.PricePattern{
						ArrIndex:    4,
						PricePoint:  func() *float64 { f := 0.7; return &f }(),
						IsOver:      nil,
						IsMatchRank: false,
					},
					&entity.PricePattern{
						ArrIndex:    5,
						PricePoint:  func() *float64 { f := 0.8; return &f }(),
						IsOver:      nil,
						IsMatchRank: false,
					},
					&entity.PricePattern{
						ArrIndex:    6,
						PricePoint:  func() *float64 { f := 0.5; return &f }(),
						IsOver:      nil,
						IsMatchRank: true,
					},
					&entity.PricePattern{
						ArrIndex:    7,
						PricePoint:  func() *float64 { f := 0.6; return &f }(),
						IsOver:      nil,
						IsMatchRank: false,
					},
				},
			},
			args: args{
				sc: entity.StocksCalc{
					Stocks: []*entity.StockCalc{
						{
							Stock: entity.Stock{
								OpenedPrice: func() *float64 { f := 200.00; return &f }(),
								High:        func() *float64 { f := 400.00; return &f }(),
								Low:         func() *float64 { f := 100.00; return &f }(),
								Price:       func() *float64 { f := 300.00; return &f }(),
							},
							Ma: nil,
						},
						{
							Stock: entity.Stock{
								OpenedPrice: func() *float64 { f := 1000.00; return &f }(),
								High:        func() *float64 { f := 1100.00; return &f }(),
								Low:         func() *float64 { f := 500.00; return &f }(),
								Price:       func() *float64 { f := 700.00; return &f }(),
							},
						},
					},
				},
			},
			want: true,
		},
		{
			name: `
			正常系
			最終日のみ陽線で最高値
			`,
			fields: fields{
				PricePatterns: entity.PricePatterns{
					&entity.PricePattern{
						ArrIndex:    0,
						PricePoint:  func() *float64 { f := 0.0; return &f }(),
						IsOver:      nil,
						IsMatchRank: false,
					},
					&entity.PricePattern{
						ArrIndex:    1,
						PricePoint:  func() *float64 { f := 0.0; return &f }(),
						IsOver:      nil,
						IsMatchRank: false,
					},
					&entity.PricePattern{
						ArrIndex:    2,
						PricePoint:  func() *float64 { f := 0.0; return &f }(),
						IsOver:      nil,
						IsMatchRank: false,
					},
					&entity.PricePattern{
						ArrIndex:    3,
						PricePoint:  func() *float64 { f := 0.0; return &f }(),
						IsOver:      nil,
						IsMatchRank: false,
					},
					&entity.PricePattern{
						ArrIndex:    4,
						PricePoint:  func() *float64 { f := 0.0; return &f }(),
						IsOver:      nil,
						IsMatchRank: false,
					},
					&entity.PricePattern{
						ArrIndex:    5,
						PricePoint:  func() *float64 { f := 0.0; return &f }(),
						IsOver:      nil,
						IsMatchRank: false,
					},
					&entity.PricePattern{
						ArrIndex:    6,
						PricePoint:  func() *float64 { f := 0.0; return &f }(),
						IsOver:      nil,
						IsMatchRank: false,
					},
					&entity.PricePattern{
						ArrIndex:    7,
						PricePoint:  func() *float64 { f := 0.0; return &f }(),
						IsOver:      nil,
						IsMatchRank: false,
					},
					&entity.PricePattern{
						ArrIndex:    8,
						PricePoint:  func() *float64 { f := 1.0; return &f }(),
						IsOver:      nil,
						IsMatchRank: false,
					},
					&entity.PricePattern{
						ArrIndex:    9,
						PricePoint:  func() *float64 { f := 2.0; return &f }(),
						IsOver:      nil,
						IsMatchRank: false,
					},
					&entity.PricePattern{
						ArrIndex:    10,
						PricePoint:  func() *float64 { f := 3.0; return &f }(),
						IsOver:      nil,
						IsMatchRank: false,
					},
					&entity.PricePattern{
						ArrIndex:    11,
						PricePoint:  func() *float64 { f := 4.0; return &f }(),
						IsOver:      nil,
						IsMatchRank: false,
					},
				},
			},
			args: args{
				sc: entity.StocksCalc{
					Stocks: []*entity.StockCalc{
						{
							Stock: entity.Stock{
								OpenedPrice: func() *float64 { f := 100.00; return &f }(),
								High:        func() *float64 { f := 200.00; return &f }(),
								Low:         func() *float64 { f := 350.00; return &f }(),
								Price:       func() *float64 { f := 400.00; return &f }(),
							},
						},
						{
							Stock: entity.Stock{
								OpenedPrice: func() *float64 { f := 1000.00; return &f }(),
								High:        func() *float64 { f := 700.00; return &f }(),
								Low:         func() *float64 { f := 700.00; return &f }(),
								Price:       func() *float64 { f := 500.00; return &f }(),
							},
						},
						{
							Stock: entity.Stock{
								OpenedPrice: func() *float64 { f := 1100.00; return &f }(),
								High:        func() *float64 { f := 1200.00; return &f }(),
								Low:         func() *float64 { f := 1050.00; return &f }(),
								Price:       func() *float64 { f := 1150.00; return &f }(),
							},
						},
					},
				},
			},
			want: true,
		},
		{
			name: `
			異常系
			最終日のみ陽線で最高値かどうか確認し、そうではない
			`,
			fields: fields{
				PricePatterns: entity.PricePatterns{
					&entity.PricePattern{
						ArrIndex:    0,
						PricePoint:  func() *float64 { f := 0.0; return &f }(),
						IsOver:      nil,
						IsMatchRank: false,
					},
					&entity.PricePattern{
						ArrIndex:    1,
						PricePoint:  func() *float64 { f := 0.0; return &f }(),
						IsOver:      nil,
						IsMatchRank: false,
					},
					&entity.PricePattern{
						ArrIndex:    2,
						PricePoint:  func() *float64 { f := 0.0; return &f }(),
						IsOver:      nil,
						IsMatchRank: false,
					},
					&entity.PricePattern{
						ArrIndex:    3,
						PricePoint:  func() *float64 { f := 0.0; return &f }(),
						IsOver:      nil,
						IsMatchRank: false,
					},
					&entity.PricePattern{
						ArrIndex:    4,
						PricePoint:  func() *float64 { f := 0.0; return &f }(),
						IsOver:      nil,
						IsMatchRank: false,
					},
					&entity.PricePattern{
						ArrIndex:    5,
						PricePoint:  func() *float64 { f := 0.0; return &f }(),
						IsOver:      nil,
						IsMatchRank: false,
					},
					&entity.PricePattern{
						ArrIndex:    6,
						PricePoint:  func() *float64 { f := 0.0; return &f }(),
						IsOver:      nil,
						IsMatchRank: false,
					},
					&entity.PricePattern{
						ArrIndex:    7,
						PricePoint:  func() *float64 { f := 0.0; return &f }(),
						IsOver:      nil,
						IsMatchRank: false,
					},
					&entity.PricePattern{
						ArrIndex:    8,
						PricePoint:  func() *float64 { f := 1.0; return &f }(),
						IsOver:      nil,
						IsMatchRank: false,
					},
					&entity.PricePattern{
						ArrIndex:    9,
						PricePoint:  func() *float64 { f := 2.0; return &f }(),
						IsOver:      nil,
						IsMatchRank: false,
					},
					&entity.PricePattern{
						ArrIndex:    10,
						PricePoint:  func() *float64 { f := 3.0; return &f }(),
						IsOver:      nil,
						IsMatchRank: false,
					},
					&entity.PricePattern{
						ArrIndex:    11,
						PricePoint:  func() *float64 { f := 4.0; return &f }(),
						IsOver:      nil,
						IsMatchRank: false,
					},
				},
			},
			args: args{
				sc: entity.StocksCalc{
					Stocks: []*entity.StockCalc{
						{
							Stock: entity.Stock{
								OpenedPrice: func() *float64 { f := 100.00; return &f }(),
								High:        func() *float64 { f := 200.00; return &f }(),
								Low:         func() *float64 { f := 350.00; return &f }(),
								Price:       func() *float64 { f := 400.00; return &f }(),
							},
						},
						{
							Stock: entity.Stock{
								OpenedPrice: func() *float64 { f := 1000.00; return &f }(),
								High:        func() *float64 { f := 700.00; return &f }(),
								Low:         func() *float64 { f := 700.00; return &f }(),
								Price:       func() *float64 { f := 500.00; return &f }(),
							},
						},
						{
							Stock: entity.Stock{
								OpenedPrice: func() *float64 { f := 1100.00; return &f }(),
								High:        func() *float64 { f := 1200.00; return &f }(),
								Low:         func() *float64 { f := 700.00; return &f }(),
								Price:       func() *float64 { f := 500.00; return &f }(),
							},
						},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &entity.SearchStockPattern{
				PricePatterns: tt.fields.PricePatterns,
			}
			if got := s.IsMatchPricePatterns(tt.args.sc); got != tt.want {
				t.Errorf("IsMatchPricePatterns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPricePatterns_RankIndex(t *testing.T) {
	rankIndexTestPricePatterns1 := entity.PricePatterns{
		&entity.PricePattern{
			ArrIndex:    0,
			PricePoint:  func() *float64 { f := 0.2; return &f }(),
			IsOver:      nil,
			IsMatchRank: true,
		},
		&entity.PricePattern{
			ArrIndex:    1,
			PricePoint:  func() *float64 { f := 0.4; return &f }(),
			IsOver:      nil,
			IsMatchRank: true,
		},
		&entity.PricePattern{
			ArrIndex:    2,
			PricePoint:  func() *float64 { f := 0.1; return &f }(),
			IsOver:      nil,
			IsMatchRank: true,
		},
		&entity.PricePattern{
			ArrIndex:    3,
			PricePoint:  func() *float64 { f := 0.3; return &f }(),
			IsOver:      nil,
			IsMatchRank: false,
		},
		&entity.PricePattern{
			ArrIndex:    4,
			PricePoint:  func() *float64 { f := 0.7; return &f }(),
			IsOver:      nil,
			IsMatchRank: false,
		},
		&entity.PricePattern{
			ArrIndex:    5,
			PricePoint:  func() *float64 { f := 0.8; return &f }(),
			IsOver:      nil,
			IsMatchRank: false,
		},
		&entity.PricePattern{
			ArrIndex:    6,
			PricePoint:  func() *float64 { f := 0.5; return &f }(),
			IsOver:      nil,
			IsMatchRank: true,
		},
		&entity.PricePattern{
			ArrIndex:    7,
			PricePoint:  func() *float64 { f := 0.6; return &f }(),
			IsOver:      nil,
			IsMatchRank: false,
		},
	}

	type args struct {
		idx int
	}
	tests := []struct {
		name string
		v    entity.PricePatterns
		args args
		want int
	}{
		{
			name: "pattern 1",
			v: entity.PricePatterns{
				{
					PricePatternID:       "",
					SearchStockPatternID: "",
					ArrIndex:             0,
					PricePoint:           func() *float64 { f := 0.0; return &f }(),
					IsOver:               nil,
					IsMatchRank:          false,
				},
				{
					PricePatternID:       "",
					SearchStockPatternID: "",
					ArrIndex:             1,
					PricePoint:           func() *float64 { f := 1.0; return &f }(),
					IsOver:               nil,
					IsMatchRank:          false,
				},
				{
					PricePatternID:       "",
					SearchStockPatternID: "",
					ArrIndex:             2,
					PricePoint:           func() *float64 { f := 0.5; return &f }(),
					IsOver:               nil,
					IsMatchRank:          false,
				},
				{
					PricePatternID:       "",
					SearchStockPatternID: "",
					ArrIndex:             3,
					PricePoint:           func() *float64 { f := 0.7; return &f }(),
					IsOver:               nil,
					IsMatchRank:          false,
				},
				{
					PricePatternID:       "",
					SearchStockPatternID: "",
					ArrIndex:             4,
					PricePoint:           func() *float64 { f := 0.6; return &f }(),
					IsOver:               nil,
					IsMatchRank:          false,
				},
			},
			args: args{
				idx: 1,
			},
			want: 0,
		},
		{
			name: "pattern1-0",
			v:    rankIndexTestPricePatterns1,
			args: args{
				idx: 0,
			},
			want: 6,
		},
		{
			name: "pattern1-1",
			v:    rankIndexTestPricePatterns1,
			args: args{
				idx: 1,
			},
			want: 4,
		},
		{
			name: "pattern1-2",
			v:    rankIndexTestPricePatterns1,
			args: args{
				idx: 2,
			},
			want: 7,
		},
		{
			name: "pattern1-3",
			v:    rankIndexTestPricePatterns1,
			args: args{
				idx: 3,
			},
			want: 5,
		},
		{
			name: "pattern1-4",
			v:    rankIndexTestPricePatterns1,
			args: args{
				idx: 4,
			},
			want: 1,
		},
		{
			name: "pattern1-5",
			v:    rankIndexTestPricePatterns1,
			args: args{
				idx: 5,
			},
			want: 0,
		},
		{
			name: "pattern1-6",
			v:    rankIndexTestPricePatterns1,
			args: args{
				idx: 6,
			},
			want: 3,
		},
		{
			name: "pattern1-6",
			v:    rankIndexTestPricePatterns1,
			args: args{
				idx: 7,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.RankIndex(tt.args.idx); got != tt.want {
				t.Errorf("RankIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
