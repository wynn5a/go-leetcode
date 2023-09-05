package solution2605

import (
	"testing"
)

func TestCase(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  int
	}{{
		name:  "test case 1",
		nums1: []int{4, 1, 3},
		nums2: []int{5, 7},
		want:  15,
	},
		{
			name:  "test case 2",
			nums1: []int{3, 5, 2, 6},
			nums2: []int{3, 1, 7},
			want:  3,
		}, {
			name:  "test case 3",
			nums1: []int{3, 8, 4, 2, 6, 1},
			nums2: []int{4, 7, 8, 5, 2, 3, 6},
			want:  2,
		},
	}

	for _, v := range tests {
		actual := minNumber(v.nums1, v.nums2)
		if v.want != actual {
			t.Errorf("%s failed, expected: %d, actual: %d", v.name, v.want, actual)
		}
	}
}
