package problems

type Data[T any] struct {
	x T
	y T
	d T
}

// input - matrix with 0's and 1's
// op - matrix with the distance.
func FindNearestOne(input [][]int) [][]int {
	queue := make([]Data[int], 0)

	// make other 2D slices
	visited := make([][]int, len(input))
	distance := make([][]int, len(input))

	// make the input slices.
	for i := range input {
		visited[i] = make([]int, len(input[0]))
		distance[i] = make([]int, len(input[0]))
	}

	lx, ly := len(input), len(input[0]) // lx is the length of the row, ly is the length of the column.
	for i := 0; i < lx; i++ {
		for j := 0; j < ly; j++ {
			if input[i][j] == 1 {
				queue = append(queue, Data[int]{i, j, 0})
				visited[i][j] = 1
			}
		}
	}

	for len(queue) > 0 {
		front := queue[0]
		// deque
		queue = queue[1:]

		// update the distance matrix
		distance[front.x][front.y] = front.d

		// navigate to different directions
		moves := [][]int{
			{-1, 0}, // top
			{0, 1},  // right
			{1, 0},  // down
			{0, -1}, // left
		}

		for i := 0; i < 4; i++ {
			nRow := front.x + moves[i][0]
			nCol := front.y + moves[i][1]

			// indicates move within the bound
			if ((nRow) >= 0 && (nRow < lx)) && ((nCol) >= 0 && ((nCol) < ly)) {
				// check if not visited
				if visited[nRow][nCol] == 0 {
					visited[nRow][nCol] = 1
					// push to the queue
					queue = append(queue, Data[int]{nRow, nCol, front.d + 1})
				}
			}
		}
	}

	return distance
}
