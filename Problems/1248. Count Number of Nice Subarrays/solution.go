package solution

// Runtime 120 ms
// Memory 8.2 MB

func numberOfSubarrays(nums []int, k int) int {
	odd := make([]int, 0, len(nums))
	for i := range nums {
		if nums[i]%2 == 1 {
			odd = append(odd, i)
		}
	}
	if len(odd) < k {
		return 0
	}
	odd = append([]int{-1}, odd...)
	odd = append(odd, len(nums))
	// fmt.Println(odd)
	result := 0
	for i := 1; i+k < len(odd); i++ {
		result += (odd[i] - odd[i-1]) * (odd[i+k] - odd[i+k-1])
		// fmt.Println(odd[i]-odd[i-1], odd[i+k]-odd[i+k-1])
		// fmt.Println(result)
	}
	return result
}
