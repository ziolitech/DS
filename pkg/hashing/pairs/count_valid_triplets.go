package pairs

import "fmt"

// Given an array find valid triplets such that (a[i] + a[j] + a[k]) % d = 0
// Constraint:
// N (size of array) :  3 <= N <= 10³
// 2 <= a[i] <= 10⁹
// 2 <= d <= 10⁶

// Logic
// (a[i] + a[j]) % d = x  assuming p = a[k] %d
//  then x = nd - p ∀ n ∈ [0,1,2 ...] as x%d will always result in -p

// Approach:
// we fix the k and find the pair (i,j) on the left hand side.
// we go from k - 2 ..N-1
// we keep finding pair which has value as [(d-p)%d] on the left hand side of the running hashmap.
func countValidTriplets(a []int, d int) int {
	// we need to have a frequency map
	pairSumMap := make(map[int]int, 0)

	// seed value.
	pairSumMap[(a[0]%d+a[1]%d)%d] = 1

	// triplet count
	count := 0

	for i := 2; i < len(a); i++ {
		aiN := a[i] % d
		need := (d - aiN) % d

		if freq, ok := pairSumMap[need]; ok {
			count = count + freq
			fmt.Println("Need : ", need, "aiN : ", aiN, "count : ", count)
		}

		// add all the normalized pairs with the current number
		for j := 0; j < i; j++ {
			// normalised pair sum
			// {1, 2, 3, 4, 5, 6}
			pairSum := (a[j]%d + a[i]%d) % d
			pairSumMap[pairSum] = pairSumMap[pairSum] + 1
			fmt.Println(" pairSum is : ", pairSum, "updated freq : ", pairSumMap[pairSum])
		}
	}

	return count
}
