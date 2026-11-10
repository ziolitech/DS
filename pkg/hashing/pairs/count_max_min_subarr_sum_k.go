package pairs

// Given an array find the max and min length of the subarray which has sum as K
// Think 💡 : Did you think about the edge case ? array just having one element and
// k equal/uneual to that number, empty
func findMaxMinSubarrSum_K(input []int, k int) (int, int) {
	// ✅ state vars
	// stores running prefix sum
	prefixSum := 0
	// stores key = prefix_sum and value as index
	latest := make(map[int]int, 0)
	earliest := make(map[int]int, 0)

	// ✅ seed values
	// this is needed to account for the array starting from the 0th index.
	latest[0] = -1
	earliest[0] = -1

	// ✅ result vars
	// stores largest and smallest subarray sum
	longest, smallest := -1, -1

	for i, num := range input {
		prefixSum += num
		needed := prefixSum - k

		// check if needed exists in the map
		if idx, ok := earliest[needed]; ok {
			length := i - idx
			if length > longest {
				longest = length
			}
			// dont update
		}

		// update the prefix sum index in the map

		// check if needed exists in the map
		if idx, ok := latest[needed]; ok {
			length := i - idx
			// IMP (Edge cond) : Below condition will fail when
			// we have just single element in the given input array and K equal that element.
			// idx found would be -1 in that case and length will be 1 and {length(1) < smallest(-1)}
			// would turn out to be false. So, we need to add a special condition to check if idx == -1
			if idx == -1 || length < smallest {
				smallest = length
			}
		}

		// Updation
		// if not present in the earliest map - then update the map
		if _, ok := earliest[prefixSum]; !ok {
			earliest[prefixSum] = i
		}

		// update latest in any case
		latest[prefixSum] = i
	}

	// Update the longest and smallest maps

	return longest, smallest
}
