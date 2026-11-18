package misc

import "testing"

func Test_findCommonChars(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		input []string
		want  string
	}{
		{
			"simple test - 1",
			[]string{"abczdfs", "abcmnop", "abcitej"},
			"abc",
		},
		{
			"simple test - 2",
			[]string{"abczdfs", "abcmnop", ""},
			"",
		},
		{
			"simple test - 3",
			[]string{"abczdfs", "abcmnop", "a"},
			"a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findCommonChars(tt.input)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("findCommonChars() = %v, want %v", got, tt.want)
			}
		})
	}
}
