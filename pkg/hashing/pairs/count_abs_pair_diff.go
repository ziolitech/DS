package pairs

// Given an array, find pairs (i,j) ∀ i < j | abs(arr[i] - arr[j]) = target
// |arr[i] + arr[j]| = arr[i] - arr[j] or -(arr[i] - arr[j])
// posibility 1 : a[j] = a[i] - target
// posibility 2 : a[j] = a[i] + target

func countPairExistsWithAbsDiff(arr []int, targetDiff int) int {
	m := map[int]struct{}{}

	count := 0
	for i := range arr {
		if _, ok := m[arr[i]]; ok {
			count++
		}

		target1 := arr[i] + targetDiff
		target2 := arr[i] - targetDiff
		m[target1] = struct{}{}
		m[target2] = struct{}{}
	}

	return count
}
