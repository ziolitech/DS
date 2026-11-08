package pairs

import "testing"

func Test_countPairsDivisibleByK(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		arr  []int
		k    int
		want int
	}{
		{
			"Check if the pair exists with difference K",
			[]int{1, 3, 4, 5, 6},
			3,
			2,
		},
		{
			// This test case fails if we don't consider frequency.
			"Check if the pair exists with difference K",
			[]int{3, 6, 9}, // (3,6), (3,9), (6,9)
			3,
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countPairsDivisibleByK(tt.arr, tt.k)

			if got != tt.want {
				t.Errorf("countPairsDivisibleByK() = %v, want %v", got, tt.want)
			}
		})
	}
}
