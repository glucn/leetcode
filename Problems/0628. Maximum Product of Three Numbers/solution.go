package solution

import "math"

// Runtime 52 ms
// Memory 6.3 MB

func maximumProduct(nums []int) int {
	min1, min2 := math.MaxInt32, math.MaxInt32
	max1, max2, max3 := -math.MaxInt32, -math.MaxInt32, -math.MaxInt32

	for _, n := range nums {
		if n < min1 {
			min1, min2 = n, min1
		} else if n < min2 {
			min2 = n
		}

		if n > max1 {
			max1, max2, max3 = n, max1, max2
		} else if n > max2 {
			max2, max3 = n, max2
		} else if n > max3 {
			max3 = n
		}
	}

	result1 := min1 * min2 * max1
	result2 := max1 * max2 * max3

	if result1 > result2 {
		return result1
	}
	return result2
}
