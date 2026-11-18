package misc

import (
	"sort"
)

// Given a pile with height of boxes .. we need to reduce the height of the pile
// to make all of the pile of boxes equal at the end. count the min number of ways.
// Note : Max heap way of implementing this is covered under rand-tests.
func reduceHeightOfPile(input []int) int {

	// state vars
	freq := make(map[int]int, 0)

	// loop to build the map
	for i := range input {
		num := input[i]

		// simply add the frequenecies.
		freq[num]++
	}

	// now we need to delete the highest freq and make it equal to second highest
	// while delete the highest freq .. second highest .. will become highest and
	// .. we need to find subsequent second highest .. either we use heap .. or a
	// sorted set (How it is implemented ..? ) - Check ? - we need to sort by val.

	// sort by values

	// Collect distinct heights.
	distinct := make([]int, 0, len(freq))

	for h := range freq {
		distinct = append(distinct, h)
	}

	// Sort distinct heights in descending order. - understand this.
	sort.Sort(sort.Reverse(sort.IntSlice(distinct)))

	// Sort distinct heights in descending order.
	sort.Sort(sort.Reverse(sort.IntSlice(distinct)))

	operations := 0
	cumulativeCount := 0

	for i, h := range distinct {
		// Every pile at this height needs to be
		// reduced to the next distinct height.
		if i == len(distinct)-1 {
			break
		}

		// cumalative count is needed to count the previous
		// number of ops + number of ops with 2nd highest elements frequency.
		cumulativeCount = cumulativeCount + freq[h]
		// operations track all the transition operations.
		operations += cumulativeCount
	}

	return operations

}

// Break down the implementation:
// Operation	Complexity	Explanation
// Build frequency map	O(n) average	Visit every pile once
// Extract distinct heights	O(k)	Visit frequency map
// Sort distinct heights	O(k log k)	Sort only unique heights
// Calculate operations	O(k)	Iterate distinct heights

// Therefore:
// \[
// T(n,k)=O(n)+O(k)+O(klogk)+O(k)
// \]

// Ignoring lower-order terms:
//
// O(n+klog k)
//

// Space complexity:
//
// O(k)
//

// This is particularly useful when there are many duplicate heights.

// For example:
// n = 1,000,000
// k = 10

// You only sort 10 distinct heights rather than one million individual values.
