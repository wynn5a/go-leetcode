package solution1706

import (
	"reflect"
	"testing"
)

func TestFindBall(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want []int
	}{
		{
			name: "Test Case 1",
			grid: [][]int{
				{1, 1, 1, 1, 1, 1},
				{-1, -1, -1, -1, -1, -1},
				{1, 1, 1, 1, 1, 1},
				{-1, -1, -1, -1, -1, -1},
			},
			want: []int{0, 1, 2, 3, 4, -1},
		},
		{
			name: "Test Case 2",
			grid: [][]int{
				{-1},
			},
			want: []int{-1},
		},
		{
			name: "Test Case 3",
			grid: [][]int{
				{1, 1, 1, -1, -1},
				{1, 1, 1, -1, -1},
				{-1, -1, -1, 1, 1},
				{1, 1, 1, 1, -1},
				{-1, -1, -1, -1, -1},
			},
			want: []int{1, -1, -1, -1, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findBall(tt.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findBall() = %v, want %v", got, tt.want)
			}
		})
	}
}
