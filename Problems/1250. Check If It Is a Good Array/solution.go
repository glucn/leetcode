package solution

// Runtime 56 ms
// Memory 9 MB

func isGoodArray(nums []int) bool {
	g := nums[0]
	for i := 0; i < len(nums); i++ {
		g = gcd(g, nums[i])
		if g == 1 {
			return true
		}
	}
	return false
}

func gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}
