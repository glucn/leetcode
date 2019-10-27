package solution

// Runtime 16 ms
// Memory 5.9 MB

func missingNumber(nums []int) int {
	total := 0
	for i := range nums {
		total += nums[i]
	}
	return (len(nums)*(len(nums)+1))>>1 - total
}
