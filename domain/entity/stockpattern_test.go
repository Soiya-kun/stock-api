package entity

import "testing"

func TestStockPattern_IsMatchMaXUpDownPattern(t *testing.T) {
	type fields struct {
		MaXUpDownPattern map[int][]bool
	}
	type args struct {
		stocks StocksCalc
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
				MaXUpDownPattern: map[int][]bool{
					5:  {true, false, true, true, true},
					20: {true, true, true, true, true},
				},
			},
			args: args{
				stocks: StocksCalc{
					Stocks: []*StockCalc{
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
			s := &StockPattern{
				MaXUpDownPattern: tt.fields.MaXUpDownPattern,
			}
			if got := s.IsMatchMaXUpDownPattern(tt.args.stocks); got != tt.want {
				t.Errorf("IsMatchMaXUpDownPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStockPattern_IsMaxVolumeInDaysOverAverage(t *testing.T) {
	type fields struct {
		MaxVolumeInDaysIsOverAverage struct {
			Day         int     // N日間
			OverAverage float64 // 平均出来高の何倍か
		}
		MaXUpDownPattern map[int][]bool
	}
	type args struct {
		sc StocksCalc
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
				MaxVolumeInDaysIsOverAverage: struct {
					Day         int
					OverAverage float64
				}{Day: 5, OverAverage: 1.5},
			},
			args: args{
				sc: StocksCalc{
					Stocks: []*StockCalc{
						{Stock: Stock{Volume: func() *float64 { var v float64 = 100; return &v }()}},
						{Stock: Stock{Volume: func() *float64 { var v float64 = 100; return &v }()}},
						{Stock: Stock{Volume: func() *float64 { var v float64 = 100; return &v }()}},
						{Stock: Stock{Volume: func() *float64 { var v float64 = 100; return &v }()}},
						{Stock: Stock{Volume: func() *float64 { var v float64 = 100; return &v }()}},
						{Stock: Stock{Volume: func() *float64 { var v float64 = 100; return &v }()}},
						{Stock: Stock{Volume: func() *float64 { var v float64 = 160; return &v }()}},
						{Stock: Stock{Volume: func() *float64 { var v float64 = 100; return &v }()}},
					},
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StockPattern{
				MaxVolumeInDaysIsOverAverage: tt.fields.MaxVolumeInDaysIsOverAverage,
				MaXUpDownPattern:             tt.fields.MaXUpDownPattern,
			}
			if got := s.IsMaxVolumeInDaysOverAverage(tt.args.sc); got != tt.want {
				t.Errorf("IsMaxVolumeInDaysOverAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStockPattern_IsMatchPricePattern(t *testing.T) {
	type fields struct {
		MaxVolumeInDaysIsOverAverage struct {
			Day         int     // N日間
			OverAverage float64 // 平均出来高の何倍か
		}
		PricePattern []struct {
			PriceRank       *int // "終値"の順位
			OpenedPriceRank *int // "始値"の順位
			HighRank        *int // "高値"の順位
			LowRank         *int // "安値"の順位
		}
		MaXUpDownPattern map[int][]bool
	}
	type args struct {
		sc StocksCalc
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
				PricePattern: []struct {
					PriceRank       *int
					OpenedPriceRank *int
					HighRank        *int
					LowRank         *int
				}{
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
				sc: StocksCalc{
					Stocks: []*StockCalc{
						{
							Stock: Stock{
								Price:       func() *float64 { v := 20.0; return &v }(),
								OpenedPrice: func() *float64 { v := 0.0; return &v }(),
								High:        func() *float64 { v := -10.0; return &v }(),
								Low:         func() *float64 { v := 5.0; return &v }(),
							},
						},
						{
							Stock: Stock{
								OpenedPrice: func() *float64 { v := 8.0; return &v }(),
								High:        func() *float64 { v := 7.0; return &v }(),
								Low:         func() *float64 { v := 6.0; return &v }(),
								Price:       func() *float64 { v := 5.0; return &v }(),
							},
						},
						{
							Stock: Stock{
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
				PricePattern: []struct {
					PriceRank       *int
					OpenedPriceRank *int
					HighRank        *int
					LowRank         *int
				}{
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
				sc: StocksCalc{
					Stocks: []*StockCalc{
						{
							Stock: Stock{
								Price:       func() *float64 { v := 20.0; return &v }(),
								OpenedPrice: func() *float64 { v := 0.0; return &v }(),
								High:        func() *float64 { v := -10.0; return &v }(),
								Low:         func() *float64 { v := 5.0; return &v }(),
							},
						},
						{
							Stock: Stock{
								OpenedPrice: func() *float64 { v := 120.0; return &v }(),
								High:        func() *float64 { v := 150.0; return &v }(),
								Low:         func() *float64 { v := 110.0; return &v }(),
								Price:       func() *float64 { v := 130.0; return &v }(),
							},
						},
						{
							Stock: Stock{
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
			s := &StockPattern{
				MaxVolumeInDaysIsOverAverage: tt.fields.MaxVolumeInDaysIsOverAverage,
				PricePattern:                 tt.fields.PricePattern,
				MaXUpDownPattern:             tt.fields.MaXUpDownPattern,
			}
			if got := s.IsMatchPricePattern(tt.args.sc); got != tt.want {
				t.Errorf("IsMatchPricePattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
