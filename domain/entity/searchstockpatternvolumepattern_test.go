package entity_test

import (
	"testing"

	"gitlab.com/soy-app/stock-api/domain/entity"
)

func floatPointer(f float64) *float64 {
	return &f
}

func TestVolumePatterns_RankIndex(t *testing.T) {
	type args struct {
		idx int
	}
	tests := []struct {
		name string
		v    entity.VolumePatterns
		args args
		want int
	}{
		{
			name: "RankIndex(0) should return 1",
			v: entity.VolumePatterns{
				{ArrIndex: 0, VolumePoint: floatPointer(0.4)},
				{ArrIndex: 1, VolumePoint: floatPointer(0.3)},
				{ArrIndex: 2, VolumePoint: floatPointer(0.3)},
				{ArrIndex: 3, VolumePoint: floatPointer(0.5)},
				{ArrIndex: 4, VolumePoint: floatPointer(0.1)},
			},
			args: args{
				idx: 0,
			},
			want: 1,
		},
		{
			name: "RankIndex(1) should return 2",
			v: entity.VolumePatterns{
				{ArrIndex: 0, VolumePoint: floatPointer(0.4)},
				{ArrIndex: 1, VolumePoint: floatPointer(0.3)},
				{ArrIndex: 2, VolumePoint: floatPointer(0.3)},
				{ArrIndex: 3, VolumePoint: floatPointer(0.5)},
				{ArrIndex: 4, VolumePoint: floatPointer(0.1)},
			},
			args: args{
				idx: 1,
			},
			want: 2,
		},
		{
			name: "RankIndex(2) should return 2",
			v: entity.VolumePatterns{
				{ArrIndex: 0, VolumePoint: floatPointer(0.4)},
				{ArrIndex: 1, VolumePoint: floatPointer(0.3)},
				{ArrIndex: 2, VolumePoint: floatPointer(0.3)},
				{ArrIndex: 3, VolumePoint: floatPointer(0.5)},
				{ArrIndex: 4, VolumePoint: floatPointer(0.1)},
			},
			args: args{
				idx: 2,
			},
			want: 2,
		},
		{
			name: "RankIndex(3) should return 0",
			v: entity.VolumePatterns{
				{ArrIndex: 0, VolumePoint: floatPointer(0.4)},
				{ArrIndex: 1, VolumePoint: floatPointer(0.3)},
				{ArrIndex: 2, VolumePoint: floatPointer(0.3)},
				{ArrIndex: 3, VolumePoint: floatPointer(0.5)},
				{ArrIndex: 4, VolumePoint: floatPointer(0.1)},
			},
			args: args{
				idx: 3,
			},
			want: 0,
		},
		{
			name: "RankIndex(4) should return 4",
			v: entity.VolumePatterns{
				{ArrIndex: 0, VolumePoint: floatPointer(0.4)},
				{ArrIndex: 1, VolumePoint: floatPointer(0.3)},
				{ArrIndex: 2, VolumePoint: floatPointer(0.3)},
				{ArrIndex: 3, VolumePoint: floatPointer(0.5)},
				{ArrIndex: 4, VolumePoint: floatPointer(0.1)},
			},
			args: args{
				idx: 4,
			},
			want: 4,
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

func TestSearchStockPattern_IsMatchVolumePatterns(t *testing.T) {
	type fields struct {
		SearchStockPatternID string
		UserID               string
		User                 entity.User
		VolumePatterns       entity.VolumePatterns
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
			name: "IsMatchVolumePatterns should return true",
			fields: fields{
				VolumePatterns: []*entity.VolumePattern{
					{ArrIndex: 0, VolumePoint: floatPointer(0.4), IsMatchRank: true},
					{ArrIndex: 1, VolumePoint: floatPointer(0.3), IsMatchRank: true},
					{ArrIndex: 2, VolumePoint: floatPointer(0.3), IsMatchRank: true},
					{ArrIndex: 3, VolumePoint: floatPointer(1.0), IsMatchRank: true},
					{ArrIndex: 4, VolumePoint: floatPointer(0.0), IsMatchRank: true},
				},
			},
			args: args{
				sc: entity.StocksCalc{
					Stocks: []*entity.StockCalc{
						{Stock: entity.Stock{Volume: floatPointer(400)}},
						{Stock: entity.Stock{Volume: floatPointer(300)}},
						{Stock: entity.Stock{Volume: floatPointer(300)}},
						{Stock: entity.Stock{Volume: floatPointer(1000)}},
						{Stock: entity.Stock{Volume: floatPointer(0)}},
					},
				},
			},
			want: true,
		},
		{
			name: "IsMatchVolumePatterns should return false",
			fields: fields{
				VolumePatterns: []*entity.VolumePattern{
					{ArrIndex: 0, VolumePoint: floatPointer(0.4), IsMatchRank: true},
					{ArrIndex: 1, VolumePoint: floatPointer(0.5), IsMatchRank: true},
					{ArrIndex: 2, VolumePoint: floatPointer(0.3), IsMatchRank: true},
					{ArrIndex: 3, VolumePoint: floatPointer(1.0), IsMatchRank: true},
					{ArrIndex: 4, VolumePoint: floatPointer(0.0), IsMatchRank: true},
				},
			},
			args: args{
				sc: entity.StocksCalc{
					Stocks: []*entity.StockCalc{
						{Stock: entity.Stock{Volume: floatPointer(400)}},
						{Stock: entity.Stock{Volume: floatPointer(300)}},
						{Stock: entity.Stock{Volume: floatPointer(300)}},
						{Stock: entity.Stock{Volume: floatPointer(1000)}},
						{Stock: entity.Stock{Volume: floatPointer(0)}},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &entity.SearchStockPattern{
				VolumePatterns: tt.fields.VolumePatterns,
			}
			if got := s.IsMatchVolumePatterns(tt.args.sc); got != tt.want {
				t.Errorf("IsMatchVolumePatterns() = %v, want %v", got, tt.want)
			}
		})
	}
}
