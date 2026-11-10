package subarray

import "testing"

func Test_findMaxMinSubarrSum_K(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		input []int
		k     int
		want  int
		want2 int
	}{
		{
			"Simple TC : Find max min subarr with K sum",
			[]int{1, 4, 4, 5, 6, 9},
			9,
			3,
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got2 := findMaxMinSubarrSum_K(tt.input, tt.k)
			if got != tt.want {
				t.Errorf("findMaxMinSubarrSum_K() = %v, want %v", got, tt.want)
			}
			if got2 != tt.want2 {
				t.Errorf("findMaxMinSubarrSum_K() = %v, want %v", got2, tt.want2)
			}
		})
	}
}
