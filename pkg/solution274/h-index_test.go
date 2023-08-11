package solution274

import "testing"

func TestHIndex(t *testing.T) {
	tests := []struct {
		name      string
		citations []int
		want      int
	}{
		{
			name:      "Test 1",
			citations: []int{3, 0, 6, 1, 5},
			want:      3,
		},
		{
			name:      "Test 2",
			citations: []int{1, 3, 5, 7, 9},
			want:      3,
		},
		{
			name:      "Test 3",
			citations: []int{5, 5, 5, 5},
			want:      4,
		},
		{
			name:      "Test 4",
			citations: []int{1, 1, 1, 1},
			want:      1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hIndex(tt.citations); got != tt.want {
				t.Errorf("hIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
