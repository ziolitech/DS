package misc

import "testing"

func Test_reduceHeightOfPile(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		input []int
		want  int
	}{
		{
			"simple test - reduce heights",
			[]int{4, 5, 5, 2, 4},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reduceHeightOfPile(tt.input)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("reduceHeightOfPile() = %v, want %v", got, tt.want)
			}
		})
	}
}
