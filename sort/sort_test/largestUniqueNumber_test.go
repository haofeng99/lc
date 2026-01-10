package base_sort_test

import (
	"testing"

	"lc/sort"
)

func TestLargestUniqueNumber(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{"unique max", []int{1, 2, 3, 2, 1, 4}, 4},
		{"no unique", []int{1, 1, 2, 2}, -1},
		{"single", []int{10}, 10},
		{"zeroes", []int{0, 0, 0, 1}, 1},
		{"max duplicate but smaller unique", []int{5, 5, 3, 4, 4}, 3},
		{"large range", []int{1000, 999, 1000}, 999},
		{"empty", []int{}, -1},
	}

	for _, tc := range tests {
		got := sort.TestlargestUniqueNumber(tc.arr)
		if got != tc.want {
			t.Errorf("%s: got %d, want %d (arr=%v)", tc.name, got, tc.want, tc.arr)
		}
	}
}
