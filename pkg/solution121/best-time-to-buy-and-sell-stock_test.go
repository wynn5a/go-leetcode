package solution121

import (
	"math"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{
			name:   "Empty prices",
			prices: []int{},
			want:   0,
		},
		{
			name:   "Descending prices",
			prices: []int{7, 6, 5, 4, 3, 2, 1},
			want:   0,
		},
		{
			name:   "Ascending prices",
			prices: []int{1, 2, 3, 4, 5, 6, 7},
			want:   6,
		},
		{
			name:   "Constant prices",
			prices: []int{3, 3, 3, 3, 3},
			want:   0,
		},
		{
			name:   "Example prices",
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   5,
		},
		{
			name:   "Single price difference",
			prices: []int{1, 7},
			want:   6,
		},
		{
			name:   "Large numbers",
			prices: []int{math.MaxInt64 - 1, math.MaxInt64},
			want:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
