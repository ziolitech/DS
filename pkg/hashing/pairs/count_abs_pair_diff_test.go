package pairs

import "testing"

func Test_checkPairExistsWithAbsDiff(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		arr        []int
		targetDiff int
		want       int
	}{
		{
			"Check if the pair exists with difference K",
			[]int{1, 3, 4, 5, 6},
			3,
			2,
		},
		{
			"Check if the pair exists with difference K",
			[]int{1, 3, 4, 5, 6},
			7,
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countPairExistsWithAbsDiff(tt.arr, tt.targetDiff)
			if got != tt.want {
				t.Errorf("checkPairExistsWithAbsDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
