package subarray

func findMaxMinSubarrXor_K(input []int, k int) (int, int) {
	//result
	longest, smallest := -1, -1

	//state vars
	prefixXor := 0
	earliestPrefixXorIdxMap := make(map[int]int, 0)
	latestPrefixXorIdxMap := make(map[int]int, 0)

	// seed vals - this is needed for prefix starting 0th position
	earliestPrefixXorIdxMap[0] = -1
	latestPrefixXorIdxMap[0] = -1

	for i, num := range input {
		prefixXor ^= num

		required := prefixXor ^ k

		if idx, ok := earliestPrefixXorIdxMap[required]; ok {
			// check the length with the earlier longest length
			length := i - idx
			if length > longest {
				longest = length
			}
		}

		if idx, ok := latestPrefixXorIdxMap[required]; ok {
			// check the length with the earlier longest length
			length := i - idx
			if smallest == -1 || length < smallest {
				smallest = length
			}
		}

		// maps updates if prefixXor exist or not
		if _, ok := earliestPrefixXorIdxMap[prefixXor]; !ok {
			earliestPrefixXorIdxMap[prefixXor] = i
		}

		// update latest map irrespective
		latestPrefixXorIdxMap[prefixXor] = i
	}

	return longest, smallest
}
