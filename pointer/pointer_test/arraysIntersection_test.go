package pointer_test

import (
	pointer "lc/pointer"
	"reflect"
	"testing"
)

func TestArraysIntersection(t *testing.T) {
	tests := []struct {
		name    string
		a, b, c []int
		want    []int
	}{
		{"common 1 and 5", []int{1, 2, 3, 4, 5}, []int{1, 2, 5, 7, 9}, []int{1, 3, 4, 5, 8}, []int{1, 5}},
		{"no common", []int{1, 2}, []int{3, 4}, []int{5, 6}, []int{}},
		{"all equal", []int{1, 2}, []int{1, 2}, []int{1, 2}, []int{1, 2}},
		{"single element", []int{1}, []int{1}, []int{1}, []int{1}},
		{"one empty", []int{}, []int{1, 2, 3}, []int{1, 2, 3}, []int{}},
		{"different lengths", []int{1, 2, 3, 6, 7}, []int{2, 3, 4}, []int{3, 6, 7}, []int{3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pointer.TestArraysIntersection(tt.a, tt.b, tt.c)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("arraysIntersection(%v, %v, %v) = %v, want %v", tt.a, tt.b, tt.c, got, tt.want)
			}
		})
	}
}
