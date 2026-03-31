package adjlist

import (
	"fmt"

	"github.com/lossdev/stack"
)

// DFSRec - implementation with passing a slice as value, IMP : internal changes in slice will
// change the header pointing to a different struct backing a different array.
// so - return the header so that it can be propogated upwards in the call-stack.
func (g2 *Graph2) DFSRec(start int, visited map[int]struct{}, result []int) []int {
	for i := range g2.AdjList[start] {
		num := g2.AdjList[start][i]
		if _, ok := visited[num]; !ok {
			// if not visited then call the recursive dfs func.
			visited[num] = struct{}{}
			result = append(result, num)
			// we need to store the result in slice - as header
			// will keep changing on re-allocations of objects.
			// it needs to be ``` result = g2.DFSRec(num, visited, result) ```
			// else it will not work. or you need to point a slice struct which contains the
			// backing array - header address, capacity and the length.
			result = g2.DFSRec(num, visited, result)
		}
	}

	fmt.Println("result is : ", result, "visited : ", visited)

	return result
}

// No need to return due to the pointer to slice, it always update the same struct with
// update length information.
func (g2 *Graph2) DFSRec2(start int, visited map[int]struct{}, result *[]int) {
	for i := range g2.AdjList[start] {
		num := g2.AdjList[start][i]
		if _, ok := visited[num]; !ok {
			// if not visited then call the recursive dfs func.
			visited[num] = struct{}{}
			*result = append(*result, num)
			// we need to store the result in slice - as header
			// will keep changing on re-allocations of objects.
			// it needs to be ``` result = g2.DFSRec(num, visited, result) ```
			// else it will not work. or you need to point a slice struct which contains the
			// backing array - header address, capacity and the length.
			g2.DFSRec2(num, visited, result)
		}
	}

	fmt.Println("result is : ", result, "visited : ", visited)
}

// DFSWithStack - iterative implementation
func (g2 *Graph2) DFSWithStack(start int) []int {
	st := stack.NewGenericStack()
	visited := make(map[int]struct{}, 0)

	// starting as a seed
	st.Push(start)

	// mark seed visited
	visited[start] = struct{}{}

	dfs := make([]int, 0)

	for st.Size() > 0 {
		el, _ := st.Pop()

		// push it to result - now.
		dfs = append(dfs, el.(int))

		// get neighbours
		if neighbours, ok := g2.AdjList[el.(int)]; ok {
			// loop over neb
			for n := range neighbours {
				if _, ok := visited[n]; !ok {
					visited[n] = struct{}{}
					// push it to stack
					st.Push(n)
				}
			}
		}
	}

	return dfs
}
