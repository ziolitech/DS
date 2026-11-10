package subarray

// # IMP
// Given an array, cout the number of subarrays present with sum as k.
// TODO : Try different variations of this as per https://chatgpt.com/s/t_6af0b0775500819180943a4e329790a0

// Flow of optimisation :
// Brute Force -> Optimised wih prefixSum array + Hashmap
// -> Final optimised using just a var + Hashmap (saves on the space.)

// Optimsied Approach : We leverage prefix sum to find the subarray
// p[r] - p[l-1] = k { where we check from l -> r if the subarray sum is k ... }
// we don't to maintain a prefix complete array .. we can maintain a variable
// to know the prefixSum till current num as we are just interested in finding the required value of
// p[l-1] = p[r] - k on the left hand side.
func CountSubarraysWith_K_Sum(arr []int, k int) int {
	freq := make(map[int]int)

	// IMP : seed value - considering imaginary numbers from -∞ to -1
	//  as they dont exist so prefix sum is 0. This is needed when we need to calculate the subarray from
	// the very first beginning of the array
	freq[0] = 1

	prefixSum := 0
	count := 0

	for _, num := range arr {
		prefixSum = prefixSum + num

		needed := prefixSum - k

		// add the needed freq to count
		count += freq[needed]

		//IMP : Remember we can not put prefixSum early in the map as it
		// will hamper the calculation - lets say if k = 0, then it will look for prefixSum in the
		// map .. we were trying to check the prefixSum occured earlier as freq not including this current
		// prefixSum value.
		freq[prefixSum]++

	}

	return count
}
