package misc

import "testing"

func Test_findLongestCommonSequence(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want int
	}{
		{
			"simple test",
			[]int{1, 2, 3, 4, 5, 11, 12, 13, 14, 15, 16},
			6,
		},
		{
			"simple test - 2",
			[]int{5, 6, 2, 15, 3, 1, 4, 12, 10, 9, 8, 11, 16},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findLongestCommonSequence(tt.nums)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("findLongestCommonSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
