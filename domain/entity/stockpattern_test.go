package entity

import "testing"

func TestStockPattern_IsMatchMaXUpDownPattern(t *testing.T) {
	type fields struct {
		MaXUpDownPattern map[int][]bool
	}
	type args struct {
		stocks []*StockCalc
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
				stocks: []*StockCalc{
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
