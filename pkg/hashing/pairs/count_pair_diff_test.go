package pairs

import "testing"

func Test_checkPairExistsWithDiff(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		arr  []int // arg param 1
		diff int   // arg param 2
		want int   // return type
	}{
		{
			"Check if the pair exists with difference K",
			[]int{1, 3, 4, 5, 6},
			1,
			3,
		},
		{
			"Check if the pair exists with difference K",
			[]int{1, 5, 7, 3, 2},
			7,
			0,
		},
		{
			"Check if the pair exists with difference K",
			[]int{1, 5, 7, 3, 2},
			2,
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPairExistsWithDiff(tt.arr, tt.diff); got != tt.want {
				t.Errorf("checkPairExistsWithDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
