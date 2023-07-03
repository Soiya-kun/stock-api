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
						PricePoint:  func() *float64 { f := 0.1; return &f }(),
						IsOver:      nil,
						IsMatchRank: true,
					},
					&entity.PricePattern{
						ArrIndex:    1,
						PricePoint:  func() *float64 { f := 0.2; return &f }(),
						IsOver:      nil,
						IsMatchRank: true,
					},
					&entity.PricePattern{
						ArrIndex:    2,
						PricePoint:  func() *float64 { f := 0.3; return &f }(),
						IsOver:      nil,
						IsMatchRank: true,
					},
					&entity.PricePattern{
						ArrIndex:    3,
						PricePoint:  func() *float64 { f := 0.4; return &f }(),
						IsOver:      nil,
						IsMatchRank: false,
					},
					&entity.PricePattern{
						ArrIndex:    4,
						PricePoint:  func() *float64 { f := 1.0; return &f }(),
						IsOver:      nil,
						IsMatchRank: false,
					},
					&entity.PricePattern{
						ArrIndex:    5,
						PricePoint:  func() *float64 { f := 0.7; return &f }(),
						IsOver:      nil,
						IsMatchRank: false,
					},
					&entity.PricePattern{
						ArrIndex:    6,
						PricePoint:  func() *float64 { f := 0.6; return &f }(),
						IsOver:      nil,
						IsMatchRank: true,
					},
					&entity.PricePattern{
						ArrIndex:    7,
						PricePoint:  func() *float64 { f := 0.5; return &f }(),
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
							Ma: nil,
						},
						{
							Stock: entity.Stock{
								OpenedPrice: func() *float64 { f := 1000.00; return &f }(),
								High:        func() *float64 { f := 700.00; return &f }(),
								Low:         func() *float64 { f := 700.00; return &f }(),
								Price:       func() *float64 { f := 500.00; return &f }(),
							},
						},
					},
				},
			},
			want: true,
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
