package subarray

// Intuition :
// As we need to get the array which are non - overlapping somehow we need to divide the main array
// into two halves .. by following a blue stick method .. which divides the array into 2 halves.
// this stick we will place in every position to divide the array - from 1 to N-1 position.(1-based) (check)
// when divided for the left half - left half we can just maintain a variable to know the subarray max sum including that element.
// Here we dont need to know Pmax which is the max subarray sum on the left side ..as that will be considered already while knowing
// the total subarray sum on the left and right and when the blue stick was until that element.
// So we just need know the p array on the left side or we can maintain a var to know the max until
// that element(including that element) and for the right hand side we can create a Smax array (dependent on the suffix) array
// to know the sum until that element.
// 📥 i/p : array
// 📤 o/p : max non ovelapping sum
func countTwoNonOverlapSubarray(input []int) int {
	// state vars
	prefix := -1
	size := len(input)
	// suffix array
	s := make([]int, size+1)
	// smax array
	smax := make([]int, size+1)

	// o/p vars
	result := 0

	// pre-compute stuff - create suffix array and smax array
	for i := size - 1; i >= 0; i-- {
		s[i] = max(input[i]+s[i+1], input[i])
	}

	for i := size - 1; i >= 0; i-- {
		// smax[size] would be zero.
		smax[i] = max(smax[i+1], s[i])
	}

	// core logic involving - blue stick.
	// Loop goes until (n-1)th position as we want to have a right party.
	for bluestick := 0; bluestick <= size-1; bluestick++ {
		left := bluestick
		right := bluestick + 1

		// left part sum
		if bluestick == 0 {
			prefix = input[bluestick]
		} else {
			prefix = max(prefix+input[bluestick], input[bluestick])
		}
		left = prefix

		right = smax[right]

		result = max(result, left+right)
	}

	return result
}
