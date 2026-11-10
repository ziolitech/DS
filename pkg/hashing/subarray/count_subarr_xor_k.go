package subarray

// approach : maintain prefix running xor and store previous xor
// freq in a map.
func countSubArrWithXOR_K(input []int, k int) int {
	//result - count of subaarrays with XOR k
	count := 0

	// state vars
	prefixXOR := 0
	prefixXorFreqMap := make(map[int]int, 0)

	// seed -  this is needed as a seed value if - lets say xor value
	// from 0 ..i = k in that case we want p[-1] xor until -1th position
	// which is imaginary or dont exist but we assum xor until that position is
	// 0 with freq of 1. This helps add up subarray from [0 ... i] into the count.
	prefixXorFreqMap[0] = 1

	// for i := 0; i < len(input); i++ {
	for _, val := range input {
		prefixXOR = prefixXOR ^ val
		// needed to check on the left of the array
		needed := prefixXOR ^ k

		if freq, ok := prefixXorFreqMap[needed]; ok {
			count += freq
		}

		prefixXorFreqMap[prefixXOR] += 1
	}

	return count
}
