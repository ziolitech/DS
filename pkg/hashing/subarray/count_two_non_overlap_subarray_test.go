package subarray

import "testing"

func Test_countTwoNonOverlapSubarray(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		input []int
		want  int
	}{
		{
			"basic test - non overlapping array",
			[]int{1, 2, 3, -7, 8, 9, -15, 1},
			23,
		},
		{
			"basic test - 2 ",
			[]int{1, 2, 3, 4, 5, 6},
			21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countTwoNonOverlapSubarray(tt.input)
			if got != tt.want {
				t.Errorf("countTwoNonOverlapSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
