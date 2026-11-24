package amz

// Understanding :-> Given an array :-> count the number of good subarrays;

// Definition :-> [i……j] is good if sum[i……j]%k == length of that subarray

// Brute Force : Iterate over all subarrays and maintain their
// sum and check if the condition is getting satisfied or not.
// Take the subarray - find the sum .. calc the remainder .. check if it is equal
// to (j-i+1)

// idx : 0 1 2 3 4
// nums: 1 2 3 4 5

// Approach optimised
// simple version :
// prefix_sum[i..j] = j-i+1 => prefix_sum[j] - prefix_sum[i-1] == j-i+1
// sum[j]-j = sum[i-1] - (i - 1)
// i.e : key[k] = prefixSum[k] - k.. it means we see if prefixSum[j]-j occurs prior if yes .. then
// that many LHS satisfy the condition. So, we can have a map of (map[prefixSum[j]-j] = freq)
// we can store 'sum[l] - l + 1' as prefix sum? we don't know l on the left yet.
// One important detail: initialize freq[0] = 1 to account for subarrays starting at index 0.
// we keep storing p[r] - r in prefix array ...

// and check if that is equal to length of the array.

// simple version
// prefix_sum[i..j] = j-i+1
func countTheNumberOfGoodSubarraySimple(input []int) int {
	if len(input) == 0 {
		return 0
	}

	// result arr
	count := 0

	// internal state vars
	prefixSumMap := make(map[int]int, len(input)+1)
	prefixSum := 0

	// needed for subarrays to count from index 0.
	prefixSumMap[0] = 1

	// core logic.
	for i := 0; i < len(input); i++ {
		// running sum variable.
		prefixSum += input[i]

		// why i+1 ? ..read more here :
		// https://chatgpt.com/s/t_6b05d3919af08191adbf8e0356609bda
		key := prefixSum - (i + 1)

		// seems like we need to iterate the whole .. map ?
		// checking sol.
		if frq, ok := prefixSumMap[key]; ok {
			count += frq
		}

		// update the map
		prefixSumMap[key]++
	}

	return count
}

// TODO : Need to deep dive to realise ..

// Approach optimised
// simple version :
// prefix_sum[i..j]%k = j-i+1 => prefix_sum[j]%k - prefix_sum[i-1]%k == (j-i+1)%k
// sum[j]-j = sum[i-1] - (i - 1)
// i.e : key[k] = prefixSum[k] - k.. it means we see if prefixSum[j]-j occurs prior if yes .. then
// that many LHS satisfy the condition. So, we can have a map of (map[prefixSum[j]-j] = freq)
// we can store 'sum[l] - l + 1' as prefix sum? we don't know l on the left yet.
// One important detail: initialize freq[0] = 1 to account for subarrays starting at index 0.
// we keep storing p[r] - r in prefix array ...

// and check if that is equal to length of the array.

// simple version
// prefix_sum[i..j] = j-i+1
func countTheNumberOfGoodSubarray(input []int, k int) int {
	if len(input) == 0 {
		return 0
	}

	// result arr
	count := 0

	// internal state vars
	prefixSumMap := make(map[int]int, len(input)+1)
	prefixSum := 0

	// needed for subarrays to count from index 0.
	prefixSumMap[0] = 1

	// core logic.
	for i := 0; i < len(input); i++ {
		// running sum variable.
		prefixSum += input[i]

		// why i+1 ? ..read more here :
		// https://chatgpt.com/s/t_6b05d3919af08191adbf8e0356609bda
		key := prefixSum - (i + 1)

		// seems like we need to iterate the whole .. map ?
		// checking sol.
		if frq, ok := prefixSumMap[key]; ok {
			count += frq
		}

		// update the map
		prefixSumMap[key]++
	}

	return count
}
