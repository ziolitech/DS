package problems_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/craftizmv/DS/pkg/graph/adj-list/problems"
)

func TestFindNearestOne(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		input [][]int
		want  [][]int
	}{
		{
			"sample test 1",
			[][]int{
				{0, 0, 1},
				{0, 1, 0},
				{0, 0, 0},
			},
			[][]int{
				{2, 1, 0},
				{1, 0, 1},
				{2, 1, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := problems.FindNearestOne(tt.input)

			fmt.Println("Got ...", got)

			for i := range got {
				x := got[i]
				if !slices.Equal(x, tt.want[i]) {
					t.Errorf("BFS() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
