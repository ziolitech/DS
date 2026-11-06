package hashing

// Given an array or slice, find if pair exists whose diff is K

func checkPairExistsWithDiff(arr []int, diff int) bool {
	m := map[int]struct{}{}

	for i := range arr {
		if _, ok := m[arr[i]]; ok {
			return true
		} else {
			target := arr[i] + diff
			m[target] = struct{}{}
		}
	}

	return false
}
