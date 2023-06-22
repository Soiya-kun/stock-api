package entity_test

import (
	"testing"

	"gitlab.com/soy-app/stock-api/domain/entity"
)

func TestStockPattern_IsMatchMaXUpDownPattern(t *testing.T) { //nolint:paralleltest
	type fields struct {
		MaXUpDownPattern []*entity.MaXUpDownPattern
	}
	type args struct {
		stocks entity.StocksCalc
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "success",
			fields: fields{
				MaXUpDownPattern: []*entity.MaXUpDownPattern{
					{
						MaX:     5,
						Pattern: "10111",
					},
					{
						MaX:     20,
						Pattern: "11111",
					},
				},
			},
			args: args{
				stocks: entity.StocksCalc{
					Stocks: []*entity.StockCalc{
						{
							Ma: map[int]float64{
								5:  100,
								20: 100,
							},
						},
						{
							Ma: map[int]float64{
								5:  90,
								20: 110,
							},
						},
						{
							Ma: map[int]float64{
								5:  120,
								20: 120,
							},
						},
						{
							Ma: map[int]float64{
								5:  130,
								20: 130,
							},
						},
						{
							Ma: map[int]float64{
								5:  140,
								20: 140,
							},
						},
						{
							Ma: map[int]float64{
								5:  150,
								20: 150,
							},
						},
						{
							Ma: map[int]float64{
								5:  160,
								20: 160,
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
				MaXUpDownPatterns: tt.fields.MaXUpDownPattern,
			}
			if got := s.IsMatchMaXUpDownPattern(tt.args.stocks); got != tt.want {
				t.Errorf("IsMatchMaXUpDownPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStockPattern_IsMaxVolumeInDaysOverAverage(t *testing.T) {
	type fields struct {
		MaxVolumeInDaysIsOverAverage *entity.MaxVolumeInDaysIsOverAverage
	}
	type args struct {
		sc entity.StocksCalc
	}
	var tests = []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "success",
			fields: fields{
				MaxVolumeInDaysIsOverAverage: &entity.MaxVolumeInDaysIsOverAverage{Day: 5, OverAverage: 1.5},
			},
			args: args{
				sc: entity.StocksCalc{
					Stocks: []*entity.StockCalc{
						{Stock: entity.Stock{Volume: func() *float64 { var v float64 = 100; return &v }()}},
						{Stock: entity.Stock{Volume: func() *float64 { var v float64 = 100; return &v }()}},
						{Stock: entity.Stock{Volume: func() *float64 { var v float64 = 100; return &v }()}},
						{Stock: entity.Stock{Volume: func() *float64 { var v float64 = 100; return &v }()}},
						{Stock: entity.Stock{Volume: func() *float64 { var v float64 = 100; return &v }()}},
						{Stock: entity.Stock{Volume: func() *float64 { var v float64 = 100; return &v }()}},
						{Stock: entity.Stock{Volume: func() *float64 { var v float64 = 160; return &v }()}},
						{Stock: entity.Stock{Volume: func() *float64 { var v float64 = 100; return &v }()}},
					},
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &entity.SearchStockPattern{
				MaxVolumeInDaysIsOverAverage: tt.fields.MaxVolumeInDaysIsOverAverage,
			}
			if got := s.IsMaxVolumeInDaysOverAverage(tt.args.sc); got != tt.want {
				t.Errorf("IsMaxVolumeInDaysOverAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStockPattern_IsMatchPricePattern(t *testing.T) {
	type fields struct {
		MaxVolumeInDaysIsOverAverage *entity.MaxVolumeInDaysIsOverAverage
		PricePattern                 []*entity.PricePattern
		MaXUpDownPatterns            []*entity.MaXUpDownPattern
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
			name: "success",
			fields: fields{
				PricePattern: []*entity.PricePattern{
					{
						OpenedPriceRank: func() *int { v := 1; return &v }(),
						HighRank:        func() *int { v := 2; return &v }(),
						LowRank:         func() *int { v := 3; return &v }(),
						PriceRank:       func() *int { v := 4; return &v }(),
					},
					{
						OpenedPriceRank: func() *int { v := 5; return &v }(),
						HighRank:        func() *int { v := 6; return &v }(),
						LowRank:         func() *int { v := 7; return &v }(),
						PriceRank:       func() *int { v := 8; return &v }(),
					},
				},
			},
			args: args{
				sc: entity.StocksCalc{
					Stocks: []*entity.StockCalc{
						{
							Stock: entity.Stock{
								Price:       func() *float64 { v := 20.0; return &v }(),
								OpenedPrice: func() *float64 { v := 0.0; return &v }(),
								High:        func() *float64 { v := -10.0; return &v }(),
								Low:         func() *float64 { v := 5.0; return &v }(),
							},
						},
						{
							Stock: entity.Stock{
								OpenedPrice: func() *float64 { v := 8.0; return &v }(),
								High:        func() *float64 { v := 7.0; return &v }(),
								Low:         func() *float64 { v := 6.0; return &v }(),
								Price:       func() *float64 { v := 5.0; return &v }(),
							},
						},
						{
							Stock: entity.Stock{
								OpenedPrice: func() *float64 { v := 4.0; return &v }(),
								High:        func() *float64 { v := 3.0; return &v }(),
								Low:         func() *float64 { v := 2.0; return &v }(),
								Price:       func() *float64 { v := 1.0; return &v }(),
							},
						},
					},
				},
			},
			want: true,
		},
		{
			name: "success",
			fields: fields{
				PricePattern: []*entity.PricePattern{
					{
						OpenedPriceRank: nil,
						HighRank:        func() *int { v := 1; return &v }(),
						LowRank:         nil,
						PriceRank:       nil,
					},
					{
						OpenedPriceRank: nil,
						HighRank:        func() *int { v := 2; return &v }(),
						LowRank:         nil,
						PriceRank:       nil,
					},
				},
			},
			args: args{
				sc: entity.StocksCalc{
					Stocks: []*entity.StockCalc{
						{
							Stock: entity.Stock{
								Price:       func() *float64 { v := 20.0; return &v }(),
								OpenedPrice: func() *float64 { v := 0.0; return &v }(),
								High:        func() *float64 { v := -10.0; return &v }(),
								Low:         func() *float64 { v := 5.0; return &v }(),
							},
						},
						{
							Stock: entity.Stock{
								OpenedPrice: func() *float64 { v := 120.0; return &v }(),
								High:        func() *float64 { v := 150.0; return &v }(),
								Low:         func() *float64 { v := 110.0; return &v }(),
								Price:       func() *float64 { v := 130.0; return &v }(),
							},
						},
						{
							Stock: entity.Stock{
								OpenedPrice: func() *float64 { v := 135.0; return &v }(),
								High:        func() *float64 { v := 140.0; return &v }(),
								Low:         func() *float64 { v := 100.0; return &v }(),
								Price:       func() *float64 { v := 105.0; return &v }(),
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
				MaxVolumeInDaysIsOverAverage: tt.fields.MaxVolumeInDaysIsOverAverage,
				PricePattern:                 tt.fields.PricePattern,
				MaXUpDownPatterns:            tt.fields.MaXUpDownPatterns,
			}
			if got := s.IsMatchPricePattern(tt.args.sc); got != tt.want {
				t.Errorf("IsMatchPricePattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
