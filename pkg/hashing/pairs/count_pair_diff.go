package pairs

// Given an array or slice, find if pair exists whose diff is K
// TC = O(n) in worst case, SC = O(n) in worst case

// Not considering the duplicates.

func countPairExistsWithDiff(arr []int, diff int) int {
	m := map[int]struct{}{}

	count := 0
	for i := range arr {
		println("i val : ", i)
		if _, ok := m[arr[i]]; ok {
			count++
		}

		// this needs to be put irrespective as below as a target will be used for other nums.
		target := arr[i] + diff
		m[target] = struct{}{}
	}

	return count
}
