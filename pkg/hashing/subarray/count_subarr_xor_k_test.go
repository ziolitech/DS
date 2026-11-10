package subarray

import (
	"testing"
)

func Test_countSubArrWithXOR_K(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		input []int
		k     int
		want  int
	}{
		{
			"simple test - count subarr with XOR",
			[]int{4, 2, 2, 6, 4},
			6,
			4,
		},
		{
			"Edge test - count subarr with XOR",
			[]int{4},
			4,
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countSubArrWithXOR_K(tt.input, tt.k)
			if got != tt.want {
				t.Errorf("countSubArrWithXOR_K() = %v, want %v", got, tt.want)
			}
		})
	}
}
