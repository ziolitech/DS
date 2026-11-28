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
// Mod version :
// prefix_sum[i..j]%k = j-i+1 => prefix_sum[j]%k - prefix_sum[i-1]%k == (j-i+1)%k
//  For safety purpose; we add %k & +k wherever required to normalize
// as prefixSum - i can lead to negative nums:-
// -> (p[j]%k -j%k + k)%k = (p[i-1]%k -(i - 1)%k + k)%k
// Define our prefix state
// Let's define:
//
// R[x]=(P[x]-x + k)%k
// residue[i] = ((prefixSum - int64(i))%K + K) % K

/*
1 <= n <= 2*10⁵.. n is the length of array
1 <= input[i] <= 10⁹ .. input[i] represent value range of each element
1 <= k <= 10⁹ ... k is the number which divides the sum of the subarray.

constraint that the subarray length cannot be greater than or equal to k.
0 <= sum(i...j) < k

Note: Based on constraint we use the int64 as return type
#IMP #Sliding_Window
*/
func countTheNumberOfGoodSubarray(input []int, k int) int64 {
	n := len(input)
	if n == 0 || k == 1 {
		return 0
	}

	// result arr
	var count int64

	// using the bigger val
	K := int64(k)

	// internal state vars
	// key : normalised prefix sum R[x] . value = frequency of previous occurrrence R[i] for given R[j].
	freq := make(map[int64]int64, len(input)+1)
	var prefixSum int64
	// to store prefix state : (PrefixSum - position) % K
	residue := make([]int64, n+1)

	// creating normalised residue
	for i := 1; i < len(input); i++ {
		// running sum variable.
		prefixSum += int64(input[i-1])

		residue[i] = (prefixSum%K - int64(i)%K + K) % K
	}

	// loop to calculate the satifying conditions
	// we go upto j=n as thats where the prefix or residue/remainder lie.
	for j := 1; j <= n; j++ {
		// adding the previois residue to the map.. or including to the window
		// takes care of the boudary start condition.
		freq[residue[j-1]]++ // for j = 1, residue[0] would be 0 freq. (Need for Seed)

		// Window handling
		// check for item pos which need to deleted out of window
		expired := j - k
		if expired >= 0 {
			// redude the freq ..
			freq[residue[expired]]--

			if freq[residue[expired]] == 0 {
				delete(freq, residue[expired])
			}
		}

		count += freq[residue[j]]
	}

	return count
}
