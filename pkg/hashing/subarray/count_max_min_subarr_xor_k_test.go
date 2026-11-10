package subarray

import "testing"

func Test_findMaxMinSubarrXor_K(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		input []int
		k     int
		want  int
		want2 int
	}{
		{
			"Simple test - find max min xor",
			[]int{4, 2, 2, 6, 4},
			6,
			5,
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got2 := findMaxMinSubarrXor_K(tt.input, tt.k)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("findMaxMinSubarrXor_K() = %v, want %v", got, tt.want)
			}
			if got2 != tt.want2 {
				t.Errorf("findMaxMinSubarrXor_K() = %v, want %v", got2, tt.want2)
			}
		})
	}
}
