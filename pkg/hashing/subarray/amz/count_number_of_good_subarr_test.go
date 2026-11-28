package amz

import "testing"

func Test_countTheNumberOfGoodSubarraySimple(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		input []int
		want  int
	}{
		{
			"simple count number of good subarr -1",
			[]int{1},
			1,
		},
		{
			"simple count number of good subarr -2",
			[]int{1, 1, 1},
			6,
		},
		{
			"simple count number of good subarr -2",
			[]int{2, 0, 1},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countTheNumberOfGoodSubarraySimple(tt.input)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("countTheNumberOfGoodSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countTheNumberOfGoodSubarray(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		input []int
		k     int
		want  int64
	}{
		{
			"count number of good subarr",
			[]int{2, 0, 1},
			3,
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countTheNumberOfGoodSubarray(tt.input, tt.k)
			// TODO: update the condition below to compare got with tt.want.
			if tt.want != got {
				t.Errorf("countTheNumberOfGoodSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
