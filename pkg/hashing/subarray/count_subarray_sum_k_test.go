package subarray

import (
	"testing"
)

func TestCountSubarraysWith_K_Sum(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		arr  []int
		k    int
		want int
	}{
		{
			"test 1",
			[]int{10, 2, -2, -20, 10},
			-10,
			3,
		},
		{
			"test 2",
			[]int{9, 4, 20, 3, 10, 5},
			33,
			2,
		},
		{
			"test 3",
			[]int{1, 3, 5},
			2,
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountSubarraysWith_K_Sum(tt.arr, tt.k)
			if got != tt.want {
				t.Errorf("CountSubarraysWith_K_Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
