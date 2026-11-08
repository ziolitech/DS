package pairs

import "testing"

func Test_countValidTriplets(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		a    []int
		d    int
		want int
	}{
		{
			"simple valid triplet",
			[]int{1, 2, 3, 4, 5, 6}, // {4,2,6}, {1,5,6}, {1,2,3}
			3,
			8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countValidTriplets(tt.a, tt.d)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("countValidTriplets() = %v, want %v", got, tt.want)
			}
		})
	}
}
