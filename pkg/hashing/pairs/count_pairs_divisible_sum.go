package pairs

// # Given an array, find count of pairs whose sum is is divisible by K
// Approch : J trick .. maintain a running hashset to check if the normalized value
// of target is present or not.
// a[i]%k + a[j]%k == 0 => Means that sum should be 0 or divisible by k

// IMP Note : if we just use a hashmap without storing the frequency, we will
// be missing out if multiple numbers exists to form a pair with on the left hand side.

// SC : min(O(N),O(K)), TC = O(N) as we use a running map to check pairs

func countPairsDivisibleByK(arr []int, k int) int {
	nMap := map[int]struct{}{}
	count := 0
	for i := range len(arr) {
		// ai_n means a[i] normalised
		ai_n := arr[i] % k
		// we do again mod as k - ai_n can be k as ai_n is 0.
		need := (k - ai_n) % k
		if _, ok := nMap[need]; ok {
			// target found
			count++
		}

		// store of the normalised val of a[i]
		nMap[arr[i]%k] = struct{}{}
	}

	return count
}
