package misc

// https://leetcode.com/problems/longest-consecutive-sequence/description/
// Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

// You must write an algorithm that runs in O(n) time.
// Constraints:

// 0 <= nums.length <= 10⁵
// -10⁹ <= nums[i] <= 10⁹ (1 billion)

// Approach - create a set of starting points.
func findLongestCommonSequence(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// state var
	uniqueIntMap := make(map[int]bool)
	startNums := make([]int, 0, 0)

	// result var
	result := 0
	for i := range nums {
		uniqueIntMap[nums[i]] = false
	}

	for i := range nums {
		exists := checkPrevExists(uniqueIntMap, nums[i])
		if !exists {
			startNums = append(startNums, nums[i])
		}
	}

	for i := range startNums {
		length := findLongLengthForeward(uniqueIntMap, startNums[i], 1)
		result = max(result, length)
	}

	return result
}

func checkPrevExists(uniqueIntMap map[int]bool, num int) bool {
	key := num - 1
	if _, ok := uniqueIntMap[key]; ok {
		return true
	}

	return false
}

func findLongLengthForeward(uniqueIntMap map[int]bool, key int, length int) int {
	newKey2 := key + 1
	if visited, ok := uniqueIntMap[newKey2]; ok && !visited {
		// uniqueIntMap[newKey2] = true
		length++
		length = findLongLengthForeward(uniqueIntMap, newKey2, length)
	}

	return length
}

// TODO Analyse time complexity
