package problems

type Pair struct {
	e int // element
	p int // parent
}

// Below detects cycle in just one connected component.
func detect(adjList map[int][]int, start int, vis *map[int]struct{}) bool {
	queue := []Pair{}
	queue = append(queue, Pair{start, -1})

	(*vis)[start] = struct{}{}

	for len(queue) > 0 {
		frontPair := queue[0]

		// dequeue
		queue = queue[1:]

		// get adj neighbours and push them to queue
		// which are not equal to parents.
		for _, n := range adjList[frontPair.e] {
			// if not visited.
			if _, ok := (*vis)[n]; !ok {
				// below condition might be redundant as the parent can be already visited.
				(*vis)[n] = struct{}{}
				queue = append(queue, Pair{n, frontPair.e})
			} else if n != frontPair.p {
				// if the neighbour is visited, but not equal to parent
				// it means that .. neighbour which is parent is already visited and hence the cycle
				return true
			}
		}
	}

	return false
}

// find the cycle in the graph.
func FindCycle(adjList map[int][]int) bool {
	vis := make(map[int]struct{})

	// just looking for the key in the map.
	for k := range adjList {
		if _, ok := vis[k]; !ok {
			// not visited
			res := detect(adjList, k, &vis)
			if res == true {
				// it indicates the presence of a cycle
				return res
			}
		}
	}

	return false
}
